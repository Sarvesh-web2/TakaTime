package main

import (
	"context"
	"flag"
	"log"
	"path/filepath"
	"time"

	utils "github.com/Rtarun3606k/TakaTime/internal/Utils"
	"github.com/Rtarun3606k/TakaTime/internal/db"
	"github.com/Rtarun3606k/TakaTime/internal/debugger"
	"github.com/Rtarun3606k/TakaTime/internal/types"
)

// pes2ug23cs645
func main() {

	uri := flag.String("uri", "", "MongoDB Atlas Connection URI")
	project := flag.String("project", "unknown", "Project Name")
	file := flag.String("file", "", "File Name")
	duration := flag.Float64("duration", 0, "Duration in seconds")
	language := flag.String("language", "unknown", "Lnaguage")
	editor := flag.String("editor", "unknown", "Editor Name NeoVim/VsCode")

	flag.Parse()

	if *uri == "" || *duration <= 0 {
		log.Fatalln("Arguments are empty MongoDB URI or Duration is less than 0")
		return
	}

	err := debugger.SetupLog()

	if err != nil {
		// If setup fails, we print to the standard console so the user sees it
		// We use log.Panic to print the error and crash/exit the app
		log.Panic("Failed to initialize logger: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := db.ConnectToDataBase(*uri)
	if err != nil {
		log.Fatalln("Counld not connect to mongo db", err)
	}

	collection := client.Database("takatime").Collection("logs")

	fileDir := filepath.Dir(*file)

	gitBranch, err := utils.GetGitBranch(fileDir)
	if err != nil {
		log.Printf("Could not get git branch there might not be one or initiated yet !! %s", err)
		gitBranch = ""
	}

	entry := types.LogEntry{
		FileName:  *file,
		Project:   *project,
		Duration:  *duration,
		TimeStamp: time.Now(),
		Date:      time.Now().Format("2006-01-02"),
		Language:  *language,
		Os:        utils.GetOS(),
		GitBranch: gitBranch,
		Editor:    *editor,
	}

	_, err = collection.InsertOne(ctx, entry)

	if err != nil {
		log.Fatal("Insert Failed:", err)
	}

	log.Println("Log processed sucessfullty")
}
