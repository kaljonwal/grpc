package main

import (
	"context"
	"io"
	pb "kalyan/greet/proto"
	"log"
)

func doCalaculatePrime(c pb.GreetServiceClient) {
	log.Println("doCalaculatePrime invoked")
	req := &pb.PrimeRequest{
		Number1: 120,
	}

	stream, err := c.CalculatePrimeNumbers(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while invoking the server")
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			log.Fatal("EOF received")
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream %v \n", err)
		}

		log.Println("Prime numbers are : %v \n", msg)
	}
}
