package main

import (
	"context"
	"log"
	"net"

	pb "go-grpc-demo/helloworld/helloworld"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreetingMessageServer
}

// this method is exposed to the clients
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {

	s := grpc.NewServer() //from grpc native library
	pb.RegisterGreetingMessageServer(s, &server{}) //from generated code
	
	lis, err := net.Listen("tcp", ":3000")         //open the port 3000
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	
	log.Println("Server listening at", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
