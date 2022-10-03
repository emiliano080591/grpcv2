package main

import (
	"context"
	pb "github.com/emiliano080591/grpcv2/greet/proto"
	"io"
	"log"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	req := &pb.GreetRequest{FirstName: "Emiliano"}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("errow while calling GreetManyTimes: %v\n", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading the stream: %v\n", err)
		}
		log.Printf("GreetManytimes: %s\n", msg.Result)
	}
}
