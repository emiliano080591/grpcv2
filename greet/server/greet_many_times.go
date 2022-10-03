package main

import (
	"fmt"
	pb "github.com/emiliano080591/grpcv2/greet/proto"
	"log"
)

func (s *Server) GreetManyTimes(req *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with: %v\n", req)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", req.FirstName, i)
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}
	return nil
}
