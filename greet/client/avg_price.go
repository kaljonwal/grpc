package main

import (
	"context"
	pb "kalyan/greet/proto"
	"log"
	"time"
)

func doAvg(c pb.GreetServiceClient) {
	log.Println("doAvg invoked \n")

	reqs := []*pb.PrimeRequest{
		{Number1: 10},
		{Number1: 20},
		{Number1: 30},
	}

	stream, err := c.AvgRpc(context.Background())
	if err != nil {
		log.Fatal("Error while calling avgRpc %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req to AvgRpc %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from longgreet %v \n", err)
	}

	log.Printf("Avg %v\n ", res.Result)
}
