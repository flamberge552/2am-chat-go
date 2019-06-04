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
}
