package main

import (
	"fmt"
	pb "kalyan/greet/proto"
	"log"
)

func (s *Server) GreetServerStreaming(in *pb.GreetRequest, stream pb.GreetService_GreetServerStreamingServer) error {
	log.Println("GreetServerStreaming got invoked", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, Number %d", in.FirstName, i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
