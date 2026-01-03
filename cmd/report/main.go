package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	dbquery "github.com/Rtarun3606k/TakaTime/internal/DBquery"
	gogist "github.com/Rtarun3606k/TakaTime/internal/GoGist"
	"github.com/Rtarun3606k/TakaTime/internal/db"
)

func main() {

	// Flags
	days := flag.Int("days", 0, "Number of past days to include (0 = today)")
	flag.Parse()

	// Connect
	mongoURI := os.Getenv("MONGO_URI")
	// gistID := os.Getenv("GIST_ID")
	gistToken := os.Getenv("GIST_TOKEN")
	targetRepo := os.Getenv("TARGET_REPO")

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

	content := dbquery.GenerateReport(logs)
	log.Println(content)

	if gistToken != "" && targetRepo != "" {
		fmt.Printf("🚀 Updating README for %s...\n", targetRepo)

		err := gogist.UpdateReadMe(gistToken, targetRepo, content)
		if err != nil {
			log.Fatalf("❌ Failed to update README: %v", err)
		}

		fmt.Println("✅ README Updated Successfully!")
	} else {
		fmt.Println("ℹ️ Skipping README update (GIST_TOKEN or TARGET_REPO missing)")
	}

	//
	// updateContentError := gogist.UpdateGist(gistToken, gistID, content)
	// if updateContentError != nil {
	// 	log.Fatalln("Some error occured gogist", updateContentError)
	// 	return
	// }
	//

}
