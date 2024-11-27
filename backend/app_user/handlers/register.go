package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

	coreValidator "SimpleChat/backend/core/validator"
	"SimpleChat/backend/app_user/serializers"
	"SimpleChat/backend/core/db"
	"SimpleChat/backend/core/db/models"
	"SimpleChat/backend/core/services"
)


// эндпоинт для регистрации юзера
//	@Summary		Register user
//	@Description	Register new user with form
//	@Router			/user/register [post]
//	@ID				user-register
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			RegisterUserIn	body		serializers.RegisterUserIn	true	"Register params"
//	@Success		201				{object}	models.User
//	@Failure		400				{object}	errors.GeneralValidateError400
//	@Failure		409				{object}	errors.UserRegusterAlreadyExistsError409
//	@Failure		500				{object}	errors.GeneralInternalError500
func Register(context echo.Context) error {
	var err error
	var dataIn serializers.RegisterUserIn
	var newUser models.User

	// парсинг JSON-body
	if err = context.Bind(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = coreValidator.Validate(&dataIn); err != nil {
		return err
	}
	// создание нового юзера в БД
	err = db.NewDB().CreateUser(&newUser, dataIn.Username, dataIn.Password)
	if err != nil {
		return err
	}
	// получение куки авторизации
	var newAuthCookie *http.Cookie
	newAuthCookie, err = services.GetAuthCookie(newUser.ID)
	if err != nil {
		return err
	}
	// добавление куки авторизации в ответ
	context.SetCookie(newAuthCookie)

	return context.JSON(http.StatusCreated, newUser)
}
