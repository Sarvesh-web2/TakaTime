package main

import (
	"log"

	dbqueryv2 "github.com/Rtarun3606k/TakaTime/internal/DBQueryV2"
	"github.com/Rtarun3606k/TakaTime/internal/db"
	"github.com/Rtarun3606k/TakaTime/internal/types"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var labels = [4]string{"language", "project", "os", "editor"}

func (m Model) GetData(URI string) (Model, *mongo.Client, error) {
	Client, err := db.ConnectToDataBase(URI)
	if err != nil {
		log.Println("Database connection failed:", err)
		return m, nil, err // Return the unmodified model on error
	}

	for _, value := range labels {
		data, err := dbqueryv2.GetListStats(Client, value, 30, types.CatppuccinTheme, 0)
		if err != nil {
			log.Println("Failed to fetch stats for", value, ":", err)
			return m, nil, err
		}

		// Optional: Keep your debug printing if you want
		// fmt.Println("---" + value + "---")
		// for _, stat := range data {
		// 	fmt.Printf("%-15s | %-10s | %.1f%%\n", stat.Label, stat.Value, stat.Percent*100)
		// }

		// Assign the data directly to the model's fields
		switch value {
		case "language":
			m.LanguageListStats = data
		case "project":
			m.ProjectListStats = data
		case "os":
			m.OsListStats = data
		case "editor":
			m.editorListStats = data
		}
	}

	// Return model
	return m, Client, nil
}
