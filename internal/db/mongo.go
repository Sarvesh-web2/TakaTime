package db

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"time"
)

func ConnectToDataBase(uri string) (*mongo.Client, error) {

	opts := options.Client().ApplyURI(uri)

	// We create a temporary context just for the connection handshake
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var Client *mongo.Client
	var err error
	Client, err = mongo.Connect(opts)
	if err != nil {
		log.Fatal("Error creating client:", err)
		return Client, err
	}

	// Ping is the safest way to check.
	if err := Client.Ping(ctx, nil); err != nil {
		log.Fatal("Could not ping MongoDB:", err)
		return Client, err
	}

	log.Println("✅ Connected to the database successfully!")

	return Client, nil
}
