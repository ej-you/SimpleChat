package services

import (
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	echo "github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// кодирование пароля в хэш
func EncodePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// возвращаем 400, потому что скорее всего ошибка длины пароля
		return "", echo.NewHTTPError(400, map[string]string{"encodePassword": err.Error()})
	}
	return string(hash), nil
}

// проверка введённого юзером пароля на совпадение с хэшем из БД
func PasswordIsCorrect(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// получение ID юзера из контекста запроса
func GetUserIDFromRequest(context echo.Context) (uuid.UUID, error) {
	var contextUserID uuid.UUID

	// ошибка, которую возвратит функция при неудаче
	getTokenUserIDError := echo.NewHTTPError(400, map[string]string{"parseToken": "failed to get user id from token"})

	// достаём map значений JWT-токена из контекста context
	token, ok := context.Get("user").(*jwt.Token)
	if !ok {
		return contextUserID, getTokenUserIDError
	}
	tokenClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return contextUserID, getTokenUserIDError
	}

	// приведение значения id юзера к string
	stringContextUserID, ok := tokenClaims["userID"].(string)
	if !ok {
		return contextUserID, getTokenUserIDError
	}
	// парсинг строки с uuid в объект uuid.UUID
	contextUserID, err := uuid.Parse(stringContextUserID)
	if err != nil {
		return contextUserID, echo.NewHTTPError(500, map[string]string{"parseToken": "failed to parse uuid from string"})
	}

	return contextUserID, nil
}
