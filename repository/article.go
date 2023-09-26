package repository

import (
	"context"

	"github.com/dg/acordia/models"
)

func InsertArticle(ctx context.Context, article *models.InsertArticle) (insertArticle *models.Article, err error) {
	return implementation.InsertArticle(ctx, article)
}
func GetArticleById(ctx context.Context, id string) (*models.Article, error) {
	return implementation.GetArticleById(ctx, id)
}

func ListArticles(ctx context.Context) ([]models.Article, error) {
	return implementation.ListArticles(ctx)
}

func UpdateArticle(ctx context.Context, data models.UpdateArticle, id string) (*models.Article, error) {
	return implementation.UpdateArticle(ctx, data, id)
}

func DeleteArticle(ctx context.Context, id string) error {
	return implementation.DeleteArticle(ctx, id)
}
