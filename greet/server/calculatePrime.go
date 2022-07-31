package main

import (
	pb "kalyan/greet/proto"
	"log"
)

func (s *Server) CalculatePrimeNumbers(in *pb.PrimeRequest, stream pb.GreetService_CalculatePrimeNumbersServer) error {
	log.Printf("CalculatePrimeNumber invoked %v \n", in)

	k := int64(2)
	N := in.Number1

	for N > 1 {
		if N%k == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: k,
			})

			N = N / k
		} else {
			k = k + 1
		}
	}

	return nil
}
