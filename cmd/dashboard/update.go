package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type dataLoadedMsg struct {
	updatedModel Model
	err          error
}

func fetchData(uri string) tea.Cmd {
	return func() tea.Msg {
		// Create a temporary model just to call your fetch method
		tempModel := Model{}

		// Run your heavy DB query
		filledModel, _, err := tempModel.GetData(uri)

		// Return the result as a message to the Update function
		return dataLoadedMsg{
			updatedModel: filledModel,
			err:          err,
		}
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// 1. Catch our background DB result!
	case dataLoadedMsg:
		m.Loading = false // Turn off the loading screen

		if msg.err != nil {
			m.Err = msg.err
			return m, nil
		}

		// Transfer the fetched data into our active model
		m.LanguageListStats = msg.updatedModel.LanguageListStats
		m.ProjectListStats = msg.updatedModel.ProjectListStats
		m.OsListStats = msg.updatedModel.OsListStats
		m.editorListStats = msg.updatedModel.editorListStats

		return m, nil

	// 2. Handle normal keypresses
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		}
	}

	return m, nil
}
