package service

import "outofmemory/models"

type Article struct {
	ArticleID    uint32
	Title        string
	Content      string
	Type         string
	State        uint8
	Tags         []uint32
	Category     uint32
	AuthorID     uint32
	AuthorName   string
	AuthorAvatar string
	Page         int
	PerPage      int
}

func (article *Article) GetArticlesByTagID() ([]*models.Article, error) {
	return models.GetArticlesByTagID(article.Tags, article.Page, article.PerPage)
}

func (article *Article) GetArticlesByCategoryID() ([]*models.Article, error) {
	return models.GetArticlesByCategoryID(article.Category, article.Page, article.PerPage)
}

func (article *Article) GetArticlesByAuthorID() ([]*models.Article, error) {
	return models.GetArticlesByAuthorID(article.AuthorID, article.Page, article.PerPage)
}

func (article *Article) GetArticleByID() (*models.Article, error) {
	return models.GetArticleByID(article.ArticleID)
}

func (article *Article) AddArticle() (interface{}, error) {
	data := makeArticleData(article)
	return models.AddArticle(data)
}

func (article *Article) UpdateArticle() (*models.Article, error) {
	data := makeArticleData(article)
	return models.UpdateArticle(data)
}

func (article *Article) DeleteArticle() error {
	data := makeArticleData(article)
	return models.DeleteArticle(data)
}

func makeArticleData(article *Article) map[string]interface{} {
	data := map[string]interface{}{
		"article_id":    article.ArticleID,
		"author_id":     article.AuthorID,
		"author_name":   article.AuthorName,
		"author_avatar": article.AuthorAvatar,
		"title":         article.Title,
		"content":       article.Content,
		"type":          article.Type,
		"state":         article.State,
		"tags":          article.Tags,
		"category":      article.Category,
	}
	return data
}
