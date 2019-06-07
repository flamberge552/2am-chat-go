package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func currentContext() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Printf("$PORT either undeclared or empty, falling back to default port %s", port)
	}
	return ":" + port
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/createRoom", CreateRoom)
	r.HandleFunc("/getRooms", ReturnRooms)
	r.HandleFunc("/msg", handleConnections)

	go handleMessages()

	defer log.Fatal(http.ListenAndServe(currentContext(), r))
}
