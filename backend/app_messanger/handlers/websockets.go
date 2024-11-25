package handlers


import (
	"fmt"
	// "net/http"
	// "time"

    "github.com/google/uuid"
    "github.com/gorilla/websocket"
)


type client struct {
    // соединение с клиентом
    Conn *websocket.Conn
    // канал для хранения входящего сообщения от клиента
    Message chan []byte
}

// словарь со всеми подключёнными клиентами
var clients = make(map[uuid.UUID]client)


// обработка входящих сообщений
func (client *client) HandleReadMessage(userUuid uuid.UUID) {
    defer client.Conn.Close()
    defer delete(clients, userUuid)

	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			fmt.Println("ERROR (HandleReadMessage):", err)
			break
		}
		fmt.Printf("Received: %s\n", message)
		fmt.Printf("Clients: %v\n", clients)

		client.Message <- message
	}
}

// отправка сообщения
func (client *client) HandleWriteMessage(userUuid uuid.UUID) {
	defer client.Conn.Close()
    defer delete(clients, userUuid)

	for {
		select {
			case message, _ := <- client.Message:
				if err := client.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
					fmt.Println("ERROR (HandleWriteMessage):", err)
					break
				}
				fmt.Printf("Send: %s\n", message)
		}
	}
}


// // список подключённых клиентов
// var Clients = make(map[*websocket.Conn]bool)

// // время, по истечении которого соединение будет закрыто, если не будет получено новое сообщение
// var connectionTimeout = 10 * time.Second

// var upgrader = websocket.Upgrader{
//     ReadBufferSize:  1024, // Размер буфера чтения
//     WriteBufferSize: 1024, // Размер буфера записи
//     // Позволяет определить, должен ли сервер сжимать сообщения
//     EnableCompression: true,
// }


// // ответ всем подключённым клиентам
// func answerToAllClients(message []byte) {
// 	for conn := range Clients {
// 		conn.WriteMessage(websocket.TextMessage, message)
// 	}
// }

// // обработка сообщения (например, работа с БД)
// func messageHandler(conn *websocket.Conn, message []byte) {
// 	messageSuccess := fmt.Sprintf("Message %q was handled successfully!", message)
// 	log.Printf(messageSuccess)
// 	conn.WriteMessage(websocket.TextMessage, []byte(messageSuccess))
// }


// // обработка соединения
// func handleConnection(conn *websocket.Conn) {
// 	// установка таймаута для чтения сообщения
//     conn.SetReadDeadline(time.Now().Add(connectionTimeout))

//     // добавляем нового клиента в общий список
//     Clients[conn] = true
//     defer delete(Clients, conn)

//     for {
//     	// получаем сообщение от клиента
// 		mt, message, err := conn.ReadMessage()

// 		// выходим из цикла, если:
// 		// связь с клиентом прервана: *websocket.netError
// 		// или клиент пытается закрыть соединение: *websocket.CloseError
// 		if err != nil || mt == websocket.CloseMessage {
// 			break
// 		}

// 		// рассылаем сообщения всем клиентам
// 		answerToAllClients(message)

// 		// обработка сообщения
// 		go messageHandler(conn, message)

//         // обновление таймаута после успешного чтения сообщения
//         conn.SetReadDeadline(time.Now().Add(connectionTimeout))
// 	}
// }

// // обработка обновления соединения до WebSocket
// func handleUpgrade(reponse http.ResponseWriter, request *http.Request) {
// 	url := request.URL.Path
//     log.Println("URL:", url)

// 	// обновление соединения до WebSocket
//     conn, err := upgrader.Upgrade(reponse, request, nil)
//     if err != nil {
//         log.Fatal(err)
//     }
//     defer conn.Close()

//     // обработка соединения
//     handleConnection(conn)
// }
