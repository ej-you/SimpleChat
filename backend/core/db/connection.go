package db

import (
	"fmt"

	"gorm.io/gorm"
	_ "modernc.org/sqlite"
	sqlite "github.com/genc-murat/gorm-sqlite-cgo-free"

	"SimpleChat/backend/core/db/models"
	"SimpleChat/backend/settings"
)


// соединение с БД
var dbConnection *gorm.DB = func() *gorm.DB {
	connection, err := gorm.Open(sqlite.Open(settings.PathDB), &gorm.Config{})
	settings.DieIf(err)
	return connection
}()


// создание таблиц в БД по структурам в Go
func Migrate() {
	fmt.Println("Start migration...")
	
	fmt.Println("Migrate \"User\" model...")
	err := dbConnection.AutoMigrate(&models.User{})
	settings.DieIf(err)

	fmt.Println("Migrate \"Chat\" model...")
	err = dbConnection.AutoMigrate(&models.Chat{})
	settings.DieIf(err)
	
	fmt.Println("Migrate \"Message\" model...")
	err = dbConnection.AutoMigrate(&models.Message{})
	settings.DieIf(err)

	fmt.Println("DB -- Migrated successfully!")
}
