package error_handler

import (
	"errors"
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	echo "github.com/labstack/echo/v4"
)

// настройка обработчика ошибок для JWT middleware
func CustomJWTErrorHandler(err error) error {
	var tokenParsingError *echojwt.TokenParsingError
	var tokenExtractionError *echojwt.TokenExtractionError

	switch {
	// ошибка валидации токена (кривой токен)
	case errors.As(err, &tokenParsingError):
		return &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: map[string]string{"token": tokenParsingError.Error()},
		}
	// токен не был отправлен в куках (токен и запись в куках истекли)
	case errors.As(err, &tokenExtractionError):
		return &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: map[string]string{"token": "missing auth cookie"},
		}
	default:
		return err
	}
}
