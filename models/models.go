package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID		primitive.ObjectID	`json:"_id,omitempty" bson:"_id,omitempty"`
	Isbn	string				`json:"isbn,omitempty" bson:"isbn,omitempty"`
	Title	string				`json:"title,omitempty" bson:"title,omitempty"`
	Author	*Author				`json:"author,omitempty" bson:"author,omitempty"`
}

type Author struct {
	FirstName 	string 	`json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName 	string 	`json:"lastname,omitempty" bson:"lastname,omitempty"`
}