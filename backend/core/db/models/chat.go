package models

import (
	"time"

	"github.com/google/uuid"
)


// модель чата
type Chat struct {
	ID			uuid.UUID	`gorm:"not null; type:uuid; primaryKey" json:"id"`
	// FirstUserID 	uuid.UUID	`gorm:"type:uuid; uniqueIndex:idx_unique_chat_participants" json:"-"`
	// SecondUserID 	uuid.UUID	`gorm:"type:uuid; uniqueIndex:idx_unique_chat_participants" json:"-"`
	// // ассоциация первого участника чата
	// FirstUser 		User 		`gorm:"foreignKey:FirstUserID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"firstUser"`
	// // ассоциация второго участника чата
	// SecondUser 		User 		`gorm:"foreignKey:SecondUserID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"secondUser"`
	// ассоциация участников чата
	Users		[]User		`gorm:"many2many:chat_participants" json:"users"`
	// ассоциация сообщений
	Messages 	[]Message 	`gorm:"foreignKey:ChatID" json:"messages"`
}

// модель сообщения
type Message struct {
	ID			uuid.UUID	`gorm:"not null; type:uuid; primaryKey; autoIncrement" json:"-"`
	ChatID 		uuid.UUID	`gorm:"not null; type:uuid; index" json:"-"`
	SenderID 	uuid.UUID	`gorm:"not null; type:uuid" json:"-"`
	Content		string 		`gorm:"not null; type: LONGTEXT" json:"content"`
	CreatedAt	time.Time 	`gorm:"not null; autoCreateTime:true" json:"createdAt"`
	// ассоциация c чатом, к которому принадлежат сообщения
	Chat     	Chat     	`gorm:"foreignKey:ChatID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	// ассоциация отправителя сообщения
	Sender 		User 		`gorm:"foreignKey:SenderID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"sender"`
}
