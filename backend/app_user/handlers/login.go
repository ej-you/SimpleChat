package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

	coreValidator "SimpleChat/backend/core/validator"
	"SimpleChat/backend/app_user/serializers"
	"SimpleChat/backend/core/db"
	"SimpleChat/backend/core/services"
)


func Login(context echo.Context) error {
	var err error
	var dataIn serializers.LoginUserIn

	// парсинг JSON-body
	if err = context.Bind(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = coreValidator.Validate(&dataIn); err != nil {
		return err
	}
	// получение юзера из БД по username'у
	userFromDB, err := db.NewDB().GetUserByUsername(dataIn.Username)
	if err != nil {
		return err
	}
	// проверка на совпадение введённого пароля и хэша из БД
	if ok := services.PasswordIsCorrect(dataIn.Password, userFromDB.Password); !ok {
		return echo.NewHTTPError(400, map[string]string{"password": "Invalid password"})
	}
	// получение куки авторизации
	var newAuthCookie *http.Cookie
	newAuthCookie, err = services.GetAuthCookie(userFromDB.ID)
	if err != nil {
		return err
	}
	// добавление куки авторизации в ответ
	context.SetCookie(newAuthCookie)

	return context.JSON(http.StatusOK, userFromDB)
}
