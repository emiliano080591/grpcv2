package main

import (
	"context"
	pb "github.com/emiliano080591/grpcv2/greet/proto"
	"io"
	"log"
	"time"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("error while creating stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Ruben"},
		{FirstName: "Xochitl"},
		{FirstName: "Tania"},
		{FirstName: "Jose"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while receiving: %v\n", err)
				break
			}
			log.Printf("Received: %v\n", res.Result)
		}
		close(waitc)
	}()
	<-waitc
}
