package main

import (
	"context"
	"log"
	"time"

	pb "go-grpc-demo/helloworld/helloworld"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.NewClient("localhost:3000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreetingMessageClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Kiran"})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
