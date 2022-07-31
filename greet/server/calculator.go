package main

import (
	"context"
	pb "kalyan/greet/proto"
	"log"
)

func (s *Server) SumRpc(ctx context.Context, input *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("SumRpc Got called")

	return &pb.CalculatorResponse{
		Result: input.Number1 + input.Number2,
	}, nil
}
