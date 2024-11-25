package handlers

import (
	"fmt"
	"net/http"

	echo "github.com/labstack/echo/v4"
    "github.com/gorilla/websocket"

	"SimpleChat/backend/core/services"
)


var upgrader = websocket.Upgrader{
    // Размер буфера чтения
    ReadBufferSize:  1024,
    // Размер буфера записи
    WriteBufferSize: 1024,
    // Включаем поддержку сжатия
    EnableCompression: true,
    // Разрешить любые запросы на обновление протокола с разных источников
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}


// обработка обновления соединения до WebSocket
func UpgradeWebSocket(context echo.Context) error {
    // получение uuid юзера из контекста запроса
    userUuid, err := services.GetUserIDFromRequest(context)
    if err != nil {
        return err
    }
    fmt.Printf("Connected UserID: %s\n", userUuid)

	// обновление соединения до WebSocket
    conn, err := upgrader.Upgrade(context.Response(), context.Request(), nil)
    if err != nil {
        return echo.NewHTTPError(400, map[string]string{"websocket": "failed to upgrade connection: " + err.Error()})
    }

    // создание новой структуры клиента и добавление его в список подключённых
    newClient := client{
        Conn: conn,
        Message: make(chan []byte),
    }
    clients[userUuid] = newClient
    
    go newClient.HandleReadMessage(userUuid)
    go newClient.HandleWriteMessage(userUuid)

    return nil
}
