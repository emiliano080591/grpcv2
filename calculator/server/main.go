package main

import (
	pb "github.com/emiliano080591/grpcv2/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	addr = "0.0.0.0:50051"
)

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}
	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(s, &Server{})
	reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
