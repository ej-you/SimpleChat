// Структуры ошибок для Swagger документации
package errors

import (
	"time"
)

// --------------
// general errors
// --------------

// @Description ошибка валидации входных данных
type GeneralValidateError400 struct {
	Errors     map[string]string `json:"errors" example:"username:username field must not be blank"`
	Path       string            `json:"path" example:"/api/user/login"`
	Status     string            `json:"status" example:"error"`
	StatusCode int               `json:"statusCode" example:"400"`
	Timestamp  time.Time         `json:"timestamp" example:"24-11-11 11:57:28 +03"`
}

// @Description ошибка отсутствия куков (истёк токен и соответственно куки авторизации вместе с ним)
type GeneralExpiredCredentials401 struct {
	Errors     map[string]string `json:"errors" example:"token:missing auth cookie"`
	Path       string            `json:"path" example:"/api/chat/check"`
	Status     string            `json:"status" example:"error"`
	StatusCode int               `json:"statusCode" example:"401"`
	Timestamp  time.Time         `json:"timestamp" example:"24-11-11 11:57:28 +03"`
}

// @Description ошибка ненахождения юзера с таким логином в БД
type GeneralUserNotFound404 struct {
	Errors     map[string]string `json:"errors" example:"getUser:user with such username was not found"`
	Path       string            `json:"path" example:"/api/user/login"`
	Status     string            `json:"status" example:"error"`
	StatusCode int               `json:"statusCode" example:"404"`
	Timestamp  time.Time         `json:"timestamp" example:"24-11-11 11:57:28 +03"`
}

// @Description обычная пятисотка от сервера, которую никто не ждёт
type GeneralInternalError500 struct {
	Errors     map[string]string `json:"errors" example:"unknown:some error desc"`
	Path       string            `json:"path" example:"/api/some/shit"`
	Status     string            `json:"status" example:"error"`
	StatusCode int               `json:"statusCode" example:"500"`
	Timestamp  time.Time         `json:"timestamp" example:"24-11-11 11:57:28 +03"`
}

// ------------------
// /api/user/register
// ------------------

// @Description ошибка регистрации юзера с уже существующим (занятым) логином
type UserRegusterAlreadyExistsError409 struct {
	Errors     map[string]string `json:"errors" example:"username:user with such username already exists"`
	Path       string            `json:"path" example:"/api/user/register"`
	Status     string            `json:"status" example:"error"`
	StatusCode int               `json:"statusCode" example:"409"`
	Timestamp  time.Time         `json:"timestamp" example:"24-11-11 11:57:28 +03"`
}

// ---------------
// /api/user/login
// ---------------

// @Description ошибка неверного пароля
type UserLoginInvalidPassword401 struct {
	Errors     map[string]string `json:"errors" example:"password:invalid password"`
	Path       string            `json:"path" example:"/api/user/login"`
	Status     string            `json:"status" example:"error"`
	StatusCode int               `json:"statusCode" example:"401"`
	Timestamp  time.Time         `json:"timestamp" example:"24-11-11 11:57:28 +03"`
}

// ------------------------
// /api/chat/with/:username
// ------------------------

// @Description ошибка, возникающая при указании второго участника чата как себя
type ChatWithSameUser400 struct {
	Errors     map[string]string `json:"errors" example:"chatWith:another chat participant cannot be the same user"`
	Path       string            `json:"path" example:"/api/chat/with/:username"`
	Status     string            `json:"status" example:"error"`
	StatusCode int               `json:"statusCode" example:"400"`
	Timestamp  time.Time         `json:"timestamp" example:"24-11-11 11:57:28 +03"`
}

// -------------
// /api/chat/:id
// -------------

// @Description ошибка ненахождения чата с таким uuid в БД
type ChatNotFound404 struct {
	Errors     map[string]string `json:"errors" example:"getChat:chat with such id was not found"`
	Path       string            `json:"path" example:"/api/chat/:id"`
	Status     string            `json:"status" example:"error"`
	StatusCode int               `json:"statusCode" example:"404"`
	Timestamp  time.Time         `json:"timestamp" example:"24-11-11 11:57:28 +03"`
}

// @Description ошибка, возникающая при запросе юзером чата, в котором он не состоит
type ChatUserIsNotParticipant403 struct {
	Errors     map[string]string `json:"errors" example:"getChat:forbidden"`
	Path       string            `json:"path" example:"/api/chat/:id"`
	Status     string            `json:"status" example:"error"`
	StatusCode int               `json:"statusCode" example:"409"`
	Timestamp  time.Time         `json:"timestamp" example:"24-11-11 11:57:28 +03"`
}
