package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

	coreValidator "SimpleChat/backend/core/validator"
	"SimpleChat/backend/app_chat/serializers"
	"SimpleChat/backend/core/db"
	"SimpleChat/backend/core/db/models"
	"SimpleChat/backend/core/services"
)


// эндпоинт для получения чата
//	@Summary		Get chat
//	@Description	Get chat messages and chat participants by chat uuid in path parameters
//	@Router			/chat/{id} [get]
//	@ID				get-chat
//	@Tags			chat
//	@Accept			plain
//	@Produce		json
//	@Param			id	path		string	true	"Chat uuid"
//	@Success		200	{object}	models.Chat
//	@Failure		401	{object}	errors.GeneralExpiredCredentials401
//	@Failure		403	{object}	errors.ChatUserIsNotParticipant403
//	@Failure		404	{object}	errors.ChatNotFound404
//	@Failure		500	{object}	errors.GeneralInternalError500
func GetChat(context echo.Context) error {
	var err error
	var dataIn serializers.GetChatIn
	var chatFromDB models.Chat

	// парсинг path-параметров 
	if err = context.Bind(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = coreValidator.Validate(&dataIn); err != nil {
		return err
	}
	// получение существующего чата из БД по path-параметру-id
	err = db.NewDB().GetFullChatByID(&chatFromDB, dataIn.ID)
	if err != nil {
		return err
	}

	// получение uuid юзера из контекста запроса
	userUUID, err := services.GetUserIDFromRequest(context)
	if err != nil {
		return err
	}
	// если текущий юзер не состоит в запрашиваемом чате
	if userUUID != chatFromDB.Users[0].ID && userUUID != chatFromDB.Users[1].ID {
		return echo.NewHTTPError(403, map[string]string{"getChat": "forbidden"})
	}

	return context.JSON(http.StatusOK, chatFromDB)
}
