package main

import (
	"context"
	pb "github.com/emiliano080591/grpcv2/blog/proto"
	"log"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("----createBlog was invoked----")

	blog := &pb.Blog{
		AuthorId: "Emiliano",
		Title:    "My first blog",
		Content:  "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %s\n", res.Id)
	return res.Id
}
