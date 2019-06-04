package main

import (
	"math/rand"
	"net/http"
)

var room = make(chan Message)

func generateID() int64 {
	return rand.Int63()
}

func createRoom(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	check(err)

	defer ws.Close()

	clients[ws] = true

}
