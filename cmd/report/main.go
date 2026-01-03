package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	dbquery "github.com/Rtarun3606k/TakaTime/internal/DBquery"
	"github.com/Rtarun3606k/TakaTime/internal/db"
)

func main() {

	// Flags
	days := flag.Int("days", 0, "Number of past days to include (0 = today)")
	flag.Parse()

	// Connect
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable is required")
	}

	client, err := db.ConnectToDataBase(mongoURI)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	// Run Analysis
	logs, err := dbquery.FetchLogs(client, *days)
	if err != nil {
		log.Fatal(err)
	}

	if len(logs) == 0 {
		fmt.Println("No logs found for this period.")
		return
	}

	dbquery.GenerateReport(logs)
}
