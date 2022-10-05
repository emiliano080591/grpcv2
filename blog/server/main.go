package main

import (
	"context"
	pb "github.com/emiliano080591/grpcv2/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	addr            = "0.0.0.0:50051"
	mongoUrl        = "mongodb://root:roo@localhost:27050/"
	mongoDatabase   = "blogdb"
	mongoCollection = "blog"
)

var collection *mongo.Collection

type Server struct {
	pb.BlogServiceServer
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUrl))
	if err != nil {
		log.Fatalf("error trying to connect with mongodb->%v\n", err.Error())
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatalf("error with mongodb client->%v\n", err.Error())
	}

	collection = client.Database(mongoDatabase).Collection(mongoCollection)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}
	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()

	pb.RegisterBlogServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
