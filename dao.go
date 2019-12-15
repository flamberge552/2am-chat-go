package main

import (
	"log"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
)

// DAO is the database access object containing the database address and database name
type DAO struct {
	Server   string
	Database string
}

var db *mgo.Database

//COLLECTION is the equivalent of the room where the messages will be temporarily stored
const COLLECTION = "messages"

//Connect is a simple dial method to connect to the mongo instance
func (dao *DAO) Connect() {
	session, err := mgo.Dial(dao.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(dao.Database)
}

//Insert adds 1 message into the collection
func (dao *DAO) Insert(message Message) error {
	err := db.C(COLLECTION).Insert(&message)
	return err
}

//FindAllMessages returns all messages found within the implied collection
func (dao *DAO) FindAllMessages() ([]Message, error) {
	var messages []Message
	err := db.C(COLLECTION).Find(bson.M{}).All(&messages)
	return messages, err
}

//FindAllRooms returns all collections
func (dao *DAO) FindAllRooms() ([]string, error) {
	rooms, err := db.CollectionNames()
	if err != nil {
		log.Fatal(err)
	}
	return rooms, err
}
