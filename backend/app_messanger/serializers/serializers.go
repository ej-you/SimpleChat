package serializers

import (
	"encoding/json"

	validate "github.com/gobuffalo/validate/v3"
	"github.com/google/uuid"
	echo "github.com/labstack/echo/v4"

	"SimpleChat/backend/core/db"
	"SimpleChat/backend/core/db/models"
	coreValidator "SimpleChat/backend/core/validator"
)

// входные данные для получения чата для двух юзеров и сообщений из этого чата
type GetChatIn struct {
	// id чата
	ParticipantID uuid.UUID `header:"ParticipantID" myvalid:"required" example:"0aafe1fd-0088-455b-9269-0307aae15bcc"`
}

// дополнительная валидация входных данных
func (chat *GetChatIn) IsValid(errors *validate.Errors) {}

// входные данные отправки сообщения через WebSocket
type MessageIn struct {
	ChatID  uuid.UUID `json:"chatId" myvalid:"required" example:"0aafe1fd-0088-455b-9269-0307aae15bcc"`
	Content string    `json:"content" myvalid:"required" example:"sample message"`
}

// дополнительная валидация входных данных
func (mes *MessageIn) IsValid(errors *validate.Errors) {}

// десериализация сырого сообщения в структуру и её валидация
func (mes *MessageIn) ParseAndValidate(rowMessage []byte) error {
	// десериализация сообщения
	err := json.Unmarshal(rowMessage, mes)
	if err != nil {
		return echo.NewHTTPError(400, map[string]string{"message": "failed to parse JSON from message: " + err.Error()})
	}
	// валидация сообщения
	if err = coreValidator.Validate(mes); err != nil {
		return err
	}
	return nil
}

// получение второго участника чата
func GetChatParticipantUUID(chatUUID, firstParticipantUUID uuid.UUID) (uuid.UUID, error) {
	var chat models.Chat
	err := db.NewDB().GetChatParticipantsByID(&chat, chatUUID)
	if err != nil {
		return uuid.UUID{}, err
	}

	// если первый юзер не является участником этого чата, то возвращаем ошибку
	if firstParticipantUUID != chat.Users[0].ID && firstParticipantUUID != chat.Users[1].ID {
		return uuid.UUID{}, echo.NewHTTPError(403, map[string]string{"message": "forbidden"})
	}

	if chat.Users[0].ID == firstParticipantUUID {
		return chat.Users[1].ID, nil
	} else {
		return chat.Users[0].ID, nil
	}
}
