package main

import (
	"log"
	"net/http"
)

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

func main() {
	http.HandleFunc("/msg", handleConnections)
	go handleMessages()
	// go handleJoin() TODO
	// go handleLeave() TODO
	// go handleGetRooms() TODO

	log.Fatal(http.ListenAndServe(":8080", nil))
}
