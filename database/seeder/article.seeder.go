package seeder

import (
	"blog/database/models"
	"fmt"

	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

func ArticleSeeder(db *gorm.DB, userID uint) {
	fmt.Printf("Article seeder running for userID:%d...", userID)

	for i := 0; i < 10; i++ {
		article := new(models.Article)
		article.UserID = userID
		article.Slug = faker.Username()
		article.Title = faker.Sentence()
		article.Content = faker.Paragraph()

		db.Create(&article)
		fmt.Printf("Article created by title:%s", article.Title)
	}

	fmt.Printf("Article seeder ended for userID:%d.", userID)
	fmt.Println("")
}
