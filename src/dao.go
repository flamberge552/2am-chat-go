package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	mgo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.Server))
	if err != nil {
		log.Fatal(err)
	}
	db = client.Database(m.Database)
}

// Insert adds 1 message into the DB
func (m *MessagesDAO) Insert(msg Message) error {
	collection := db.Collection(COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, &msg)
	if err != nil {
		log.Fatal(err)
	}

	return err
}

// FindAll returns ALL messages stored
func (m *MessagesDAO) FindAll() ([]*Message, error) {
	var msgs []*Message

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := db.Collection(COLLECTION).Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(ctx) {
		var elem Message
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		msgs = append(msgs, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	defer cur.Close(ctx)

	return msgs, err
}

// Flush clears out the entire DB
func (m *MessagesDAO) Flush() error {
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	err := db.Collection(COLLECTION).Drop(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return err
}
