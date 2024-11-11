package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

	"SimpleChat/backend/core/db"
	"SimpleChat/backend/core/services"
)


// эндпоинт для тестирования запроса с авторизацией через куки
//	@Summary		Restricted endpoint for test cookie auth
//	@Description	Requires auth cookie
//	@Router			/chat/restricted [get]
//	@ID				chat-restricted
//	@Tags			chat
//	@Produce		json
//	@Success		200	{object}	models.User
//	@Failure		401	{object}	errors.ChatRestricted401
//	@Failure		500	{object}	errors.General500
func Restricted(context echo.Context) error {
	userUuid, err := services.GetUserIDFromRequest(context)
	if err != nil {
		return err
	}

	userFromDB, err := db.NewDB().GetUserByID(userUuid)
	if err != nil {
		return err
	}	

	return context.JSON(http.StatusOK, userFromDB)
}
