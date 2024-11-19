package serializers

import (
	validate "github.com/gobuffalo/validate/v3"
)


// входные данные для проверки юзера на существование
type CheckIn struct {
	// логин собеседника
	Username string `param:"username" myvalid:"required" example:"boris_2007"`
}

// дополнительная валидация входных данных
func (self *CheckIn) IsValid(errors *validate.Errors) {}

// @Description выходные данные для проверки юзера на существование
type CheckOut struct {
	// подтверждение существования такого юзера
	IsExists bool `json:"isExists" example:"true"`
}
