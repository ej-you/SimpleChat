package serializers

import (
	validate "github.com/gobuffalo/validate/v3"
)

// структура для входных данных регистрации юзера
type LoginUserIn struct {
	Username 	string `json:"username" myvalid:"required" example:"vasya_2007"`
	Password 	string `json:"password" myvalid:"required" example:"qwerty123"`
}

// дополнительная валидация входных данных
func (self *LoginUserIn) IsValid(errors *validate.Errors) {}
