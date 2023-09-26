package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Article struct {
	Id        primitive.ObjectID `bson:"_id" json:"_id"`
	Name      string             `bson:"name" json:"name"`
	Autor     string             `bson:"autor" json:"autor"`
	Date      string             `bson:"date" json:"date"`
	Content   string             `bson:"content" json:"content"`
	Documents []string           `bson:"documents" json:"documents"`
	Image     string             `bson:"image" json:"image"`
}

type InsertArticle struct {
	Name      string   `bson:"name" json:"name"`
	Autor     string   `bson:"autor" json:"autor"`
	Date      string   `bson:"date" json:"date"`
	Content   string   `bson:"content" json:"content"`
	Documents []string `bson:"documents" json:"documents"`
	Image     string   `bson:"image" json:"image"`
}

type UpdateArticle struct {
	Name      string   `bson:"name" json:"name"`
	Autor     string   `bson:"autor" json:"autor"`
	Content   string   `bson:"content" json:"content"`
	Documents []string `bson:"documents" json:"documents"`
	Image     string   `bson:"image" json:"image"`
}
