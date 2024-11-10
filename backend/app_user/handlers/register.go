package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

	coreValidator "SimpleChat/backend/core/validator"
	"SimpleChat/backend/app_user/serializers"
	"SimpleChat/backend/core/db"
	"SimpleChat/backend/core/services"
)


func Register(context echo.Context) error {
	var err error
	var dataIn serializers.RegisterUserIn

	// парсинг JSON-body
	if err = context.Bind(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = coreValidator.Validate(&dataIn); err != nil {
		return err
	}
	// создание нового юзера в БД
	createdUser, err := db.NewDB().CreateUser(dataIn.Username, dataIn.Password)
	if err != nil {
		return err
	}
	// получение куки авторизации
	var newAuthCookie *http.Cookie
	newAuthCookie, err = services.GetAuthCookie(createdUser.ID)
	if err != nil {
		return err
	}
	// добавление куки авторизации в ответ
	context.SetCookie(newAuthCookie)

	return context.JSON(http.StatusCreated, createdUser)
}
