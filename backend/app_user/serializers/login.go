package serializers

import (
	validate "github.com/gobuffalo/validate/v3"
)

// @Description входные данные для входа юзера
type LoginUserIn struct {
	// логин юзера
	Username string `json:"username" myvalid:"required" example:"vasya_2007"`
	// пароль юзера
	Password string `json:"password" myvalid:"required" example:"qwerty123"`
}

// дополнительная валидация входных данных
func (self *LoginUserIn) IsValid(errors *validate.Errors) {}
