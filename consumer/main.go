package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

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

func main() {
	serverAddr := "ws://exchange:8080/"

	conn, _, err := websocket.DefaultDialer.Dial(serverAddr, nil)
	if err != nil {
		log.Fatal("Erro ao estabelecer conexão WebSocket:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Conectado ao servidor WebSocket.")

	subscribeChannel(conn, "BRL_USD:ticker")
	subscribeChannel(conn, "BRL_USD:ORDERBOOK")
	subscribeChannel(conn, "BRL_USD:TRANSACTIONS")

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Erro ao receber mensagem WebSocket:", err)
			return
		}
		if messageType == websocket.TextMessage {
			fmt.Printf("Mensagem Recebida: %s\n", p)
		}
	}
}
