package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

	coreValidator "SimpleChat/backend/core/validator"
	"SimpleChat/backend/app_chat/serializers"
	"SimpleChat/backend/core/db"
	"SimpleChat/backend/core/services"
)


// эндпоинт для получения id чата для двух юзеров
//	@Summary		Get chat id
//	@Description	Get chat id by username of another chat participant in path parameters
//	@Router			/chat/with/{username} [get]
//	@ID				chat-with
//	@Tags			chat
//	@Accept			plain
//	@Produce		json
//	@Param			username	path		string	true	"Chat participant username"
//	@Success		200			{object}	serializers.WithOut
//	@Failure		400			{object}	errors.ChatWithSameUser400
//	@Failure		401			{object}	errors.GeneralExpiredCredentials401
//	@Failure		404			{object}	errors.GeneralUserNotFound404
//	@Failure		500			{object}	errors.GeneralInternalError500
func With(context echo.Context) error {
	var err error
	var dataIn serializers.WithIn

	// парсинг path-параметров 
	if err = context.Bind(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = coreValidator.Validate(&dataIn); err != nil {
		return err
	}
	// получение uuid собеседника из БД по path-параметру-логину
	secondUserFromDB, err := db.NewDB().GetUserByUsername(dataIn.Username)
	if err != nil {
		return err
	}

	// получение uuid юзера из контекста запроса
	userUuid, err := services.GetUserIDFromRequest(context)
	if err != nil {
		return err
	}
	// если второй юзер является первым
	if userUuid == secondUserFromDB.ID {
		return echo.NewHTTPError(400, map[string]string{"chatWith": "another chat participant cannot be the same user"})
	}

	// получение существующего чата для этих двух юзеров или создание нового, если для них ещё нет чата
	chatForUsers, err := db.NewDB().GetOrCreateChat(userUuid, secondUserFromDB.ID)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, serializers.WithOut{ID: chatForUsers.ID})
}
