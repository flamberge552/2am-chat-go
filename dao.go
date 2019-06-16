package main

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MessagesDAO contains the db connection data
type MessagesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

// COLLECTION is the mongodb collection name
const (
	COLLECTION = "messages"
)

// Connect opens the connection to the DB
func (m *MessagesDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Insert adds 1 message into the DB
func (m *MessagesDAO) Insert(msg Message) error {
	err := db.C(COLLECTION).Insert(&msg)
	return err
}

// FindAll returns ALL messages stored
func (m *MessagesDAO) FindAll() ([]Message, error) {
	var msg []Message
	err := db.C(COLLECTION).Find(bson.M{}).All(&msg)
	return msg, err
}
