// Структуры ошибок для Swagger документации
package errors

import (
	"time"
)


// --------------
// general errors
// --------------

// @Description обычная пятисотка от сервера, которую никто не ждёт
type General500 struct {
	Errors 		map[string]string `json:"" example:"unknown: some error desc"`
	Path 		string `json:"" example:"/api/smth/shit"`
	Status 		string `json:"" example:"error"`
	StatusCode 	int `json:"" example:"500"`
	Timestamp	time.Time `json:"" example:"24-11-11 11:57:28 +03"`
}

// ------------------
// /api/user/register
// ------------------

// @Description ошибка валидации входных данных
type UserRegister400 struct {
	Errors 		map[string]string `json:"" example:"password: password field must not be blank"`
	Path 		string `json:"" example:"/api/user/register"`
	Status 		string `json:"" example:"error"`
	StatusCode 	int `json:"" example:"400"`
	Timestamp	time.Time `json:"" example:"24-11-11 11:57:28 +03"`
}

// @Description ошибка регистрации юзера с уже существующим (занятым) логином
type UserRegister409 struct {
	Errors 		map[string]string `json:"" example:"username: user with such username already exists"`
	Path 		string `json:"" example:"/api/user/register"`
	Status 		string `json:"" example:"error"`
	StatusCode 	int `json:"" example:"409"`
	Timestamp	time.Time `json:"" example:"24-11-11 11:57:28 +03"`
}

// ---------------
// /api/user/login
// ---------------

// @Description ошибка валидации входных данных
type UserLogin400 struct {
	Errors 		map[string]string `json:"" example:"username: username field must not be blank"`
	Path 		string `json:"" example:"/api/user/login"`
	Status 		string `json:"" example:"error"`
	StatusCode 	int `json:"" example:"400"`
	Timestamp	time.Time `json:"" example:"24-11-11 11:57:28 +03"`
}

// @Description ошибка неверного пароля
type UserLogin401 struct {
	Errors 		map[string]string `json:"" example:"password: invalid password"`
	Path 		string `json:"" example:"/api/user/login"`
	Status 		string `json:"" example:"error"`
	StatusCode 	int `json:"" example:"401"`
	Timestamp	time.Time `json:"" example:"24-11-11 11:57:28 +03"`
}

// @Description ошибка ненахождения юзера с таким логином в БД
type UserLogin404 struct {
	Errors 		map[string]string `json:"" example:"getUser: user with such username was not found"`
	Path 		string `json:"" example:"/api/user/login"`
	Status 		string `json:"" example:"error"`
	StatusCode 	int `json:"" example:"404"`
	Timestamp	time.Time `json:"" example:"24-11-11 11:57:28 +03"`
}

// --------------------
// /api/chat/restricted
// --------------------

// @Description ошибка отсутствия куков (истёк токен и соответственно куки авторизации вместе с ним)
type ChatRestricted401 struct {
	Errors 		map[string]string `json:"" example:"token: missing auth cookie"`
	Path 		string `json:"" example:"/api/chat/restricted"`
	Status 		string `json:"" example:"error"`
	StatusCode 	int `json:"" example:"401"`
	Timestamp	time.Time `json:"" example:"24-11-11 11:57:28 +03"`
}
