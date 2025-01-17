package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

var room Room
var rooms Rooms

// Room object map
type Room struct {
	RoomID int32 `json:"id"`
}

// Rooms array
type Rooms struct {
	RoomIDs []int32 `json:"ids"`
}

func (rs *Rooms) generateID(r *Room) {
	r.RoomID = rand.Int31()
	rs.RoomIDs = append(rs.RoomIDs, r.RoomID)
}

// CreateRoom generates a room with and ID and returns it
func CreateRoom(w http.ResponseWriter, r *http.Request) {
	rooms.generateID(&room)
	roomJSON, err := json.Marshal(&room)
	if err != nil {
		panic(err)
	}
	w.Write(roomJSON)
}

// ReturnRooms returns all the rooms that have been generated
func ReturnRooms(w http.ResponseWriter, r *http.Request) {
	roomsJSON, err := json.Marshal(&rooms)
	if err != nil {
		panic(err)
	}
	w.Write(roomsJSON)
}
