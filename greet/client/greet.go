package main

import (
	"context"
	pb "kalyan/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Printf("doGreet was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "kalyan",
	})

	if err != nil {
		log.Fatalf("could not greet due to : %v \n", err)
	}

	log.Printf("Greeting %s\n", res.Result)
}
