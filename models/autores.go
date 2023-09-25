package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Autor struct {
	Id        primitive.ObjectID `bson:"_id" json:"_id"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	Img       string             `bson:"img" json:"img"`
	DesertRef string             `bson:"desertRef" json:"desertRef"`
}

type InsertAutor struct {
	Name      string `bson:"name" json:"name"`
	Email     string `bson:"email" json:"email"`
	Img       string `bson:"img" json:"img"`
	DesertRef string `bson:"desertRef" json:"desertRef"`
}

type UpdateAutor struct {
	Name      string `bson:"name" json:"name"`
	Email     string `bson:"email" json:"email"`
	Image     string `bson:"image" json:"image"`
	DesertRef string `bson:"desertref" json:"desertref"`
}
type DeleteAutor struct {
	Name      string `bson:"name" json:"name"`
	Email     string `bson:"email" json:"email"`
	Image     string `bson:"image" json:"image"`
	DesertRef string `bson:"desertref" json:"desertref"`
}