package main

import (
	"context"
	"encoding/json"
	"exchange_poc/pb"
	exchange "exchange_poc/pb"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
)

type server struct {
	exchange.UnimplementedExchangeSteamServer
}

type SubscribEventRequest struct {
	Channel string `json:"channel"`
}

type DataReceiver struct {
	Channel string `json:"channel"`
	Data    string `json:"data"`
}

type ActionRequest struct {
	Action    string `json:"action"`
	Market    string `json:"market"`
	MessageId string `json:"message_id"`
}

func (s *server) Heathcheck(ctx context.Context, req *pb.HeathCheck) (*pb.HeathCheck, error) {
	return &pb.HeathCheck{Status: "OK"}, nil
}

func (s *server) RequestAction(ctx context.Context, req *pb.ActionRequest) (*pb.DataReceiver, error) {
	conn := connectWebSocket()
	defer conn.Close()

	rand.Seed(time.Now().UnixNano())

	data := map[string]interface{}{
		"action":      req.Action,
		"market":      req.Market,
		"messagei_id": req.MessageId,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	if err := conn.WriteMessage(websocket.TextMessage, jsonData); err != nil {
		log.Println("Erro ao enviar mensagem WebSocket:", err)
		return nil, err
	}

	messageType, p, err := conn.ReadMessage()

	if messageType == websocket.TextMessage {
		return &pb.DataReceiver{Data: string(p), IsSuccess: true}, nil
	}
	return &pb.DataReceiver{Data: "Error", IsSuccess: false}, nil
}

func (s *server) SubscribeEvent(req *exchange.SubscribeEventRequest, stream exchange.ExchangeSteam_SubscribeEventServer) error {
	conn := connectWebSocket()
	defer conn.Close()

	subscribeChannel(conn, req.Channel)

	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				return
			default:
				messageType, p, err := conn.ReadMessage()
				if err != nil {
					log.Println("Erro ao receber mensagem WebSocket:", err)
					// Feche o stream se ocorrer um erro no WebSocket
					stream.Send(&exchange.DataReceiver{Data: "Erro no WebSocket"})
					stream.Context().Done()
					return
				}
				if messageType == websocket.TextMessage {
					response := &exchange.DataReceiver{
						Data:    string(p),
						Channel: "canal_de_resposta",
					}
					if err := stream.Send(response); err != nil {
						log.Printf("Erro ao enviar resposta: %v", err)
						done <- struct{}{} // Sinalize o encerramento da goroutine
						return
					}
					time.Sleep(2 * time.Second)
				}
			}
		}
	}()

	// Aguarde o encerramento do stream gRPC
	<-stream.Context().Done()
	done <- struct{}{} // Sinalize o encerramento da goroutine
	return nil
}

func start(port int) *grpc.Server {
	lis, error := net.Listen("tcp", fmt.Sprintf("go-consumer:%d", port))
	if error != nil {
		log.Fatalf("Não foi possível conectar: %v", error)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterExchangeSteamServer(grpcServer, &server{})

	log.Println("Listening on port :", port)

	grpcServer.Serve(lis)
	return grpcServer
}

func subscribeChannel(conn *websocket.Conn, channel string) error {
	data := map[string]interface{}{
		"action":  "subscribe",
		"channel": channel,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := conn.WriteMessage(websocket.TextMessage, jsonData); err != nil {
		log.Println("Erro ao enviar mensagem WebSocket:", err)
		return err
	}

	return nil
}

func connectWebSocket() *websocket.Conn {
	serverAddr := "ws://exchange:8080/"

	conn, _, err := websocket.DefaultDialer.Dial(serverAddr, nil)
	if err != nil {
		log.Fatal("Erro ao estabelecer conexão WebSocket:", err)
	}
	return conn
}

func main() {
	s := start(50051)
	log.Println(s)
}
