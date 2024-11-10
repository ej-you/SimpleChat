package models

import (
	"github.com/google/uuid"
)


// модель юзера
type User struct {
	// поля БД
	// ID			uuid.UUID 	`gorm:"not null; primaryKey" json:"id"`
	ID			uuid.UUID 	`gorm:"not null; type:uuid; primaryKey" json:"id"`
	Username	string 		`gorm:"not null; type: VARCHAR(50); unique" json:"username"`
	Password	string 		`gorm:"not null; type: VARCHAR(255)" json:"-"`
	// ассоциация юзера с чатами, в которых он состоит
	Chats 		[]Chat 	`gorm:"many2many:chat_participants" json:"-"`
	// ассоциация сообщений, которые отправил юзер
	// Messages 	[]Message 	`gorm:"foreignKey:SenderID" json:"-"`
	// Chats 		[]Chat 		`gorm:"foreignKey:FirstUserID" json:"-"`
}
