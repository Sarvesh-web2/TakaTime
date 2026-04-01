package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func initModel(mongoURI string) Model {
	return Model{
		Loading:  true,
		Err:      nil,
		MongoURI: mongoURI, // Save this so Init() can use it
	}
}

func (m Model) Init() tea.Cmd {
	// If no URI was provided, don't try to fetch
	if m.MongoURI == "" {
		return nil
	}
	// Start spinning the spinner and tell Bubble Tea to fetch data in the background!
	return fetchData(m.MongoURI)
}

func main() {
	var mongoDBString string
	flag.StringVar(&mongoDBString, "MongoDBString", "", "MongoDB String for dashboard to query data")
	flag.Parse()

	// Initialize the model with the URI. It will start with Loading: true
	appModel := initModel(mongoDBString)

	// Start the program instantly. The Init() function handles the async loading.
	p := tea.NewProgram(appModel)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
