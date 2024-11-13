package serializers

import (
	validate "github.com/gobuffalo/validate/v3"
)


// входные данные для получения чата для двух юзеров и сообщений из этого чата
type GetMessagesIn struct {
	// логин собеседника
	Username string `param:"username" myvalid:"required" example:"boris_2007"`
}

// дополнительная валидация входных данных
func (self *GetMessagesIn) IsValid(errors *validate.Errors) {}
