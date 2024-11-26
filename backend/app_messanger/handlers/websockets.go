package handlers


import (
	"encoding/json"
	"fmt"
	"time"

    "github.com/gorilla/websocket"
    "github.com/google/uuid"

	coreErrorHandler "SimpleChat/backend/core/error_handler"
	"SimpleChat/backend/core/db"
	"SimpleChat/backend/core/db/models"
	"SimpleChat/backend/app_messanger/serializers"
	"SimpleChat/backend/settings"
)


// структура с обработанным сообщением и с ошибкой
type jsonMessageWithError struct {
	JSONData 	serializers.MessageIn
	Error 		error
}

type client struct {
    // соединение с клиентом
    Conn *websocket.Conn
    // канал для хранения входящего сообщения от клиента
    Message chan jsonMessageWithError
}

// словарь со всеми подключёнными клиентами
var clients = make(map[uuid.UUID]client)


// обработка входящих сообщений
func (client *client) HandleReadMessage(userUUID uuid.UUID) {
    defer client.Conn.Close()
    defer close(client.Message)
    defer delete(clients, userUUID)

    // настройка таймаута чтения сообщений от клиента
    client.Conn.SetReadDeadline(time.Now().Add(settings.WebsocketPongWait))
    // настройка обработчика PONG'ов от клиента
    client.Conn.SetPongHandler(func(string) error {
    	// обновление таймаута после получения PONG сообщения от клиента
    	client.Conn.SetReadDeadline(time.Now().Add(settings.WebsocketPongWait))
    	return nil
    })

	for {
		_, byteMessage, err := client.Conn.ReadMessage()
		if err != nil {
			// если произошла неизвестная ошибка закрытия соединения
			if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure) { // websocket.CloseGoingAway, websocket.CloseAbnormalClosure
				settings.ErrorLog.Printf("unexpectedCloseError: %v", err)
			}
			return
		}

	    // структура для обработанного сообщения
	    var clientMessage jsonMessageWithError
		// десериализация и валидация сообщения
		clientMessage.Error = clientMessage.JSONData.ParseAndValidate(byteMessage)

		fmt.Printf("Received: %v\n", clientMessage)

		client.Message <- clientMessage
	}
}

// отправка сообщения
func (client *client) HandleWriteMessage(userUUID uuid.UUID) {
	// настройка тикера отправки PING сообщений для проверки активности соединения
	pongTicker := time.NewTicker(settings.WebsocketPingPeriod)

	defer pongTicker.Stop()
	defer client.Conn.Close()
    defer delete(clients, userUUID)

    dbStruct := db.NewDB()

	for {
		select {
			// новое сообщение в канале
			case clientMessage, ok := <- client.Message:
				// канал закрыт, закрываем соединение
				if !ok {
					client.Conn.WriteMessage(websocket.CloseMessage, []byte{})
					return
				}
				// если ошибка, то отправляем её отправителю и пропускаем дальнейшие действия
				if clientMessage.Error != nil {
					// если произошла ошибка при отправке, то прерываем соединение
					if err := client.SendError(clientMessage.Error); err != nil {
						return
					}
					continue
				}

				// получение uuid другого участника
				// (и проверка текущего юзера на принадлежность к чату, uuid которого он отправил в сообщении)
				participantUUID, err := serializers.GetChatParticipantUUID(clientMessage.JSONData.ChatID, userUUID)
				if err != nil {
					if err := client.SendError(err); err != nil {
						return
					}
					continue
				}

				// добавление записи сообщения в БД
				var messageFromDB models.Message
				err = dbStruct.CreateMessage(&messageFromDB, clientMessage.JSONData.ChatID, userUUID, clientMessage.JSONData.Content)
				if err != nil {
					if err := client.SendError(err); err != nil {
						return
					}
					continue
				}
				
				// сериализация сообщения
				byteMessage, err := json.Marshal(messageFromDB)
				if err != nil {
					if err := client.SendError(err); err != nil {
						return
					}
					continue
				}
				// отправка сообщения отправителю
				if err = client.Conn.WriteMessage(websocket.TextMessage, byteMessage); err != nil {
					return
				}
				// поиск второго участника и отправка сообщения ему (если он есть среди подключённых клиентов)
				participantClient, found := clients[participantUUID]
				if found {
					if err = participantClient.Conn.WriteMessage(websocket.TextMessage, byteMessage); err != nil {
						return
					}
				}

			// подошло время отправки очередного PING сообщения
			case <- pongTicker.C:
				if err := client.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
					return
				}
		}
	}
}


// отправка сообщения с ошибкой
func (client *client) SendError(errorToSend error) error {
	// создание структуры ошибки
	errStruct, _ := coreErrorHandler.GetCustomErrorMessage(settings.WebsocketURLPath, errorToSend)
	// сериализация структуры ошибки
	byteMessage, err := json.Marshal(errStruct)
	if err != nil {
		return err
	}
	if err = client.Conn.WriteMessage(websocket.TextMessage, byteMessage); err != nil {
		return err
	}
	return nil
}
