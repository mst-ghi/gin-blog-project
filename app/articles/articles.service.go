package articles

import (
	"blog/core"
	"blog/database/models"
	"blog/database/repositories"
)

type ArticlesServiceInterface interface {
	FindAll() []models.Article
	FindOne(id string) (models.Article, core.Error)
}

type ArticlesService struct {
	repository repositories.ArticleRepositoryInterface
}

func NewArticlesService() *ArticlesService {
	return &ArticlesService{
		repository: repositories.NewArticleRepository(),
	}
}

func (self *ArticlesService) FindAll() []models.Article {
	return self.repository.FindAll()
}

func (self *ArticlesService) FindOne(id string) (models.Article, core.Error) {
	article := self.repository.FindByID(id)

	if article.ID == 0 {
		return article, core.Error{"reason": "Article not found"}
	}

	return article, nil
}
