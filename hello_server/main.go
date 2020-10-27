package main

import (
	"fmt"
	"github/Ko4s/grpc-intro/hello"
	"github/Ko4s/grpc-intro/hello_server/server"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Server is running")

	localhost := "localhost:50051"
	lis, err := net.Listen("tcp", localhost)

	if err != nil {
		log.Fatalf("%v", err)
	}

	s := grpc.NewServer()
	helloService := server.NewService()

	hello.RegisterHelloServiceServer(s, helloService)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Bump sad face :(")
	}
}
