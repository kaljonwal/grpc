package main

import (
	"context"
	"io"
	pb "kalyan/greet/proto"
	"log"
	"time"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone Invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Kalyan"},
		{FirstName: "Amol"},
		{FirstName: "Avinash"},
		{FirstName: "Jayesh"},
	}

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while creating client stream %v \n", err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request : %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving msg %v\n", err)
			}

			log.Println("result received from server : ", res)
		}
		close(waitc)
	}()

	<-waitc
}
