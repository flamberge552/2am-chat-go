package main

import mgo "gopkg.in/mgo.v2"

var db *mgo.Database

// DB name
const (
	COLLECTION = "messages"
)

// MessagesDAO stands for Messages Data Access Object. It is used as a query in the DB.
type MessagesDAO struct {
	Server   string
	Database string
}

// Connect is called to establish a mongodb connection
func (m *MessagesDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	check(err)
	db = session.DB(m.Database)
}
