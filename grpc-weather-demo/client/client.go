package main

import (
	"context"
	"io"
	"log"

	pb "grpc-weather-demo/grpc/weather"// import the generated package

	"google.golang.org/grpc"
)

func main() {
	// Connect to the server
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewWeatherServiceClient(conn)

	// Prepare request
	req := &pb.ForecastRequest{City: "New York"}

	// Call GetForecast and receive a stream
	stream, err := client.GetForecast(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling GetForecast: %v", err)
	}

	// Read messages from the stream
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("End of forecast stream.")
			break
		}
		if err != nil {
			log.Fatalf("Error receiving from stream: %v", err)
		}

		log.Printf("Day: %s | Weather: %s | Temp: %.1f°C",
			resp.Day, resp.Description, resp.Temperature)
	}

	log.Println("Forecast complete.")
}
