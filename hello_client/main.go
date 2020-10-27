package main

import (
	"fmt"
	"github/Ko4s/grpc-intro/hello_client/client")


func main() {

	c := client.NewClient("localhost:50051")
	fmt.Println(c.SayHello("Piotrek"))
	fmt.Println(c.SayHello("Szymon"))
}
