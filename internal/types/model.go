package types

import "time"

type LogEntry struct {
	FileName  string    `bson:"name"`
	Project   string    `bson:"project"`
	TimeStamp time.Time `bson:"timestamp"`
	Duration  float64   `bson:"duration"`
	Date      string    `bson:"date"`
	Language  string    `bson:"language"`
	Os        string    `bson:"os"`
	GitBranch string    `bson:"gitBranch"`
	Editor    string    `bson:"editor"`
}

type StatItem struct {
	Name       string
	Duration   float64
	Percentage float64
}

// --- Config & Colors ---
const (
	TopN     = 4
	BarWidth = 10
	Reset    = "\033[0m"
	Bold     = "\033[1m"
	Green    = "\033[32m"
	Cyan     = "\033[36m"
	Gray     = "\033[90m"
)

type Stat struct {
	Name     string
	Duration float64
	Percent  float64
}
