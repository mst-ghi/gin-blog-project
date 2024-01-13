package migrations

import (
	"blog/core/config"
	"blog/database"
	"blog/database/models"
	"fmt"
	"log"
)

func Migrate() {
	config.InitializeAndSet()

	database.InitializeAndConnect()
	db := database.Connection()

	err := db.AutoMigrate(
		&models.User{},
		&models.Token{},
		&models.Article{},
	)

	if err != nil {
		log.Fatal("Cannot migrate")
		return
	}

	fmt.Println("Migration done ..")
}
