package repository

import (
	"context"

	"github.com/dg/acordia/models"
)

func InsertAutor(ctx context.Context, autor *models.InsertAutor) (insertAutor *models.Autor, err error) {
	return implementation.InsertAutor(ctx, autor)
}
