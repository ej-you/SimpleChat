package error_handler

import (
	"net/http"
	"time"

	echo "github.com/labstack/echo/v4"

	"SimpleChat/backend/settings"
)


// структура кастомной ошибки
type CustomError struct {
	Status 		string `json:"status"`
	StatusCode 	int `json:"statusCode"`
	Path 		string `json:"path"`
	Timestamp 	string `json:"timestamp"`
	Errors 		interface{} `json:"errors"`
}


// возвращает структуру для сериализации с полным описанием ошибки и код ответа
func GetCustomErrorMessage(path string, err error) (CustomError, int) {
	// проверка, является ли ошибка err структурой *echo.HTTPError (приведение типов)
	httpError, ok := err.(*echo.HTTPError)
	if !ok {
		httpError = &echo.HTTPError{
			Code: http.StatusInternalServerError,
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
		Status: "error",
		StatusCode: httpError.Code,
		Path: path,
		Timestamp: strTime,
		Errors: httpError.Message,
	}, httpError.Code
}


// настройка обработчика ошибок
func CustomErrorHandler(echoApp *echo.Echo) {
	echoApp.HTTPErrorHandler = func(err error, context echo.Context) {
		errMessage, statusCode := GetCustomErrorMessage(context.Path(), err)

		// отправка ответа
		var respErr error
		if !context.Response().Committed {
			// если метод HEAD
			if context.Request().Method == http.MethodHead {
				respErr = context.NoContent(statusCode)
			} else {
				respErr = context.JSON(statusCode, errMessage)
			}

			if respErr != nil {
				context.Echo().Logger.Error(respErr)
			}
		}
	}
}
