package main

import (
	"fmt"
	"strings"

	"github.com/Rtarun3606k/TakaTime/internal/Styles"
	"github.com/Rtarun3606k/TakaTime/internal/types"
)

// buildStatsList creates a formatted string block for any list of stats
func buildStatsList(title string, stats []types.ListStats) string {
	// If there is no data, don't render the section at all
	if len(stats) == 0 {
		return ""
	}

	var b strings.Builder

	// 1. Add the section title
	b.WriteString(fmt.Sprintf("--- %s ---\n", title))

	// 2. Loop through the stats and build each line
	for _, stat := range stats {
		// Formatting: 15 chars for label, 10 for value, 1 decimal for percent
		line := fmt.Sprintf("%-15s | %-10s | %.1f%%\n", stat.Label, stat.Value, stat.Percent*100)
		b.WriteString(line)
	}

	// 3. Add a blank line at the bottom for spacing
	b.WriteString("\n")

	return b.String()
}

// Renders the string that actually gets printed to the terminal
func (m Model) View() string {
	s := Styles.TitleStyle.Render(" ⏱️ TakaTime Dashboard ") + "\n\n"

	if m.Loading {
		s += "Fetching your coding stats from MongoDB...\n\n"
	} else {
		// Call the reusable component for each list!
		s += buildStatsList("Languages", m.LanguageListStats)
		s += buildStatsList("Projects", m.ProjectListStats)

		// You can easily add OS and Editor here too
		s += buildStatsList("Operating Systems", m.OsListStats)
		s += buildStatsList("Editors", m.editorListStats)
	}

	s += Styles.TextStyle.Render("Press 'q' to quit.")
	return s
}
