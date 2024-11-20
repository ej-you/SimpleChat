package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

	coreValidator "SimpleChat/backend/core/validator"
	"SimpleChat/backend/app_user/serializers"
	"SimpleChat/backend/core/db"
	"SimpleChat/backend/core/services"
)


// эндпоинт для проверки юзера на существование
//	@Summary		Check user is exists
//	@Description	Check user is exists by his username (returns error if checked current user)
//	@Router			/user/check/{username} [get]
//	@ID				user-check
//	@Tags			user
//	@Accept			plain
//	@Produce		json
//	@Param			username	path		string	true	"Check user is exists"
//	@Success		200			{object}	serializers.CheckOut
//	@Failure		400			{object}	errors.UserCheck400
//	@Failure		401			{object}	errors.UserCheck401
//	@Failure		404			{object}	errors.UserCheck404
//	@Failure		409			{object}	errors.UserCheck409
//	@Failure		500			{object}	errors.General500
func Check(context echo.Context) error {
	var err error
	var dataIn serializers.CheckIn

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
		return echo.NewHTTPError(409, map[string]string{"check": "current user was checked"})
	}

	return context.JSON(http.StatusOK, serializers.CheckOut{IsExists: true})
}
