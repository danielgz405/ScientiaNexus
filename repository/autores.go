package repository

import (
	"context"

	"github.com/dg/acordia/models"
)

func InsertAutor(ctx context.Context, autor *models.InsertAutor) (insertAutor *models.Autor, err error) {
	return implementation.InsertAutor(ctx, autor)
}
func GetAutorById(ctx context.Context, id string) (*models.Autor, error) {
	return implementation.GetAutorById(ctx, id)
}
func GetAutorByEmail(ctx context.Context, email string) (*models.Autor, error) {
	return implementation.GetAutorByEmail(ctx, email)
}
func UpdateAutor(ctx context.Context, data models.UpdateAutor, id string) (*models.Autor, error) {
	return implementation.UpdateAutor(ctx, data, id)
}
func DeleteAutor(ctx context.Context, id string) error {
	return implementation.DeleteAutor(ctx, id)
}

func ListAutor(ctx context.Context) ([]models.Autor, error) {
	return implementation.ListAutor(ctx)
}
