package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"
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

// CreateMessage will save 1 message into the database
func CreateMessage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var message Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	message.ID = bson.NewObjectId()
	if err := dao.Insert(message); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	RespondWithJSON(w, http.StatusCreated, message.ID)
}

// RetrieveMessages returns a json object array of the requested rooms
func RetrieveMessages(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not implemented")
}

// RetrieveAllMessages returns a json object array of all meggaes in a room
func RetrieveAllMessages(w http.ResponseWriter, r *http.Request) {
	messages, err := dao.FindAllMessages()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, messages)
}
