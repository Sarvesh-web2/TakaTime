package main

import (
	"fmt"
	"strings"
	"time"

	persnalization "github.com/Rtarun3606k/TakaTime/internal/Persnalization"
	"github.com/Rtarun3606k/TakaTime/internal/Styles"
	"github.com/Rtarun3606k/TakaTime/internal/types"
	"github.com/charmbracelet/lipgloss"
)

func buildActiveTimeBox(dist types.ActivityDistribution, styles Styles.AppStyles, width int) string {
	var b strings.Builder

	persona := persnalization.GetCoderPersona(dist)
	titleStyle := lipgloss.NewStyle().Bold(true).Width(width).Align(lipgloss.Center).MarginBottom(1)
	b.WriteString(titleStyle.Render(fmt.Sprintf("━ %s ━", persona)) + "\n")

	// Calculate the total sum of all hours for the percentage math
	totalTime := dist.Morning + dist.Afternoon + dist.Evening + dist.Night

	drawRow := func(label string, value float64, max float64) string {
		barWidth := width - 26
		if barWidth < 5 {
			barWidth = 5
		}
		if barWidth > 20 {
			barWidth = 20
		}

		// Progress bar math
		percentOfMax := 0.0
		if max > 0 {
			percentOfMax = value / max
		}

		filledCount := int(percentOfMax * float64(barWidth))
		if filledCount > barWidth {
			filledCount = barWidth
		}

		filledBar := styles.Color1.Render(strings.Repeat("█", filledCount))
		emptyBar := styles.SubText.Render(strings.Repeat("░", barWidth-filledCount))

		//  Text math
		percentageOfTotal := 0.0
		if totalTime > 0 {
			percentageOfTotal = (value / totalTime) * 100
		}

		timeStr := styles.ListPercent.Render(fmt.Sprintf("%4.1f%%", percentageOfTotal))
		labelStr := styles.ListLabel.Render(fmt.Sprintf("%-10s", label))

		return fmt.Sprintf("%s | %s%s | %s\n", labelStr, filledBar, emptyBar, timeStr)
	}

	// Group all the rows into a single text block
	var rowsBlock strings.Builder
	rowsBlock.WriteString(drawRow("Morning", dist.Morning, dist.MaxVal))
	rowsBlock.WriteString(drawRow("Afternoon", dist.Afternoon, dist.MaxVal))
	rowsBlock.WriteString(drawRow("Evening", dist.Evening, dist.MaxVal))
	rowsBlock.WriteString(drawRow("Night", dist.Night, dist.MaxVal))

	//center all
	centeredRows := lipgloss.NewStyle().
		Width(width).
		Align(lipgloss.Center).
		Render(rowsBlock.String())

	b.WriteString(centeredRows)

	return styles.Box.Width(width).Render(b.String())
}

// --------------------------------------------------------------------------------------------

// heatmap

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func BuildHeatmapBox(history map[string]float64, styles Styles.AppStyles, width int) string {
	var b strings.Builder

	titleStyle := lipgloss.NewStyle().Bold(true).Width(width).Align(lipgloss.Center).MarginBottom(1)
	b.WriteString(titleStyle.Render("━ 365-Day Contribution Graph ━") + "\n\n")

	// 1. Setup 7 empty rows (Sun through Sat)
	rows := make([]strings.Builder, 7)

	// 2. Calculate dates
	today := time.Now()
	start := today.AddDate(0, 0, -364) // 365 days total (including today)

	// Offset to the nearest Sunday so the grid aligns perfectly
	offset := int(start.Weekday())
	curr := start.AddDate(0, 0, -offset)

	// 3. Loop through every day and build the rows horizontally
	for !curr.After(today) {
		wday := int(curr.Weekday())

		if curr.Before(start) {
			// Invisible padding for days before our 365-day window
			rows[wday].WriteString("  ")
		} else {
			dateStr := curr.Format("2006-01-02")
			val := history[dateStr] // Fetch from your real database!

			// Brightness logic based on hours coded!
			char := styles.SubText.Render("░") // Default (0 hours)
			if val > 0 && val <= 1.0 {
				char = styles.Color1.Render("▒") // Light coding
			} else if val > 1.0 && val <= 3.0 {
				char = styles.Color1.Render("▓") // Solid coding
			} else if val > 3.0 {
				char = styles.Color1.Render("█") // Heavy coding!
			}

			// Add a space after each block so the grid breathes horizontally
			rows[wday].WriteString(char + " ")
		}

		curr = curr.AddDate(0, 0, 1) // Move to next day
	}

	// 4. Build the Month Header (Apr, May, Jun...)
	totalWeeks := 0
	tempWeek := start.AddDate(0, 0, -offset)
	for !tempWeek.After(today) {
		totalWeeks++
		tempWeek = tempWeek.AddDate(0, 0, 7)
	}

	// Create an array of blank spaces for the top row
	headerChars := make([]string, totalWeeks)
	for i := range headerChars {
		headerChars[i] = " "
	}

	currWeek := start.AddDate(0, 0, -offset)
	var lastMonth time.Month = 0

	for i := 0; i < totalWeeks; i++ {
		if currWeek.Month() != lastMonth {
			// It's a new month! Check if we have room to write 3 letters safely
			if i+2 < totalWeeks {
				// Make sure we aren't colliding with the previous month
				if i == 0 || headerChars[i-1] == " " {
					monthStr := currWeek.Format("Jan")
					headerChars[i] = string(monthStr[0])
					headerChars[i+1] = string(monthStr[1])
					headerChars[i+2] = string(monthStr[2])
				}
			}
			lastMonth = currWeek.Month()
		}
		currWeek = currWeek.AddDate(0, 0, 7)
	}

	// Join the array into a single flawless string
	monthHeader := strings.Join(headerChars, "")

	// 5. Assemble the grid
	var gridBuilder strings.Builder

	// Add exactly 7 spaces of offset (4 for the text + 3 for the " | ")
	monthStyle := styles.SubText.Copy().UnsetWidth().Padding(0).Margin(0)
	gridBuilder.WriteString(fmt.Sprintf("              \t\t%s\t\n", monthStyle.Render(monthHeader)))

	// Use pure strings, no spaces to accidentally delete!
	rowLabels := []string{"", "Mon", "", "Wed", "", "Fri", ""}
	labelStyle := styles.ListLabel.Copy().UnsetWidth().Padding(0).Margin(0)

	for i := 0; i < 7; i++ {
		// "Mon" becomes "Mon ". "" becomes "    ". The pipes will be laser straight!
		formattedLabel := fmt.Sprintf("%-4s", rowLabels[i])

		rowString := fmt.Sprintf("%s | %s\n", labelStyle.Render(formattedLabel), rows[i].String())
		gridBuilder.WriteString(rowString)
	}

	// 6. Center the grid (Updated height to 8 to account for the new month row!)
	centeredGrid := lipgloss.Place(width, 8, lipgloss.Center, lipgloss.Center, gridBuilder.String())
	b.WriteString(centeredGrid + "\n")

	// 7. Add the legend at the bottom
	legend := fmt.Sprintf("Less %s %s %s %s More",
		styles.SubText.Render("░"),
		styles.Color1.Render("▒"),
		styles.Color1.Render("▓"),
		styles.Color1.Render("█"))

	legendStyled := lipgloss.NewStyle().Width(width).Align(lipgloss.Center).MarginTop(1).Render(legend)
	b.WriteString(legendStyled)

	return styles.Box.Width(width).Render(b.String())
}
