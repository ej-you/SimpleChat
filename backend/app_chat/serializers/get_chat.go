package serializers

import (
	validate "github.com/gobuffalo/validate/v3"
	"github.com/google/uuid"
)

// входные данные для получения чата для двух юзеров и сообщений из этого чата
type GetChatIn struct {
	// id чата
	ID uuid.UUID `param:"id" myvalid:"required" example:"0aafe1fd-0088-455b-9269-0307aae15bcc"`
}

// дополнительная валидация входных данных
func (self *GetChatIn) IsValid(errors *validate.Errors) {}
