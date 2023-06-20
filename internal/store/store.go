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
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://root:sml12345@localhost:27017/"))
	if err != nil {
		panic(err)
	}

	return &Database{
		PasswordCollection: client.Database("passwordManager").Collection("passwords"),
		NotesCollection:    client.Database("passwordManager").Collection("notes"),
	}
}
