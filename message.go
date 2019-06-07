package main

import (
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

// Insert saves the message captured in the broadcast channel into the DB.
// will only return if there is an error
func (m *MessagesDAO) Insert(message Message) error {
	err := db.C(COLLECTION).Insert(&message)
	return err
}

// FindAll returns all messages saved in the DB
// will only return error if there is one thrown
func (m *MessagesDAO) FindAll() ([]Message, error) {
	var messages []Message
	err := db.C(COLLECTION).Find(bson.M{}).All(&messages)
	return messages, err
}
