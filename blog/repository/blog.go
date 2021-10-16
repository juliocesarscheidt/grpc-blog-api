package repository

import (
	"fmt"
	"context"

	"github.com/juliocesarscheidt/blog/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindBlogItemFromID(ctx context.Context, collection *mongo.Collection, blogID string) (*model.BlogItem, error) {
	objectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	data := &model.BlogItem{}
	filter := primitive.M{"_id": objectID}

	err = collection.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog with specified ID: %v", err),
		)
	}

	return data, nil
}
