package handlers


import (
	"slices"

    "github.com/gorilla/websocket"
    "github.com/google/uuid"

	"SimpleChat/backend/app_messanger/serializers"
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
var clients = make(map[uuid.UUID][]client)


// добавление клиента в словарь
func (this client) AddClient(userUUID uuid.UUID) {
	// если список с подключениями данного клиента уже есть, добавляем в него новое подключение
	if _, found := clients[userUUID]; found {
		clients[userUUID] = append(clients[userUUID], this)
	// если списка с подключениями данного клиента нет, то создаём его с текущим подключением
	} else {
		clients[userUUID] = []client{this}
	}
}

// удаление клиента из словаря
func (this client) RemoveClient(userUUID uuid.UUID) {
	clientConnections, found := clients[userUUID]
	// если списка с подключениями данного клиента нет
	if !found {
		return
	}

	// если список с подключениями данного клиента содержит один элемент, то удаляем этот список из словаря
	if len(clientConnections) == 1 {
		delete(clients, userUUID)
	// если список с подключениями данного клиента содержит несколько элементов, то удаляем элемент текущего подключения из этого списка
	} else {
		clientIndex := slices.IndexFunc(clientConnections, func(elem client) bool {
			return elem.Conn == this.Conn
		})
		clients[userUUID] = slices.Delete(clientConnections, clientIndex, clientIndex+1)
	}
}
