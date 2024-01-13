package seeder

import (
	"blog/core/config"
	"blog/database"
)

func Seeder() {
	config.InitializeAndSet()

	database.InitializeAndConnect()
	db := database.Connection()

	UserSeeder(db)
}
