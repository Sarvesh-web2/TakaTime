package dbquery

import (
	"context"
	"fmt"
	"math"
	"sort"
	"strings"
	"time"

	"github.com/Rtarun3606k/TakaTime/internal/types"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// --- MODULE 1: FETCH DATA ---
func FetchLogs(client *mongo.Client, days int) ([]types.LogEntry, error) {
	collection := client.Database("takatime").Collection("logs")
	ctx := context.TODO()

	// Calculate date range (Today minus 'days')
	startDate := time.Now().AddDate(0, 0, -days).Format("2006-01-02")

	// Filter: Date >= startDate
	filter := bson.M{"date": bson.M{"$gte": startDate}}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var logs []types.LogEntry
	if err = cursor.All(ctx, &logs); err != nil {
		return nil, err
	}
	return logs, nil
}

// --- MODULE 2: CALCULATE TOTAL TIME ---
func GetTotalDuration(logs []types.LogEntry) float64 {
	var total float64
	for _, log := range logs {
		total += log.Duration
	}
	return total
}

// --- MODULE 3: LANGUAGE STATS (With %) ---
func GetLanguageStats(logs []types.LogEntry, totalTime float64) []types.StatItem {
	// 1. Map to sum durations
	statsMap := make(map[string]float64)
	for _, log := range logs {
		statsMap[log.Language] += log.Duration
	}

	// 2. Convert to Slice
	var stats []types.StatItem
	for name, duration := range statsMap {
		percentage := (duration / totalTime) * 100
		stats = append(stats, types.StatItem{
			Name:       name,
			Duration:   duration,
			Percentage: percentage,
		})
	}

	// 3. Sort (High to Low)
	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Duration > stats[j].Duration
	})

	return stats
}

// --- MODULE 4: PROJECT STATS ---
func GetProjectStats(logs []types.LogEntry, totalTime float64) []types.StatItem {
	statsMap := make(map[string]float64)
	for _, log := range logs {
		statsMap[log.Project] += log.Duration
	}

	var stats []types.StatItem
	for name, duration := range statsMap {
		percentage := (duration / totalTime) * 100
		stats = append(stats, types.StatItem{
			Name:       name,
			Duration:   duration,
			Percentage: percentage,
		})
	}

	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Duration > stats[j].Duration
	})

	return stats
}

// --- HELPER: FORMAT TIME ---
func formatDuration(seconds float64) string {
	d := time.Duration(seconds) * time.Second
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	return fmt.Sprintf("%dh %02dm", h, m)
}

// --- MODULE 5: GENERATE REPORT (The Output) ---
func GenerateReport(logs []types.LogEntry) string {
	var result strings.Builder
	totalSeconds := GetTotalDuration(logs)
	langStats := GetLanguageStats(logs, totalSeconds)
	projStats := GetProjectStats(logs, totalSeconds)

	result.WriteString("\n📊 TakaTime Report (Last 24h)\n")
	result.WriteString("========================================")
	result.WriteString(fmt.Sprintf("⏱️  Total Coding Time: %s\n", formatDuration(totalSeconds)))
	result.WriteString("----------------------------------------")

	result.WriteString("\n📂 Projects:")
	for _, item := range projStats {
		bar := generateProgressBar(item.Percentage)
		line := fmt.Sprintf(" %-15s %6s  %-6.1f%% %s\n",
			item.Name,
			formatDuration(item.Duration),
			item.Percentage,
			bar,
		)
		result.WriteString(line)
	}

	result.WriteString("\n💻 Languages:")
	for _, item := range langStats {
		bar := generateProgressBar(item.Percentage)
		line := fmt.Sprintf(" %-15s %6s  %-6.1f%% %s\n",
			item.Name,
			formatDuration(item.Duration),
			item.Percentage,
			bar,
		)
		result.WriteString(line)
	}
	result.WriteString("========================================")
	return result.String()
}

// Optional: Cool ASCII Progress Bar
func generateProgressBar(percentage float64) string {
	width := 10
	filled := int(math.Round((percentage / 100) * float64(width)))
	bar := ""
	for i := 0; i < width; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}
	return bar
}
