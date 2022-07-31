package main

import (
	"io"
	pb "kalyan/greet/proto"
	"log"
)

func (s *Server) AvgRpc(stream pb.GreetService_AvgRpcServer) error {
	log.Println("AvgRpc got invoked")
	var count int64 = 0
	var sum int64 = 0
	for {

		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.PrimeResponse{
				Result: sum / count,
			})
		}
		count++
		if err != nil {
			log.Fatalf("error while reading stream %v\n", err)
		}
		sum += req.Number1
	}
}
