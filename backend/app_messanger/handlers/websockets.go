package handlers

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"

	"SimpleChat/backend/app_messanger/serializers"
	"SimpleChat/backend/core/db"
	"SimpleChat/backend/core/db/models"
	coreErrorHandler "SimpleChat/backend/core/error_handler"
	"SimpleChat/backend/settings"
)

// обработка входящих сообщений
func (c *client) HandleReadMessage(doneChan chan<- int) {
	defer c.Conn.Close()
	defer close(c.Message)

	// настройка таймаута чтения сообщений от клиента
	c.Conn.SetReadDeadline(time.Now().Add(settings.WebsocketPongWait))
	// настройка обработчика PONG'ов от клиента
	c.Conn.SetPongHandler(func(string) error {
		// обновление таймаута после получения PONG сообщения от клиента
		c.Conn.SetReadDeadline(time.Now().Add(settings.WebsocketPongWait))
		return nil
	})

	for {
		_, byteMessage, err := c.Conn.ReadMessage()
		if err != nil {
			// если произошла неизвестная ошибка закрытия соединения
			if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure) { // websocket.CloseGoingAway, websocket.CloseAbnormalClosure
				settings.ErrorLog.Printf("unexpectedCloseError: %v", err)
			}
			settings.InfoLog.Printf("-- Close connection with user %q\n", c.UserUUID)
			// посылаем сообщение в канал для завершения работы горутины отправки сообщений
			doneChan <- 0
			return
		}

		// структура для обработанного сообщения
		var clientMessage jsonMessageWithError
		// десериализация и валидация сообщения
		clientMessage.Error = clientMessage.JSONData.ParseAndValidate(byteMessage)

		settings.InfoLog.Printf("-- Received message from user %q\n", c.UserUUID)
		c.Message <- clientMessage
	}
}

// отправка сообщения
func (c *client) HandleWriteMessage(doneChan <-chan int) {
	// настройка тикера отправки PING сообщений для проверки активности соединения
	pongTicker := time.NewTicker(settings.WebsocketPingPeriod)

	defer pongTicker.Stop()
	defer c.Conn.Close()
	defer c.Remove()

	dbStruct := db.NewDB()

	for {
		select {
		// новое сообщение в канале
		case clientMessage, ok := <-c.Message:
			// канал закрыт, закрываем соединение
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				settings.InfoLog.Printf("-- Close connection with user %q\n", c.UserUUID)
				return
			}
			// если ошибка, то отправляем её отправителю и пропускаем дальнейшие действия
			if clientMessage.Error != nil {
				// если произошла ошибка при отправке, то прерываем соединение
				if err := c.SendError(clientMessage.Error); err != nil {
					settings.ErrorLog.Printf("failed to send error message: %v\n", err)
					return
				}
				continue
			}

			// получение uuid другого участника
			// (и проверка текущего юзера на принадлежность к чату, uuid которого он отправил в сообщении)
			participantUUID, err := serializers.GetChatParticipantUUID(clientMessage.JSONData.ChatID, c.UserUUID)
			if err != nil {
				if err := c.SendError(err); err != nil {
					settings.ErrorLog.Printf("failed to send error message: %v\n", err)
					return
				}
				continue
			}

			// добавление записи сообщения в БД
			var messageFromDB models.Message
			err = dbStruct.CreateMessage(&messageFromDB, clientMessage.JSONData.ChatID, c.UserUUID, clientMessage.JSONData.Content)
			if err != nil {
				if err := c.SendError(err); err != nil {
					settings.ErrorLog.Printf("failed to send error message: %v\n", err)
					return
				}
				continue
			}

			// сериализация сообщения
			byteMessage, err := json.Marshal(messageFromDB)
			if err != nil {
				if err := c.SendError(err); err != nil {
					settings.ErrorLog.Printf("failed to send error message: %v\n", err)
					return
				}
				continue
			}
			// отправка сообщения отправителю (всем подключениям юзера с текущим uuid)
			if err = c.SendMessageToClients(c.UserUUID, byteMessage); err != nil {
				settings.WarnLog.Printf("close connection with user %q\n", c.UserUUID)
				return
			}
			// поиск второго участника и отправка сообщения ему (если он есть среди подключённых клиентов)
			c.SendMessageToClients(participantUUID, byteMessage)

		// подошло время отправки очередного PING сообщения
		case <-pongTicker.C:
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				settings.WarnLog.Printf("close connection with user %q\n", c.UserUUID)
				return
			}

		// из горутины чтения сообщений было послано сообщение о закрытии соединения с юзером
		case <-doneChan:
			return
		}
	}
}

// отправка сообщения с ошибкой только текущему подключению юзера
func (c *client) SendError(errorToSend error) error {
	// создание структуры ошибки
	errStruct, _ := coreErrorHandler.GetCustomErrorMessage(settings.WebsocketURLPath, errorToSend)
	// сериализация структуры ошибки
	byteMessage, err := json.Marshal(errStruct)
	if err != nil {
		return err
	}
	if err = c.Conn.WriteMessage(websocket.TextMessage, byteMessage); err != nil {
		return err
	}
	settings.InfoLog.Printf("-- Send error message to user %q: %q\n", c.UserUUID, byteMessage)
	return nil
}

// отправка сообщения всем подключённым клиентам с uuid отправителя
func (c *client) SendMessageToClients(userUUID uuid.UUID, byteMessage []byte) error {
	var loopErr, sendErr error

	// поиск соединений с юзерами с переданным uuid
	clientConnections, found := clients.ClientsMap[userUUID]
	if !found {
		return nil
	}

	for _, c := range clientConnections {
		loopErr = c.Conn.WriteMessage(websocket.TextMessage, byteMessage)
		if loopErr != nil {
			// если ошибка при отправке сообщения текущему клиенту
			if c.Conn == c.Conn {
				sendErr = loopErr
			}
			continue
			settings.WarnLog.Printf("failed to send message to user %q (from %q)\n", userUUID, c.UserUUID)
		}
		settings.InfoLog.Printf("-- Send message to user %q (from %q)\n", userUUID, c.UserUUID)
	}
	if sendErr != nil {
		return sendErr
	}

	return nil
}
