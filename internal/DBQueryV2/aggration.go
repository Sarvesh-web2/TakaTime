package dbqueryv2

import (
	"context"
	"fmt"
	"time"

	"github.com/Rtarun3606k/TakaTime/internal/types"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// Helper struct for unmarshalling aggregation results
type StatResult struct {
	Name         string  `bson:"_id"`
	TotalSeconds float64 `bson:"totalSeconds"`
}

// 1. GENERIC STATS FETCHER (Projects, Languages, OS, Editors)
// fieldName: "project", "language", "os", or "editor"
func GetListStats(client *mongo.Client, fieldName string, limit int, theme types.ThemeConfig) ([]types.ListStats, error) {
	collection := client.Database("takatime").Collection("logs")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pipeline := mongo.Pipeline{
		// 1. Filter out empty/missing fields (Handles Legacy Data!)
		{{Key: "$match", Value: bson.D{
			{Key: fieldName, Value: bson.D{{Key: "$exists", Value: true}, {Key: "$ne", Value: ""}}},
		}}},
		// 2. Group by field and sum duration
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$" + fieldName},
			{Key: "totalSeconds", Value: bson.D{{Key: "$sum", Value: "$duration"}}},
		}}},
		// 3. Sort by usage (High to Low)
		{{Key: "$sort", Value: bson.D{{Key: "totalSeconds", Value: -1}}}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	var results []StatResult
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	// 4. Calculate Total for Percentages
	var grandTotal float64
	for _, r := range results {
		grandTotal += r.TotalSeconds
	}

	// 5. Convert to ListStats struct
	var stats []types.ListStats
	colors := []string{theme.Color1, theme.Color2, theme.Color3, theme.Color4, theme.TextColor}

	for i, r := range results {
		if i >= limit {
			break // Only top N
		}

		// Cycle through colors
		color := colors[i%len(colors)]

		// Calculate Percent
		percent := 0.0
		if grandTotal > 0 {
			percent = r.TotalSeconds / grandTotal
		}

		stats = append(stats, types.ListStats{
			Label:   r.Name, // e.g., "Go" or "Neovim"
			Value:   formatDuration(r.TotalSeconds),
			Percent: percent,
			Color:   color,
		})
	}
	return stats, nil
}

// 2. TIME GRID FETCHER (Uses $facet for efficiency)
func GetTimeStats(client *mongo.Client) (types.TimeGridStruct, error) {
	collection := client.Database("takatime").Collection("logs")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Calculate Timestamps
	now := time.Now()
	yesterdayStart := now.AddDate(0, 0, -1).Truncate(24 * time.Hour) // Midnight yesterday
	yesterdayEnd := yesterdayStart.Add(24 * time.Hour)
	weekAgo := now.AddDate(0, 0, -7)
	monthAgo := now.AddDate(0, 0, -30)

	// FACET PIPELINE: Runs 4 queries in parallel
	pipeline := mongo.Pipeline{
		{{Key: "$facet", Value: bson.D{
			// A. Yesterday
			{Key: "yesterday", Value: bson.A{
				bson.D{{Key: "$match", Value: bson.D{
					{Key: "timestamp", Value: bson.D{{Key: "$gte", Value: yesterdayStart}, {Key: "$lt", Value: yesterdayEnd}}},
				}}},
				bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: nil}, {Key: "total", Value: bson.D{{Key: "$sum", Value: "$duration"}}}}}},
			}},
			// B. Week
			{Key: "week", Value: bson.A{
				bson.D{{Key: "$match", Value: bson.D{{Key: "timestamp", Value: bson.D{{Key: "$gte", Value: weekAgo}}}}}},
				bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: nil}, {Key: "total", Value: bson.D{{Key: "$sum", Value: "$duration"}}}}}},
			}},
			// C. Month
			{Key: "month", Value: bson.A{
				bson.D{{Key: "$match", Value: bson.D{{Key: "timestamp", Value: bson.D{{Key: "$gte", Value: monthAgo}}}}}},
				bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: nil}, {Key: "total", Value: bson.D{{Key: "$sum", Value: "$duration"}}}}}},
			}},
			// D. All Time
			{Key: "allTime", Value: bson.A{
				bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: nil}, {Key: "total", Value: bson.D{{Key: "$sum", Value: "$duration"}}}}}},
			}},
		}}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return types.TimeGridStruct{}, err
	}

	// Unmarshal Facet Result (It's a bit nested)
	var facetResult []struct {
		Yesterday []struct {
			Total float64 `bson:"total"`
		} `bson:"yesterday"`
		Week []struct {
			Total float64 `bson:"total"`
		} `bson:"week"`
		Month []struct {
			Total float64 `bson:"total"`
		} `bson:"month"`
		AllTime []struct {
			Total float64 `bson:"total"`
		} `bson:"allTime"`
	}

	if err = cursor.All(ctx, &facetResult); err != nil {
		return types.TimeGridStruct{}, err
	}

	// Helper to safely get value from slice
	getVal := func(res []struct {
		Total float64 `bson:"total"`
	}) float64 {
		if len(res) > 0 {
			return res[0].Total
		}
		return 0
	}

	// Build Result
	res := facetResult[0]
	return types.TimeGridStruct{
		Yestarday: formatDuration(getVal(res.Yesterday)),
		Week:      formatDuration(getVal(res.Week)),
		Month:     formatDuration(getVal(res.Month)),
		AllTime:   formatDuration(getVal(res.AllTime)),
	}, nil
}

// Helper: 3661s -> "1h 1m"
func formatDuration(seconds float64) string {
	d := time.Duration(seconds) * time.Second
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	if h > 0 {
		return fmt.Sprintf("%dh %dm", h, m)
	}
	return fmt.Sprintf("%dm", m)
}
