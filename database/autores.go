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

	return &autor, nil
}
func (repo *MongoRepo) GetAutorByEmail(ctx context.Context, email string) (*models.Autor, error) {
	collection := repo.client.Database("ScientiaNexus").Collection("autors")
	var autor models.Autor
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&autor)
	if err != nil {
		return nil, err
	}
	return &autor, nil
}
func (repo *MongoRepo) UpdateAutor(ctx context.Context, data models.UpdateAutor, id string) (*models.Autor, error) {
	collection := repo.client.Database("ScientiaNexus").Collection("autors")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	update := bson.M{
		"$set": bson.M{},
	}
	iterableData := map[string]interface{}{
		"name":      data.Name,
		"email":     data.Email,
		"image":     data.Image,
		"desertref": data.DesertRef,
	}
	for key, value := range iterableData {
		if value != "" {
			update["$set"].(bson.M)[key] = value
		}
	}
	err = collection.FindOneAndUpdate(ctx, bson.M{"_id": oid}, update).Err()
	if err != nil {
		return nil, err
	}
	autor, err := repo.GetAutorById(ctx, oid.Hex())
	if err != nil {
		return nil, err
	}
	return autor, nil
}
func (repo *MongoRepo) DeleteAutor(ctx context.Context, id string) error {
	collection := repo.client.Database("ScientiaNexus").Collection("autors")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	return nil
}
