package articles

import (
	"blog/database/models"
	"time"
)

type ResponseType map[string]any

type Article struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"article_id"`
	Slug      string `json:"slug"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ArticleTransform(article models.Article) Article {
	return Article{
		ID:        article.ID,
		UserID:    article.UserID,
		Title:     article.Title,
		Slug:      article.Slug,
		Content:   article.Content,
		CreatedAt: article.CreatedAt.Format(time.RFC3339),
		UpdatedAt: article.UpdatedAt.Format(time.RFC3339),
	}
}

type ArticleResponseType struct {
	Article Article `json:"article"`
}

func ArticleResponse(article models.Article) ResponseType {
	return ResponseType{
		"article": ArticleTransform(article),
	}
}

type ArticlesResponseType struct {
	Articles []Article `json:"articles"`
}

func ArticlesResponse(articles []models.Article) ResponseType {
	var data []Article

	for _, article := range articles {
		data = append(data, ArticleTransform(article))
	}

	return ResponseType{
		"articles": data,
	}
}
