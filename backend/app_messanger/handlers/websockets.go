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
func (c *client) HandleReadMsg(doneChan chan<- int) {
	defer c.Conn.Close()
	defer close(c.Message)

	// настройка таймаута чтения сообщений от клиента
	err := c.Conn.SetReadDeadline(time.Now().Add(settings.WebsocketPongWait))
	if err != nil {
		settings.ErrorLog.Printf("HandleReadMsg: unexpected error when setting read deadline: %v", err)
		doneChan <- 0
		return
	}
	// настройка обработчика PONG'ов от клиента
	c.Conn.SetPongHandler(func(string) error {
		// обновление таймаута после получения PONG сообщения от клиента
		return c.Conn.SetReadDeadline(time.Now().Add(settings.WebsocketPongWait))
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
func (c *client) HandleWriteMsg(doneChan <-chan int) {
	// настройка тикера отправки PING сообщений для проверки активности соединения
	pongTicker := time.NewTicker(settings.WebsocketPingPeriod)

	defer pongTicker.Stop()
	defer c.Conn.Close()
	defer c.Remove()

	for {
		select {
		// новое сообщение в канале
		case clientMessage, ok := <-c.Message:
			// канал закрыт, закрываем соединение
			if !ok {
				if err := c.Conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					settings.ErrorLog.Printf("failed to close connection with user %q: %v\n", c.UserUUID, err)
					return
				}
				settings.InfoLog.Printf("-- Close connection with user %q\n", c.UserUUID)
				return
			}
			// если ошибка, то отправляем её отправителю и пропускаем дальнейшие действия
			if clientMessage.Error != nil {
				// если произошла ошибка при отправке, то прерываем соединение
				if err := c.sendError(clientMessage.Error); err != nil {
					settings.ErrorLog.Printf("failed to send error message: %v\n", err)
					return
				}
				continue
			}
			// обработка сообщения
			c.handleNewMessage(clientMessage.JSONData)
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

// обработка нового сообщения в канале
func (c *client) handleNewMessage(messageIn serializers.MessageIn) {
	// получение uuid другого участника
	// (и проверка текущего юзера на принадлежность к чату, uuid которого он отправил в сообщении)
	participantUUID, err := serializers.GetChatParticipantUUID(messageIn.ChatID, c.UserUUID)
	if err != nil {
		if err := c.sendError(err); err != nil {
			settings.ErrorLog.Printf("failed to send error message: %v\n", err)
			return
		}
		return
	}

	// добавление записи сообщения в БД
	var messageFromDB models.Message
	err = db.NewDB().CreateMessage(&messageFromDB, messageIn.ChatID, c.UserUUID, messageIn.Content)
	if err != nil {
		if err := c.sendError(err); err != nil {
			settings.ErrorLog.Printf("failed to send error message: %v\n", err)
			return
		}
		return
	}

	// сериализация сообщения
	byteMessage, err := json.Marshal(messageFromDB)
	if err != nil {
		if err := c.sendError(err); err != nil {
			settings.ErrorLog.Printf("failed to send error message: %v\n", err)
			return
		}
		return
	}
	// отправка сообщения отправителю (всем подключениям юзера с текущим uuid)
	if err = c.sendMessageToClients(c.UserUUID, byteMessage); err != nil {
		settings.WarnLog.Printf("close connection with user %q\n", c.UserUUID)
		return
	}
	// поиск второго участника и отправка сообщения ему (если он есть среди подключённых клиентов)
	if err = c.sendMessageToClients(participantUUID, byteMessage); err != nil {
		settings.WarnLog.Printf("error when sending message to participant connections: %v\n", err)
		return
	}
}

// отправка сообщения с ошибкой только текущему подключению юзера
func (c *client) sendError(errorToSend error) error {
	// создание структуры ошибки
	errStruct, _ := coreErrorHandler.GetCustomErrorMessage(settings.WebsocketURLPath, errorToSend)
	// сериализация структуры ошибки
	byteMessage, err := json.Marshal(errStruct)
	if err != nil {
		return err
	}
	err = c.Conn.WriteMessage(websocket.TextMessage, byteMessage)
	if err != nil {
		return err
	}
	settings.InfoLog.Printf("-- Send error message to user %q: %q\n", c.UserUUID, byteMessage)
	return nil
}

// отправка сообщения всем подключённым клиентам с uuid отправителя
func (c *client) sendMessageToClients(userUUID uuid.UUID, byteMessage []byte) error {
	var loopErr, sendErr error

	// поиск соединений с юзерами с переданным uuid
	clientConnections, found := clients.ClientsMap[userUUID]
	if !found {
		return nil
	}

	for _, clientConn := range clientConnections {
		loopErr = clientConn.Conn.WriteMessage(websocket.TextMessage, byteMessage)
		if loopErr != nil {
			// если ошибка при отправке сообщения текущему клиенту
			if clientConn.Conn == c.Conn {
				sendErr = loopErr
			}
			settings.WarnLog.Printf("failed to send message to user %q (from %q)\n", userUUID, clientConn.UserUUID)
			continue
		}
		settings.InfoLog.Printf("-- Send message to user %q (from %q)\n", userUUID, clientConn.UserUUID)
	}
	if sendErr != nil {
		return sendErr
	}

	return nil
}
