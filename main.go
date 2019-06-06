package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var port string

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

func main() {
	Block{
		Try: func() {
			if os.Getenv("PORT") == "" {
				fmt.Printf("Detected dev environment, falling back to local port 8080")
				port = "8080"
			}
		},
		Catch: func(e Exception) {
			fmt.Printf("$PORT env var not set, %v\n", e)
		},
		Finally: func() {
			port = os.Getenv("PORT")
		},
	}.Do()

	http.HandleFunc("/createRoom", CreateRoom)
	http.HandleFunc("/getRooms", ReturnRooms)
	http.HandleFunc("/msg", handleConnections)
	go handleMessages()
	// go handleJoin() TODO
	// go handleLeave() TODO
	// go handleGetRooms() TODO

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
