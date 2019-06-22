package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var config Config

// var config = Config{}
var dao = MessagesDAO{}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/msg", handleConnections)
	r.HandleFunc("/getAllMessages", getMessages)
	go handleMessages(&dao)
	log.Fatal(http.ListenAndServe(currentContext(), r))
}
