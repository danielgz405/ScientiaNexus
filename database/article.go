package database

import (
	"context"

	"github.com/dg/acordia/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *MongoRepo) InsertArticle(ctx context.Context, article *models.InsertArticle) (insertArticle *models.Article, err error) {
	collection := repo.client.Database("ScientiaNexus").Collection("article")
	result, err := collection.InsertOne(ctx, article)
	if err != nil {
		return nil, err
	}
	oid := result.InsertedID.(primitive.ObjectID)
	insertArticle, err = repo.GetArticleById(ctx, oid.Hex())
	if err != nil {
		return nil, err
	}
	return insertArticle, nil
}

func (repo *MongoRepo) GetArticleById(ctx context.Context, id string) (*models.Article, error) {
	collection := repo.client.Database("ScientiaNexus").Collection("article")
	var article models.Article
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&article)
	if err != nil {
		return nil, err
	}

	return &article, nil
}

func (repo *MongoRepo) ListArticles(ctx context.Context) ([]models.Article, error) {
	collection := repo.client.Database("ScientiaNexus").Collection("article")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var category []models.Article
	if err = cursor.All(ctx, &category); err != nil {
		return nil, err
	}
	return category, nil
}

func (repo *MongoRepo) UpdateArticle(ctx context.Context, data models.UpdateArticle, id string) (*models.Article, error) {
	collection := repo.client.Database("ScientiaNexus").Collection("article")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	update := bson.M{
		"$set": bson.M{},
	}
	iterableData := map[string]interface{}{
		"name":      data.Name,
		"autor":     data.Autor,
		"content":   data.Content,
		"documents": data.Documents,
		"image":     data.Image,
		"desertRef": data.DesertRef,
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
	article, err := repo.GetArticleById(ctx, oid.Hex())
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (repo *MongoRepo) DeleteArticle(ctx context.Context, id string) error {
	collection := repo.client.Database("ScientiaNexus").Collection("article")
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
