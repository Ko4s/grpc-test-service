package client

import (
	"context"
	"github/Ko4s/grpc-intro/hello"
	"log"

	"google.golang.org/grpc"
)

type Client struct {
	service hello.HelloServiceClient
}

func NewClient(address string) *Client {

	//cc -> Client connection
	cc, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalln(err)
	}

	helloClient := hello.NewHelloServiceClient(cc)

	return &Client{
		service: helloClient,
	}

}

func (c *Client) SayHello(msg string) string {

	req := hello.HelloRequest{
		Message: msg,
	}
	res, err := c.service.SayHello(context.TODO(), &req)

	if err != nil {
		log.Fatalln(err)
	}

	return res.GetMessage()
}
