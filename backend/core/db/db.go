package db

import (
	"strings"

	echo "github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"github.com/google/uuid"

	"SimpleChat/backend/core/services"
	"SimpleChat/backend/core/db/models"
)


// структура для запросов к БД
type DB struct {
	dbConnect *gorm.DB
}
func NewDB() *DB {
	return &DB{
		dbConnect: dbConnection,
	}
}


// создание нового юзера
func (db *DB) CreateUser(username, password string) (models.User, error) {
	// генерим новый uuid для записи
	newUuid, err := uuid.NewRandom()
	if err != nil {
		return models.User{}, echo.NewHTTPError(500, map[string]string{"user": "failed to create uuid for user"})
	}
	// делаем из пароля хэш
	passwordHash, err := services.EncodePassword(password)
	if err != nil {
		return models.User{}, err
	}

	newUser := models.User{
		ID: newUuid,
		Username: username,
		Password: passwordHash,
	}

	createResult := db.dbConnect.Create(&newUser)
	if err := createResult.Error; err != nil {
		// если юзер с таким юзернеймом уже есть
		if strings.HasSuffix(err.Error(), "(2067)") {
			return models.User{}, echo.NewHTTPError(409, map[string]string{"username": "user with such username already exists"})
		}
		return models.User{}, echo.NewHTTPError(500, map[string]string{"createUser": "failed to create user: " + err.Error()})
	}
	return newUser, nil
}

// получение юзера по его ID
func (db *DB) GetUserByID(id uuid.UUID) (models.User, error) {
	var userFromDB models.User

	selectResult := db.dbConnect.First(&userFromDB, id)
	if err := selectResult.Error; err != nil {
		// если ошибка в ненахождении записи
		if err.Error() == "record not found" {
			return models.User{}, echo.NewHTTPError(404, map[string]string{"getUser": "user with such id was not found"})
		}
		return models.User{}, echo.NewHTTPError(500, map[string]string{"getUser": "failed to get user by id: " + err.Error()})
	}
	return userFromDB, nil
}

// получение юзера по его username'у
func (db *DB) GetUserByUsername(username string) (models.User, error) {
	var userFromDB models.User

	selectResult := db.dbConnect.Where("username = ?", username).First(&userFromDB)
	if err := selectResult.Error; err != nil {
		// если ошибка в ненахождении записи
		if err.Error() == "record not found" {
			return models.User{}, echo.NewHTTPError(404, map[string]string{"getUser": "user with such username was not found"})
		}
		return models.User{}, echo.NewHTTPError(500, map[string]string{"getUser": "failed to get user by username: " + err.Error()})
	}
	return userFromDB, nil
}


// получение чата или создание нового, если такого ещё нет 
// func (db *DB) GetOrCreateChat(firstUserID, secondUserID uuid.UUID) (models.Chat, error) {
	// var chat models.Chat

	// // попытка получения чата
	// selectResult := db.dbConnect.Where("first_user_id = ? AND second_user_id = ?", firstUserID, secondUserID).Preload("FirstUser").Preload("SecondUser").Preload("Messages").First(&chat)
	// if err := selectResult.Error; err == nil { // NOT error
	// 	return chat, nil
	// }
	// // если произошла ошибка, не связанная с ненахождением записи в БД
	// if selectResult.Error.Error() != "record not found" {
	// 	return models.Chat{}, selectResult.Error
	// }

	// // генерим новый uuid для записи
	// newUuid, err := uuid.NewRandom()
	// if err != nil {
	// 	return models.Chat{}, err
	// }
	// // если не был найден чат, то создаём его
	// chat = models.Chat{
	// 	ID: newUuid,
	// 	FirstUserID: firstUserID,
	// 	SecondUserID: secondUserID,
	// }
	// createResult := db.dbConnect.Create(&chat)
	// if err := createResult.Error; err != nil {
	// 	return models.Chat{}, err
	// }
	// // получение только что созданного чата со всеми связанными объектами
	// selectResult = db.dbConnect.Where(&chat).Preload("FirstUser").Preload("SecondUser").Preload("Messages").First(&chat)
	// if err := selectResult.Error; err != nil {
	// 	return models.Chat{}, err
	// }
	// return chat, nil
// }
