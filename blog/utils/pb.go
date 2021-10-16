package utils

import (
	"github.com/juliocesarscheidt/blog/blogpb"
	"github.com/juliocesarscheidt/blog/model"
)

func DataToBlogPb(data *model.BlogItem) *blogpb.Blog {
	return &blogpb.Blog{
		Id:       data.ID.Hex(),
		AuthorId: data.AuthorID,
		Content:  data.Content,
		Title:    data.Title,
	}
}
