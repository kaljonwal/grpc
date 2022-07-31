package main

import (
	"context"
	"io"
	pb "kalyan/greet/proto"
	"log"
)

func doGreeting(c pb.GreetServiceClient) {
	log.Println("doGreeting invoked")

	req := &pb.GreetRequest{
		FirstName: "Kalyan",
	}

	stream, err := c.GreetServerStreaming(context.Background(), req)

	if err != nil {
		log.Fatalf("error while invoking server %v \n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			log.Fatal("End of file")
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream %v \n")
		}

		log.Printf("GreetManyTime : %v \n", msg)
	}
}
