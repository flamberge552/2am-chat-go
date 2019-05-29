package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var session = make(chan Message)

// Configure the WS upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Message json object mapper
type Message struct {
	Username string `json:"username"`
	Body     string `json:"body"`
	Color    struct {
		R int `json:"r"`
		G int `json:"g"`
		B int `json:"b"`
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
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

func serveStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/msg", handleConnections)
	http.HandleFunc("/", serveStatic)
	go handleMessages()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
