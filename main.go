package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var port = ":8080"

var dao = DAO{}
var config = Config{}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := mux.NewRouter()

	r.Path("/message").Queries("room", "{room:(?:[^\\s]+)}").HandlerFunc(CreateMessage).Methods("POST")
	r.HandleFunc("/messages", RetrieveAllMessages).Methods("GET")

	r.Path("/room").Queries("roomName", "{roomName:(?:[^\\s]+)}").HandlerFunc(CreateRoom).Methods("POST")
	r.HandleFunc("/rooms", RetrieveRooms).Methods("GET")

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
