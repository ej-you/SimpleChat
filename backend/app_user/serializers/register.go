package serializers

import (
	validate "github.com/gobuffalo/validate/v3"
)

// структура для входных данных регистрации юзера
type RegisterUserIn struct {
	Username 	string `json:"username" myvalid:"required" example:"vasya_2007"`
	Password 	string `json:"password" myvalid:"required|min:8|max:50" example:"qwerty123"` // minLength:"8" maxLength:"50"`
}

// дополнительная валидация входных данных
func (self *RegisterUserIn) IsValid(errors *validate.Errors) {}
