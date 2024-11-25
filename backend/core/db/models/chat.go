package models

import (
	"time"

	"github.com/google/uuid"
)


// модель чата
type Chat struct {
	// uuid чата
	ID			uuid.UUID	`gorm:"not null; type:uuid; primaryKey" json:"id" example:"0aafe1fd-0088-455b-9269-0307aae15bcc"`
	// участники чата
	Users		[]User		`gorm:"many2many:chat_participants" json:"users"`
	// сообщения чата
	Messages 	[]Message 	`gorm:"foreignKey:ChatID" json:"messages"`
}

// модель сообщения
type Message struct {
	// id сообщения
	ID			uuid.UUID	`gorm:"not null; type:uuid; primaryKey; autoIncrement" json:"-"`
	// id чата
	ChatID 		uuid.UUID	`gorm:"not null; type:uuid; index" json:"-"`
	// id отправителя
	SenderID 	uuid.UUID	`gorm:"not null; type:uuid" json:"-"`
	// текст сообщения
	Content		string 		`gorm:"not null; type: LONGTEXT" json:"content" example:"sample message"`
	// время создания сообщения
	CreatedAt	time.Time 	`gorm:"not null; autoCreateTime:true" json:"createdAt" example:"2024-11-13T12:34:56Z"`
	// ассоциация c чатом, к которому принадлежат сообщения
	Chat     	Chat     	`gorm:"foreignKey:ChatID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	// отправитель сообщения
	Sender 		User 		`gorm:"foreignKey:SenderID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"sender"`
}
