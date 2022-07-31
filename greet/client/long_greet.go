package main

import (
	"context"
	pb "kalyan/greet/proto"
	"log"
	"time"
)

func doGreetMultipleTimes(c pb.GreetServiceClient) {
	log.Println("doGreetMultipleTimes invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Kalyan"},
		{FirstName: "is"},
		{FirstName: "aquarian"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatal("Error while calling long greet %v\n", err)
	}
	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)

		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from longgreet %v \n", err)
	}

	log.Printf("LongGreet %v\n ", res.Result)
}
