package services

import (
	"net/http"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	echo "github.com/labstack/echo/v4"

	"SimpleChat/backend/settings"
)

// создание токена для юзера
func getToken(userID uuid.UUID) (string, error) {
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(settings.TokenExpiredTime).Unix(),
	})

	tokenString, err := tokenStruct.SignedString([]byte(settings.SecretForJWT))
	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, map[string]string{"obtainToken": err.Error()})
	}

	return tokenString, nil
}

// создание куки авторизации для юзера
func GetAuthCookie(userID uuid.UUID) (*http.Cookie, error) {
	// получение токена для юзера
	token, err := getToken(userID)
	if err != nil {
		return &http.Cookie{}, err
	}
	// создание куки авторизации для всех путей api
	cookie := http.Cookie{
		Name:     "simple-chat-auth",
		Value:    token,
		Path:     "/api/",
		HttpOnly: true,
		Secure:   settings.CookieSecure,
		SameSite: http.SameSiteNoneMode,
		Expires:  time.Now().Add(settings.TokenExpiredTime),
	}

	return &cookie, nil
}
