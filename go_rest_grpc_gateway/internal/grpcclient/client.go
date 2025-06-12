package grpcclient

import (
    "log"
    "google.golang.org/grpc"
    "go_rest_grpc_gateway/pb"
)

func NewUserServiceClient() pb.UserServiceClient {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to connect to gRPC server: %v", err)
    }
    return pb.NewUserServiceClient(conn)
}