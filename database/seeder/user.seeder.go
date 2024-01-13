package seeder

import (
	"blog/database/models"
	"fmt"

	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

func UserSeeder(db *gorm.DB) {
	fmt.Println("User seeder running...")

	for i := 0; i < 5; i++ {
		user := new(models.User)
		user.Name = faker.Name()
		user.Email = faker.Email()
		user.Password = "12345678"

		db.Create(&user)
		fmt.Printf("User created by name:%s, email:%s", user.Name, user.Email)

		ArticleSeeder(db, user.ID)
	}

	fmt.Println("User seeder ended.")
	fmt.Println("")
}
