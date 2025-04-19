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
func (c client) Add() {
	// если список с подключениями данного клиента уже есть, добавляем в него новое подключение
	if _, found := clients.ClientsMap[c.UserUUID]; found {
		clients.Lock.RLock()
		clients.ClientsMap[c.UserUUID] = append(clients.ClientsMap[c.UserUUID], c)
		clients.Lock.RUnlock()
		// если списка с подключениями данного клиента нет, то создаём его с текущим подключением
	} else {
		clients.Lock.RLock()
		clients.ClientsMap[c.UserUUID] = []client{c}
		clients.Lock.RUnlock()
	}
}

// удаление клиента из словаря
func (c client) Remove() {
	clientConnections, found := clients.ClientsMap[c.UserUUID]
	// если списка с подключениями данного клиента нет
	if !found {
		return
	}

	// если список с подключениями данного клиента содержит один элемент, то удаляем этот список из словаря
	if len(clientConnections) == 1 {
		clients.Lock.RLock()
		delete(clients.ClientsMap, c.UserUUID)
		clients.Lock.RUnlock()
		// если список с подключениями данного клиента содержит несколько элементов, то удаляем элемент текущего подключения из этого списка
	} else {
		clientIndex := slices.IndexFunc(clientConnections, func(elem client) bool {
			return elem.Conn == c.Conn
		})
		clients.Lock.RLock()
		clients.ClientsMap[c.UserUUID] = slices.Delete(clientConnections, clientIndex, clientIndex+1)
		clients.Lock.RUnlock()
	}
}
