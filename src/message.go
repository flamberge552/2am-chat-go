package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/websocket"
)

// Message json object mapper
type Message struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Username string        `bson:"username" json:"username"`
	Body     string        `bson:"body" json:"body"`
	Color    struct {
		R int `bson:"r" json:"r"`
		G int `bson:"g" json:"g"`
		B int `bson:"b" json:"b"`
	}
}

var clients = make(map[*websocket.Conn]bool) // connected clients
var session = make(chan Message)             // broadcast channel

func handleMessages(dao *MessagesDAO) {
	for {
		// fetch the next message from the channel
		msg := <-session
		msg.ID = bson.NewObjectId()
		log.Printf("Incoming message: %v", msg)
		dao.Insert(msg)
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

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func getMessages(w http.ResponseWriter, r *http.Request) {
	msg, err := dao.FindAll()
	if err != nil {
		respondWithError(w, 500, "Error contacting DB")
	}
	respondWithJSON(w, 200, msg)
}

func flushDB(w http.ResponseWriter, r *http.Request) {
	dao.Flush()
	body := "DB Cleared"
	response, _ := json.Marshal(body)
	w.Write(response)
}
