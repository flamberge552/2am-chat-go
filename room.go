package main

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

// Room json object mapper
type Room struct {
	ID   bson.ObjectId `bson:"_id" json:"id"`
	Name string        `bson:"roomName" json:"roomName"`
}

// CreateRoom generates a mongo collection with a given name
func CreateRoom(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not implemented yet")
}

// RetrieveRooms returns a json object array of the requested rooms
func RetrieveRooms(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	rooms, err := dao.FindAllRooms()
	if err != nil {
		log.Fatal(err)
	}
	RespondWithJSON(w, http.StatusOK, rooms)
}
