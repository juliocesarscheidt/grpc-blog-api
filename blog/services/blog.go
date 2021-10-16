package services

import (
	"fmt"
	"time"
	"context"

	"github.com/juliocesarscheidt/blog/blogpb"
	"github.com/juliocesarscheidt/blog/model"
	"github.com/juliocesarscheidt/blog/utils"
	"github.com/juliocesarscheidt/blog/repository"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Server struct {
	blogpb.UnimplementedBlogServiceServer
	Collection *mongo.Collection
}

func (s *Server) CreateBlog(ctx context.Context, req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {
	fmt.Printf("Request: %v\n", req)

	blog := req.GetBlog()

	data := model.BlogItem{
		AuthorID: blog.GetAuthorId(),
		Title: blog.GetTitle(),
		Content: blog.GetContent(),
	}

	result, err := s.Collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert to OID"),
		)
	}

	return &blogpb.CreateBlogResponse {
		Blog: &blogpb.Blog {
			Id: objectID.Hex(),
			AuthorId: blog.GetAuthorId(),
			Title: blog.GetTitle(),
			Content: blog.GetContent(),
		},
	}, nil
}

func (s *Server) ReadBlog(ctx context.Context, req *blogpb.ReadBlogRequest) (*blogpb.ReadBlogResponse, error) {
	fmt.Printf("Request: %v\n", req)

	data, err := repository.FindBlogItemFromID(ctx, s.Collection, req.GetId())
	if err != nil {
		return nil, err
	}

	return &blogpb.ReadBlogResponse{
		Blog: utils.DataToBlogPb(data),
	}, nil
}

func (s *Server) UpdateBlog(ctx context.Context, req *blogpb.UpdateBlogRequest) (*blogpb.UpdateBlogResponse, error) {
	fmt.Printf("Request: %v\n", req)

	blog := req.GetBlog()

	data, err := repository.FindBlogItemFromID(ctx, s.Collection, blog.GetId())
	if err != nil {
		return nil, err
	}

	// updating internal struct
	if blog.GetAuthorId() != "" {
		data.AuthorID = blog.GetAuthorId()
	}
	if blog.GetTitle() != "" {
		data.Title = blog.GetTitle()
	}
	if blog.GetContent() != "" {
		data.Content = blog.GetContent()
	}

	objectID, err := primitive.ObjectIDFromHex(blog.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}
	filter := primitive.M{"_id": objectID}

	_, updateErr := s.Collection.ReplaceOne(ctx, filter, data)
	if updateErr != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot update object: %v", updateErr),
		)
	}

	return &blogpb.UpdateBlogResponse {
		Blog: utils.DataToBlogPb(data),
	}, nil
}

func (s *Server) DeleteBlog(ctx context.Context, req *blogpb.DeleteBlogRequest) (*blogpb.DeleteBlogResponse, error) {
	fmt.Printf("Request: %v\n", req)

	data, err := repository.FindBlogItemFromID(ctx, s.Collection, req.GetId())
	if err != nil {
		return nil, err
	}

	objectID, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}
	filter := primitive.M{"_id": objectID}

	_, deleteErr := s.Collection.DeleteOne(ctx, filter)
	if deleteErr != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete object: %v", deleteErr),
		)
	}

	return &blogpb.DeleteBlogResponse {
		Id: data.ID.Hex(),
	}, nil
}

func (s *Server) ListBlog(req *blogpb.ListBlogRequest, reqStream blogpb.BlogService_ListBlogServer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := s.Collection.Find(ctx, model.BlogItem{})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		data := &model.BlogItem{}

		err := cursor.Decode(&data)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Internal error: %v", err),
			)
		}

		res := &blogpb.ListBlogResponse {
			Blog: utils.DataToBlogPb(data),
		}
		reqStream.Send(res)
	}

	return nil
}
