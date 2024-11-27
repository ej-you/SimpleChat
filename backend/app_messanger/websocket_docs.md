# WebSocket Docs

## URL:
```text
GET wss://150.241.82.68/api/messanger
```

## Message:
```json5
{
    // uuid чата
    "chatId": "8098dceb-46d6-428d-9000-9ee417ec3203",
    // текст сообщения
    "content": "Hi! How do u do?"
}
```

## Response:
```json5
{
    // текст сообщения
    "content": "Hi! How do u do?",
    // время создания сообщения
    "createdAt": "2024-11-27T22:03:05.531397654+03:00",
    // отправитель
    "sender": {
        "id": "c127a824-c5ce-4de7-bdb1-7f2fc6bd1c56",
        "username": "username1"
    }
}
```

## Errors:

* Неверный формат сообщения (не JSON) или неверные данные в JSON-сообщении
```json5
{
    "status": "error",
    "statusCode": 400,
    "path": "/api/messanger",
    "timestamp": "24-11-27 22:05:46 +03",
    "errors": {
        "message": "failed to parse JSON from message: unexpected end of JSON input"
    }
}
```

* Отсутствие куки авторизации (не авторизован или истёк токен)
```json5
{
    "status": "error",
    "statusCode": 401,
    "path": "/api/messanger",
    "timestamp": "24-11-27 22:07:44 +03",
    "errors": {
        "token": "missing auth cookie"
    }
}
```

* Текущий юзер не является участником чата, uuid которого он указал в сообщении
```json5
{
    "status": "error",
    "statusCode": 403,
    "path": "/api/messanger",
    "timestamp": "24-11-27 22:07:44 +03",
    "errors": {
        "message": "forbidden"
    }
}
```

* Неизвестная ошибка
```json5
{
    "status": "error",
    "statusCode": 500,
    "path": "/api/messanger",
    "timestamp": "24-11-27 22:07:44 +03",
    "errors": {
        "unknown": "some shit"
    }
}
```
