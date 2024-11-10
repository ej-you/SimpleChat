package services

import (
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
