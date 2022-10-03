package main

import (
	"context"
	pb "github.com/emiliano080591/grpcv2/greet/proto"
	"log"
	"time"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Emiliano"},
		{FirstName: "Jose"},
		{FirstName: "Tania"},
		{FirstName: "Xochitl"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("error while calling LongGreet %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("error while receiving response from LongGreet: %v\n", err)
	}
	log.Printf("LongGreet: %s\n", res.Result)
}
