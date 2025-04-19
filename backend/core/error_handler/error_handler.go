package error_handler

import (
	"errors"
	"net/http"
	"time"

	echo "github.com/labstack/echo/v4"

	"SimpleChat/backend/settings"
)

// структура кастомной ошибки
type CustomError struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	Path       string      `json:"path"`
	Timestamp  string      `json:"timestamp"`
	Errors     interface{} `json:"errors"`
}

// возвращает структуру для сериализации с полным описанием ошибки и код ответа
func GetCustomErrorMessage(path string, err error) (errorMessage CustomError, errorCode int) {
	var httpError *echo.HTTPError
	// если ли ошибка err не является структурой *echo.HTTPError
	if !errors.As(err, &httpError) {
		httpError = &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: map[string]string{"unknown": err.Error()},
		}
	}
	// если пришла *echo.HTTPError ошибка со строкой в качестве httpError.Message
	stringErrorMessage, ok := (httpError.Message).(string)
	if ok {
		httpError.Message = map[string]string{"unknown": stringErrorMessage}
	}

	// поле timestamp
	strTime := time.Now().Format(settings.TimeFmt)

	return CustomError{
		Status:     "error",
		StatusCode: httpError.Code,
		Path:       path,
		Timestamp:  strTime,
		Errors:     httpError.Message,
	}, httpError.Code
}

// настройка обработчика ошибок
func CustomErrorHandler(echoApp *echo.Echo) {
	echoApp.HTTPErrorHandler = func(err error, context echo.Context) {
		errMessage, statusCode := GetCustomErrorMessage(context.Path(), err)

		// отправка ответа
		var respErr error
		if context.Response().Committed {
			return
		}

		// если метод HEAD
		if context.Request().Method == http.MethodHead {
			respErr = context.NoContent(statusCode)
		} else {
			respErr = context.JSON(statusCode, errMessage)
		}
		// log if error occurs
		if respErr != nil {
			context.Echo().Logger.Error(respErr)
		}
	}
}
