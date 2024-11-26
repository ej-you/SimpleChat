package db

import (
	"strings"

	echo "github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"github.com/google/uuid"

	"SimpleChat/backend/core/db/models"
	"SimpleChat/backend/core/services"
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
func (db *DB) CreateUser(user *models.User, username, password string) error {

	// генерим новый uuid для записи
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return echo.NewHTTPError(500, map[string]string{"createUser": "failed to create uuid for user"})
	}
	// делаем из пароля хэш
	passwordHash, err := services.EncodePassword(password)
	if err != nil {
		return err
	}

	(*user).ID = newUUID
	(*user).Username = username
	(*user).Password = passwordHash

	createResult := db.dbConnect.Create(user)
	if err := createResult.Error; err != nil {
		// если юзер с таким юзернеймом уже есть
		if strings.HasSuffix(err.Error(), "(2067)") {
			return echo.NewHTTPError(409, map[string]string{"username": "user with such username already exists"})
		}
		return echo.NewHTTPError(500, map[string]string{"createUser": "failed to create user: " + err.Error()})
	}
	return nil
}

// получение юзера по его ID
func (db *DB) GetUserByID(user *models.User, id uuid.UUID) error {
	selectResult := db.dbConnect.First(user, id)
	if err := selectResult.Error; err != nil {
		// если ошибка в ненахождении записи
		if err.Error() == "record not found" {
			return echo.NewHTTPError(404, map[string]string{"getUser": "user with such id was not found"})
		}
		return echo.NewHTTPError(500, map[string]string{"getUser": "failed to get user by id: " + err.Error()})
	}
	return nil
}

// получение юзера по его username'у
func (db *DB) GetUserByUsername(user *models.User, username string) error {
	selectResult := db.dbConnect.Where("username = ?", username).First(user)
	if err := selectResult.Error; err != nil {
		// если ошибка в ненахождении записи
		if err.Error() == "record not found" {
			return echo.NewHTTPError(404, map[string]string{"getUser": "user with such username was not found"})
		}
		return echo.NewHTTPError(500, map[string]string{"getUser": "failed to get user by username: " + err.Error()})
	}
	return nil
}


// получение чата (с подгрузкой его участников и сообщений) по его ID
func (db *DB) GetFullChatByID(chat *models.Chat, id uuid.UUID) error {
	selectResult := db.dbConnect.Preload("Users").Preload(
		"Messages",
		// добавление сортировки сообщений по времени от старых к новым
		func(db *gorm.DB) *gorm.DB {
			return db.Order("messages.created_at ASC")
		},
	).Preload("Messages.Sender").First(chat, id)

	if err := selectResult.Error; err != nil {
		// если ошибка в ненахождении записи
		if err.Error() == "record not found" {
			return echo.NewHTTPError(404, map[string]string{"getChat": "chat with such id was not found"})
		}
		return echo.NewHTTPError(500, map[string]string{"getChat": "failed to get chat by id: " + err.Error()})
	}
	return nil
}

// получение чата (с подгрузкой его участников) по его ID
func (db *DB) GetChatParticipantsByID(chat *models.Chat, id uuid.UUID) error {
	selectResult := db.dbConnect.Preload("Users").First(chat, id)
	if err := selectResult.Error; err != nil {
		// если ошибка в ненахождении записи
		if err.Error() == "record not found" {
			return echo.NewHTTPError(404, map[string]string{"getChat": "chat with such id was not found"})
		}
		return echo.NewHTTPError(500, map[string]string{"getChat": "failed to get chat by id: " + err.Error()})
	}
	return nil
}

// создание нового чата
func (db *DB) createChat(chat *models.Chat, firstUserFromDB, secondUserFromDB models.User) error {
	// генерим новый uuid для записи
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return echo.NewHTTPError(500, map[string]string{"createChat": "failed to create uuid for chat"})
	}
	// создаём чат с привязкой к нему двух юзеров
	(*chat).ID = newUUID
	(*chat).Users = []models.User{firstUserFromDB, secondUserFromDB}

	createResult := db.dbConnect.Create(chat)
	if err := createResult.Error; err != nil {
		return echo.NewHTTPError(500, map[string]string{"createChat": "failed to create chat"})
	}
	return nil
}

// получение чата или создание нового, если такого ещё нет 
func (db *DB)  GetOrCreateChat(chat *models.Chat, firstUserID, secondUserID uuid.UUID) error {
	var firstUserFromDB, secondUserFromDB models.User

	// получение чатов первого юзера
	selectResult := db.dbConnect.Preload("Chats").First(&firstUserFromDB, firstUserID)
	if err := selectResult.Error; err != nil {
		return echo.NewHTTPError(500, map[string]string{"getUser": "failed to get user by id: " + err.Error()})
	}
	// получение чатов второго юзера
	selectResult = db.dbConnect.Preload("Chats").First(&secondUserFromDB, secondUserID)
	if err := selectResult.Error; err != nil {
		return echo.NewHTTPError(500, map[string]string{"getUser": "failed to get user by id: " + err.Error()})
	}

	// совместные чаты юзеров (в контексте этого проекта он должен быть один)
	joinChats := services.IntersectUserChats(firstUserFromDB.Chats, secondUserFromDB.Chats)
	
	var err error
	// если пересечение не было найдено, то создаём новый чат
	if len(joinChats) == 0 {
		err = db.createChat(chat, firstUserFromDB, secondUserFromDB)
		return err
	// если пересечение было найдено, то возвращаем чат (только его id без подгрузки участников и сообщений)
	} else {
		return err
	}

	return nil
}


// создание нового сообщения
func (db *DB) CreateMessage(message *models.Message, chatID, senderID uuid.UUID, content string) error {
	// генерим новый uuid для записи
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return echo.NewHTTPError(500, map[string]string{"createMessage": "failed to create uuid for message"})
	}

	// создаём новую запись сообщения
	(*message).ID = newUUID
	(*message).ChatID = chatID
	(*message).SenderID = senderID
	(*message).Content = content

	createResult := db.dbConnect.Create(message)
	if err := createResult.Error; err != nil {
		return echo.NewHTTPError(500, map[string]string{"createMessage": "failed to create message: " + err.Error()})
	}

	err = db.GetUserByID(&(*message).Sender, senderID)
	if err != nil {
		return err
	}
	
	return nil
}
