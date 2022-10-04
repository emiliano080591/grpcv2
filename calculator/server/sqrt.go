package main

import (
	"context"
	"fmt"
	pb "github.com/emiliano080591/grpcv2/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math"
)

func (s *Server) Sqrt(ctx context.Context, req *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt was invoked with: %v\n", req)

	number := req.Number

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %d", number),
		)
	}
	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(number))}, nil
}
