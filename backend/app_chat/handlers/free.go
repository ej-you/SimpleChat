package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)


// эндпоинт для тестирования запроса без авторизации
//	@Summary		Free endpoint for test cookie auth
//	@Description	Does not require auth cookie 
//	@Router			/chat/free [get]
//	@ID				chat-free
//	@Tags			chat
//	@Success		200	{string}	string	"Free endpoint"
//	@Failure		500	{object}	errors.General500
func Free(context echo.Context) error {
	return context.JSON(http.StatusOK, "Free endpoint")
}
