package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogItem struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string `bson:"author_id,omitempty"`
	Title string `bson:"title,omitempty"`
	Content string `bson:"content,omitempty"`
}
