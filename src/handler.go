package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

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

	go keepAlive(ws, 5*time.Second)

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
