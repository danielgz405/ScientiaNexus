package database

import (
	"context"

	"github.com/dg/acordia/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *MongoRepo) InsertAutor(ctx context.Context, autor *models.InsertAutor) (insertAutor *models.Autor, err error) {
	collection := repo.client.Database("ScientiaNexus").Collection("autors")
	result, err := collection.InsertOne(ctx, autor)
	if err != nil {
		return nil, err
	}
	oid := result.InsertedID.(primitive.ObjectID)
	insertAutor, err = repo.GetAutorById(ctx, oid.Hex())
	if err != nil {
		return nil, err
	}
	return insertAutor, nil
}
func (repo *MongoRepo) GetAutorById(ctx context.Context, id string) (*models.Autor, error) {
	collection := repo.client.Database("ScientiaNexus").Collection("autors")
	var autor models.Autor
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	// Find one and populate company
	err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&autor)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	// Populate profile

	return &autor, nil
}
