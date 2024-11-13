package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

	coreValidator "SimpleChat/backend/core/validator"
	"SimpleChat/backend/app_chat/serializers"
	"SimpleChat/backend/core/db"
	"SimpleChat/backend/core/services"
)


// эндпоинт для получения чата для двух юзеров и сообщений из этого чата
//	@Summary		Get chat messages
//	@Description	Get chat messages (also chat uuid && chat participants) by username of another chat participant in path parameters
//	@Router			/chat/get-messages/{username} [get]
//	@ID				chat-get-messages
//	@Tags			chat
//	@Accept			plain
//	@Produce		json
//	@Param			username	path		string	true	"Get messages params"
//	@Success		200			{object}	models.Chat
//	@Failure		400			{object}	errors.ChatGetMessages400
//	@Failure		401			{object}	errors.ChatGetMessages401
//	@Failure		404			{object}	errors.ChatGetMessages404
//	@Failure		500			{object}	errors.General500
func GetMessages(context echo.Context) error {
	var err error
	var dataIn serializers.GetMessagesIn

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
		return echo.NewHTTPError(400, map[string]string{"getMessages": "another chat participant cannot be the same user"})
	}

	// получение существующего чата для этих двух юзеров или создание нового, если для них ещё нет чата
	chatForUsers, err := db.NewDB().GetOrCreateChat(userUuid, secondUserFromDB.ID)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, chatForUsers)
}
