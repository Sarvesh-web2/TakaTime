package main

import (
	"context"
	"flag"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type LogEntry struct {
	FileName  string    `bson:"name"`
	Project   string    `bson:"project"`
	TimeStamp time.Time `bson:"timestamp"`
	Duration  float64   `bson:"duration"`
	Date      string    `bson:"date"`
	Language  string    `bson:"language"`
}

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

func main() {

	uri := flag.String("uri", "", "MongoDB Atlas Connection URI")
	project := flag.String("project", "unknown", "Project Name")
	file := flag.String("file", "", "File Name")
	duration := flag.Float64("duration", 0, "Duration in seconds")
	language := flag.String("language", "unknown", "Lnaguage")

	flag.Parse()

	if *uri == "" || *duration <= 0 {
		log.Fatalln("Arguments are empty MongoDB URI or Duration is less than 0")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := ConnectToDataBase(*uri)
	if err != nil {
		log.Fatalln("Counld not connect to mongo db", err)
	}

	collection := client.Database("takatime").Collection("logs")

	entry := LogEntry{
		FileName:  *file,
		Project:   *project,
		Duration:  *duration,
		TimeStamp: time.Now(),
		Date:      time.Now().Format("2006-01-02"),
		Language:  *language,
	}

	_, err = collection.InsertOne(ctx, entry)

	if err != nil {
		log.Fatal("Insert Failed:", err)
	}

	log.Println("Log processed sucessfullty")
}
