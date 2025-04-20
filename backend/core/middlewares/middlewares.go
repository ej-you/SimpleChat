package middlewares

import (
	echoJWT "github.com/labstack/echo-jwt/v4"
	echo "github.com/labstack/echo/v4"

	coreErrorHandler "SimpleChat/backend/core/error_handler"
	"SimpleChat/backend/settings"
)

// middleware для распаковки содержимого токена в содержимое context'а запроса и валидации токена
var AuthMiddleware echo.MiddlewareFunc = echoJWT.WithConfig(echoJWT.Config{
	SigningKey:  []byte(settings.SecretForJWT),
	TokenLookup: "cookie:simple-chat-auth",
	// ErrorHandler: coreErrorHandler.CustomJWTErrorHandler,
	ErrorHandler: func(_ echo.Context, err error) error {
		return coreErrorHandler.CustomJWTErrorHandler(err)
	},
})
