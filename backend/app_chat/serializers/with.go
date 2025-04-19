package serializers

import (
	validate "github.com/gobuffalo/validate/v3"
	"github.com/google/uuid"
)

// входные данные для получения id чата для двух юзеров
type WithIn struct {
	// логин собеседника
	Username string `param:"username" myvalid:"required" example:"boris_2007"`
}

// дополнительная валидация входных данных
func (chatWith *WithIn) IsValid(_ *validate.Errors) {}

// @Description выходные данные получения id чата для двух юзеров
type WithOut struct {
	// uuid чата
	ID uuid.UUID `json:"id" example:"0aafe1fd-0088-455b-9269-0307aae15bcc"`
}
