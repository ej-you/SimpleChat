package models

import (
	"github.com/google/uuid"
)

// модель юзера
// @Description выходные данные входа и регистрации юзера
type User struct {
	// uuid юзера
	ID uuid.UUID `gorm:"not null; type:uuid; primaryKey" json:"id" example:"e2f53f31-0598-4e36-b25d-41bd665764d1"`
	// логин юзера
	Username string `gorm:"not null; type: VARCHAR(50); unique" json:"username" example:"vasya_2007"`
	// хэш пароля юзера
	Password string `gorm:"not null; type: VARCHAR(255)" json:"-"`
	// ассоциация юзера с чатами, в которых он состоит
	Chats []Chat `gorm:"many2many:chat_participants" json:"-"`
}
