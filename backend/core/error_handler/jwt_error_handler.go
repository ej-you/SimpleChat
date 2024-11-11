package error_handler

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
)


// настройка обработчика ошибок для JWT middleware
func CustomJWTErrorHandler(context echo.Context, err error) error {
	// ошибка валидации токена (кривой токен)
	tokenParsingError, ok := err.(*echojwt.TokenParsingError)
	if ok {
		httpError := &echo.HTTPError{
			Code: http.StatusUnauthorized,
			Message: map[string]string{"token": tokenParsingError.Error()},
		}
		return httpError
	}

	// токен не был отправлен в куках (токен и запись в куках истекли)
	_, ok = err.(*echojwt.TokenExtractionError)
	if ok {
		httpError := &echo.HTTPError{
			Code: http.StatusUnauthorized,
			Message: map[string]string{"token": "missing auth cookie"},
		}
		return httpError
	}

	return err
}
