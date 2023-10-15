package main

import (
	// Substitua pelo nome do pacote definido em seu arquivo .proto

	"fmt"
	"log"

	"google.golang.org/grpc"
)

func start(url string, port int) *grpc.ClientConn {
	uri := fmt.Sprintf("%s:%d", url, port)
	conn, err := grpc.Dial(uri, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Não foi possível conectar: %v", err)
	}
	log.Println("Servidor GRPC inciado!")
	return conn
}

func main() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
}
