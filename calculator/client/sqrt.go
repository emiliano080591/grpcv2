package main

import (
	"context"
	pb "github.com/emiliano080591/grpcv2/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Println("doSqrt was invoked")

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Number: n})

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("error message from server: %s\n", e.Message())
			log.Printf("error code from server: %s\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Println("we probably sent a negative number")
				return
			}
		} else {
			log.Fatalf("anon grpc error: %v\n", err)
		}
	}

	log.Printf("Sqrt: %f\n", res.Result)
}
