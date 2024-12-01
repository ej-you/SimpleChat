package handlers


import (
	"encoding/json"
	"time"

	"github.com/gorilla/websocket"
	"github.com/google/uuid"

	coreErrorHandler "SimpleChat/backend/core/error_handler"
	"SimpleChat/backend/core/db"
	"SimpleChat/backend/core/db/models"
	"SimpleChat/backend/app_messanger/serializers"
	"SimpleChat/backend/settings"
)


// обработка входящих сообщений
func (this *client) HandleReadMessage(doneChan chan<- int) {
	defer this.Conn.Close()
	defer close(this.Message)

	// настройка таймаута чтения сообщений от клиента
	this.Conn.SetReadDeadline(time.Now().Add(settings.WebsocketPongWait))
	// настройка обработчика PONG'ов от клиента
	this.Conn.SetPongHandler(func(string) error {
		// обновление таймаута после получения PONG сообщения от клиента
		this.Conn.SetReadDeadline(time.Now().Add(settings.WebsocketPongWait))
		return nil
	})

	for {
		_, byteMessage, err := this.Conn.ReadMessage()
		if err != nil {
			// если произошла неизвестная ошибка закрытия соединения
			if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure) { // websocket.CloseGoingAway, websocket.CloseAbnormalClosure
				settings.ErrorLog.Printf("unexpectedCloseError: %v", err)
			}
			settings.InfoLog.Printf("-- Close connection with user %q\n", this.UserUUID)
			// посылаем сообщение в канал для завершения работы горутины отправки сообщений
			doneChan <- 0
			return
		}

		// структура для обработанного сообщения
		var clientMessage jsonMessageWithError
		// десериализация и валидация сообщения
		clientMessage.Error = clientMessage.JSONData.ParseAndValidate(byteMessage)

		settings.InfoLog.Printf("-- Received message from user %q\n", this.UserUUID)
		this.Message <- clientMessage
	}
}

// отправка сообщения
func (this *client) HandleWriteMessage(doneChan <-chan int) {
	// настройка тикера отправки PING сообщений для проверки активности соединения
	pongTicker := time.NewTicker(settings.WebsocketPingPeriod)

	defer pongTicker.Stop()
	defer this.Conn.Close()
	defer this.RemoveClient()

	dbStruct := db.NewDB()

	for {
		select {
			// новое сообщение в канале
			case clientMessage, ok := <- this.Message:
				// канал закрыт, закрываем соединение
				if !ok {
					this.Conn.WriteMessage(websocket.CloseMessage, []byte{})
					settings.InfoLog.Printf("-- Close connection with user %q\n", this.UserUUID)
					return
				}
				// если ошибка, то отправляем её отправителю и пропускаем дальнейшие действия
				if clientMessage.Error != nil {
					// если произошла ошибка при отправке, то прерываем соединение
					if err := this.SendError(clientMessage.Error); err != nil {
						settings.ErrorLog.Printf("failed to send error message: %v\n", err)
						return
					}
					continue
				}

				// получение uuid другого участника
				// (и проверка текущего юзера на принадлежность к чату, uuid которого он отправил в сообщении)
				participantUUID, err := serializers.GetChatParticipantUUID(clientMessage.JSONData.ChatID, this.UserUUID)
				if err != nil {
					if err := this.SendError(err); err != nil {
						settings.ErrorLog.Printf("failed to send error message: %v\n", err)
						return
					}
					continue
				}

				// добавление записи сообщения в БД
				var messageFromDB models.Message
				err = dbStruct.CreateMessage(&messageFromDB, clientMessage.JSONData.ChatID, this.UserUUID, clientMessage.JSONData.Content)
				if err != nil {
					if err := this.SendError(err); err != nil {
						settings.ErrorLog.Printf("failed to send error message: %v\n", err)
						return
					}
					continue
				}
				
				// сериализация сообщения
				byteMessage, err := json.Marshal(messageFromDB)
				if err != nil {
					if err := this.SendError(err); err != nil {
						settings.ErrorLog.Printf("failed to send error message: %v\n", err)
						return
					}
					continue
				}
				// отправка сообщения отправителю (всем подключениям юзера с текущим uuid)
				if err = this.SendMessageToClients(this.UserUUID, byteMessage); err != nil {
					settings.WarnLog.Printf("close connection with user %q\n", this.UserUUID)
					return
				}
				// поиск второго участника и отправка сообщения ему (если он есть среди подключённых клиентов)
				this.SendMessageToClients(participantUUID, byteMessage)

			// подошло время отправки очередного PING сообщения
			case <- pongTicker.C:
				if err := this.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
					settings.WarnLog.Printf("close connection with user %q\n", this.UserUUID)
					return
				}

			// из горутины чтения сообщений было послано сообщение о закрытии соединения с юзером
			case <- doneChan:
				return
		}
	}
}


// отправка сообщения с ошибкой только текущему подключению юзера
func (this *client) SendError(errorToSend error) error {
	// создание структуры ошибки
	errStruct, _ := coreErrorHandler.GetCustomErrorMessage(settings.WebsocketURLPath, errorToSend)
	// сериализация структуры ошибки
	byteMessage, err := json.Marshal(errStruct)
	if err != nil {
		return err
	}
	if err = this.Conn.WriteMessage(websocket.TextMessage, byteMessage); err != nil {
		return err
	}
	settings.InfoLog.Printf("-- Send error message to user %q: %q\n", this.UserUUID, byteMessage)
	return nil
}

// отправка сообщения всем подключённым клиентам с uuid отправителя
func (this *client) SendMessageToClients(userUUID uuid.UUID, byteMessage []byte) error {
	var loopErr, sendErr error

	// поиск соединений с юзерами с переданным uuid
	clientConnections, found := clients[userUUID]
	if !found {
		return nil
	}

	for _, c := range clientConnections {
		loopErr = c.Conn.WriteMessage(websocket.TextMessage, byteMessage)
		if loopErr != nil {
			// если ошибка при отправке сообщения текущему клиенту
			if c.Conn == this.Conn {
				sendErr = loopErr
			}
			continue
			settings.WarnLog.Printf("failed to send message to user %q (from %q)\n", userUUID, this.UserUUID)
		}
		settings.InfoLog.Printf("-- Send message to user %q (from %q)\n", userUUID, this.UserUUID)
	}
	if sendErr != nil {
		return sendErr
	}

	return nil
}
