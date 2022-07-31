package main

import (
	"context"
	pb "kalyan/greet/proto"
	"log"
)

func doCalculations(c pb.GreetServiceClient) {
	log.Println("inside doCalculation ")

	res, err := c.SumRpc(context.Background(), &pb.CalculatorRequest{
		Number1: 64,
		Number2: 5,
	})

	if err != nil {
		log.Fatalf("not able to do calculation because of %v \n", err)
	}

	log.Printf("num1 + num2 = %v", res.Result)

}
