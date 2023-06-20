package store

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	PasswordCollection *mongo.Collection
	NotesCollection    *mongo.Collection
}

func Connect() *Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://root:sml12345@mongo:27017/"))
	if err != nil {
		panic(err)
	}

	return &Database{
		PasswordCollection: client.Database("myDatabase").Collection("passwords"),
		NotesCollection:    client.Database("myDatabase").Collection("notes"),
	}
}
