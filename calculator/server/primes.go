package main

import (
	pb "github.com/emiliano080591/grpcv2/calculator/proto"
	"log"
)

func (s *Server) Primes(req *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("primes function was invoked with %v\n", req)

	number := req.Number
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: divisor,
			})
			number /= divisor
		} else {
			divisor++
		}
	}
	return nil
}
