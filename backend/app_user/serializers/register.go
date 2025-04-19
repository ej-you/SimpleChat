package serializers

import (
	validate "github.com/gobuffalo/validate/v3"
)

// @Description входные данные регистрации юзера
type RegisterUserIn struct {
	// логин юзера
	Username string `json:"username" myvalid:"required" example:"vasya_2007" maxLength:"50"`
	// пароль юзера
	Password string `json:"password" myvalid:"required|min:8|max:50" example:"qwerty123" minLength:"8" maxLength:"50"`
}

// дополнительная валидация входных данных
func (user *RegisterUserIn) IsValid(_ *validate.Errors) {}
