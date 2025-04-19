package handlers

import (
	"slices"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"

	"SimpleChat/backend/app_messanger/serializers"
)

// структура с обработанным сообщением и с ошибкой
type jsonMessageWithError struct {
	JSONData serializers.MessageIn
	Error    error
}

type client struct {
	// соединение с клиентом
	Conn *websocket.Conn
	// id клиента, с которым установлено соединение
	UserUUID uuid.UUID
	// канал для хранения входящего сообщения от клиента
	Message chan jsonMessageWithError
}

type clientsMap struct {
	ClientsMap map[uuid.UUID][]client
	Lock       sync.RWMutex
}

func newClientsMap() *clientsMap {
	clients := &clientsMap{}
	clients.ClientsMap = make(map[uuid.UUID][]client)

	return clients
}

// словарь со всеми подключёнными клиентами
var clients = newClientsMap()

// добавление клиента в словарь
func (this client) Add() {
	// если список с подключениями данного клиента уже есть, добавляем в него новое подключение
	if _, found := clients.ClientsMap[this.UserUUID]; found {
		clients.Lock.RLock()
		clients.ClientsMap[this.UserUUID] = append(clients.ClientsMap[this.UserUUID], this)
		clients.Lock.RUnlock()
		// если списка с подключениями данного клиента нет, то создаём его с текущим подключением
	} else {
		clients.Lock.RLock()
		clients.ClientsMap[this.UserUUID] = []client{this}
		clients.Lock.RUnlock()
	}
}

// удаление клиента из словаря
func (this client) Remove() {
	clientConnections, found := clients.ClientsMap[this.UserUUID]
	// если списка с подключениями данного клиента нет
	if !found {
		return
	}

	// если список с подключениями данного клиента содержит один элемент, то удаляем этот список из словаря
	if len(clientConnections) == 1 {
		clients.Lock.RLock()
		delete(clients.ClientsMap, this.UserUUID)
		clients.Lock.RUnlock()
		// если список с подключениями данного клиента содержит несколько элементов, то удаляем элемент текущего подключения из этого списка
	} else {
		clientIndex := slices.IndexFunc(clientConnections, func(elem client) bool {
			return elem.Conn == this.Conn
		})
		clients.Lock.RLock()
		clients.ClientsMap[this.UserUUID] = slices.Delete(clientConnections, clientIndex, clientIndex+1)
		clients.Lock.RUnlock()
	}
}
