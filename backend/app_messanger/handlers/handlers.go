package handlers

import (
    "net/http"

    "github.com/gorilla/websocket"
    echo "github.com/labstack/echo/v4"

    "SimpleChat/backend/core/services"
    "SimpleChat/backend/settings"
)

var upgrader = websocket.Upgrader{
    // Размер буфера чтения
    ReadBufferSize: 1024,
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
    userUUID, err := services.GetUserIDFromRequest(context)
    if err != nil {
        return err
    }

    // обновление соединения до WebSocket
    conn, err := upgrader.Upgrade(context.Response(), context.Request(), nil)
    if err != nil {
        return echo.NewHTTPError(400, map[string]string{"websocket": "failed to upgrade connection: " + err.Error()})
    }

    settings.InfoLog.Printf("-- Open new connection with user %q\n", userUUID)

    // создание новой структуры клиента и добавление его в список подключённых
    newClient := client{
        Conn:     conn,
        UserUUID: userUUID,
        Message:  make(chan jsonMessageWithError),
    }
    newClient.Add()

    done := make(chan int)
    go newClient.HandleReadMessage(done)
    go newClient.HandleWriteMessage(done)

    return nil
}
