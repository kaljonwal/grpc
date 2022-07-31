package main

import (
	"context"
	pb "kalyan/greet/proto"
	"log"
)

func (s *Server) Greet(context context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Println("greet function was invoked with ", in)
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
