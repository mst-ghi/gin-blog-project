package repositories

import (
	"blog/database"
	"blog/database/models"

	"gorm.io/gorm"
)

type ArticleRepositoryInterface interface {
	Connection() *gorm.DB
	Create(article models.Article) models.Article
	FindByID(id string) models.Article
	FindAll() []models.Article
}

type ArticleRepository struct {
	DB *gorm.DB
}

func NewArticleRepository() *ArticleRepository {
	return &ArticleRepository{
		DB: database.Connection(),
	}
}

func (self *ArticleRepository) Connection() *gorm.DB {
	return self.DB
}

func (self *ArticleRepository) Create(article models.Article) models.Article {
	var newArticle models.Article
	self.DB.Create(&article).Scan(&newArticle)
	return newArticle
}

func (self *ArticleRepository) FindByID(id string) models.Article {
	var article models.Article
	self.DB.First(&article, "id = ?", id)
	return article
}

func (self *ArticleRepository) FindAll() []models.Article {
	var articles []models.Article
	self.DB.Table("articles").Find(&articles)
	return articles
}
