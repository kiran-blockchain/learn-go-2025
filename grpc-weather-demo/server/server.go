package main

import (
	"log"
	"net"
	"time"

	"grpc-weather-demo/grpc/weather"
	pb "grpc-weather-demo/grpc/weather"

	"google.golang.org/grpc"
)

type  server struct {
	pb.UnimplementedWeatherServiceServer
}

func (s *server) GetForecast(req *weather.ForecastRequest, 
	stream weather.WeatherService_GetForecastServer) error {
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	for _, day := range days {
		stream.Send(&weather.ForecastReply{
			Day:         day,
			Description: "Sunny",
			Temperature: 25.0,
		})
		time.Sleep(5 * time.Second)
	}
	return nil
}

// // this method is exposed to the clients
// func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
// 	log.Printf("Received: %v", in.GetName())
// 	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
// }

func main() {

	s := grpc.NewServer() //from grpc native library
	pb.RegisterWeatherServiceServer(s, &server{}) //from generated code
	
	lis, err := net.Listen("tcp", ":3000")         //open the port 3000
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	
	log.Println("Server listening at", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
