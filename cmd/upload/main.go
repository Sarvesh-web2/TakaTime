package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/Rtarun3606k/TakaTime/internal/db"
	"github.com/Rtarun3606k/TakaTime/internal/types"
)

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

	client, err := db.ConnectToDataBase(*uri)
	if err != nil {
		log.Fatalln("Counld not connect to mongo db", err)
	}

	collection := client.Database("takatime").Collection("logs")

	entry := types.LogEntry{
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
