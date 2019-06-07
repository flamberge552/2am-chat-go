package main

import (
	"log"
	"net/http"

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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Configure the WS upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// upgrade GET to WS
	ws, err := upgrader.Upgrade(w, r, nil)
	check(err)

	defer ws.Close()

	// register new client
	clients[ws] = true

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		session <- msg
	}
}
