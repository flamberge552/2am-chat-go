package main

import (
	"log"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var session = make(chan Message)             // broadcast channel

func handleMessages() {
	for {
		// fetch the next message from the channel
		msg := <-session
		// send the message to every currently connected client
		for client := range clients {
			err := client.WriteJSON(&msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
