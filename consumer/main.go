package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	serverAddr := "ws://exchange:8080/"

	conn, _, err := websocket.DefaultDialer.Dial(serverAddr, nil)
	if err != nil {
		log.Fatal("Erro ao estabelecer conexão WebSocket:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Conectado ao servidor WebSocket.")

	data := map[string]interface{}{
		"action":  "subscribe",
		"channel": "BRL_USD:ticker",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("Erro ao serializar JSON:", err)
		return
	}

	if err := conn.WriteMessage(websocket.TextMessage, jsonData); err != nil {
		log.Println("Erro ao enviar mensagem WebSocket:", err)
		return
	}

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
