package Styles

import "github.com/charmbracelet/lipgloss"

var (
	TitleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#7D56F4")). // A nice purple
			BorderStyle(lipgloss.RoundedBorder()).
			Padding(0, 1)
	TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)
