package middlewares

import (
	echo "github.com/labstack/echo/v4"
    echoJWT "github.com/labstack/echo-jwt/v4"

    coreErrorHandler "SimpleChat/backend/core/error_handler"
    "SimpleChat/backend/settings"
)


// middleware для распаковки содержимого токена в содержимое context'а запроса и валидации токена
var AuthMiddleware echo.MiddlewareFunc = echoJWT.WithConfig(echoJWT.Config{
    SigningKey: []byte(settings.SecretForJWT),
    TokenLookup: "cookie:auth",
    ErrorHandler: coreErrorHandler.CustomJWTErrorHandler,
})
