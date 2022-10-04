package main

import (
	"context"
	pb "github.com/emiliano080591/grpcv2/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{FirstName: "Emiliano"}

	res, err := c.GreetWithDeadline(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("deadline exceeded")
				return
			} else {
				log.Fatalf("unexpected grpc error: %v\n", err)
			}
		} else {
			log.Fatalf("a non grpc error: %v\n", err)
		}
	}
	log.Printf("GreetWithDeadline: %s\n", res.Result)
}
