package main

import (
	pb "kalyan/greet/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	pb.GreetServiceServer
}


var addr string = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on %v \n", err)
	}

	log.Printf("listning on %s \n", addr)

	s := grpc.NewServer()

	pb.RegisterGreetServiceServer(s, &Server{})

	

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v \n ", err)
	}

}
