package main

import (
	"context" // Substitua pelo nome do pacote definido em seu arquivo .proto
	"exchange_poc/pb"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client.
	client := pb.NewStreamExampleClient(conn)

	// Replace the code below with your gRPC service calls.
	response, err := client.RequestAction(context.Background(), &pb.ActionRequest{
		Action:    "Buy",
		Market:    "BTC-USD",
		MessageId: 123,
	})
	if err != nil {
		log.Fatalf("Error calling RequestAction: %v", err)
	}
	fmt.Println("Response from RequestAction:", response)

	stream, err := client.SubscribeEvent(context.Background(), &pb.SubscribeEventRequest{
		Channel: "market_updates",
	})
	if err != nil {
		log.Fatalf("Error calling SubscribeEvent: %v", err)
	}

	for {
		data, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error receiving data from SubscribeEvent: %v", err)
			break
		}
		fmt.Printf("Data Received: %s (Channel: %s)\n", data.Data, data.Channel)
	}
}
