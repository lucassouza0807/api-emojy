package migration

import (
	connection "api-emoji/config"
	"api-emoji/internal/models"
	"fmt"
)

func RunMigration() {
	db, err := connection.GetDatabaseConnection()

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(
		&models.User{},
		&models.Phrase{},
	)

	sqlDB, err := db.DB()

	if err == nil {
		defer sqlDB.Close()
	}

	fmt.Println("Migration executada com sucesso!")
}
