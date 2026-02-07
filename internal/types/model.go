package types

import (
	"database/sql"
	"time"
)

var DB *sql.DB

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

// LanguageStat holds the data for a single row
type LanguageStat struct {
	Name    string
	Time    string
	Percent float64 // 0.0 to 1.0
	Color   string  // Hex code
}

// Global configuration constants for easy tweaking
const (
	CanvasWidth  = 800
	CanvasHeight = 500
	BgColor      = "#0d1117" // GitHub Dark Dimmed
	TextColor    = "#ffffff"
	SubTextColor = "#8b949e"
	BarBgColor   = "#30363d"
)

type UploadStruct struct {
	Token     string
	Owner     string
	Repo      string
	Path      string // e.g. "public/stats.png"
	Branch    string // e.g. "main"
	CommitMsg string
}

type ListStats struct {
	Label   string
	Value   string
	Percent float64
	Color   string
}

type TimeGridStruct struct {
	Yestarday string
	Week      string
	Month     string
	AllTime   string
}

type ThemeConfig struct {
	BackgroundColor    string
	TextColor          string
	SubTextColor       string
	BarBackgroundColor string

	// Palette for graphs/bars
	Color1 string // Primary (Cyan/Blue)
	Color2 string // Secondary (Green)
	Color3 string // Tertiary (Yellow/Orange)
	Color4 string // Quaternary (Red/Purple)
}
func DefaultTheme() ThemeConfig {
	return ThemeConfig{
		BackgroundColor:    "#0d1117", // GitHub Dark Dimmed
		TextColor:          "#ffffff",
		SubTextColor:       "#8b949e",
		BarBackgroundColor: "#30363d",
		Color1:             "#58a6ff", // Blue
		Color2:             "#2ea043", // Green
		Color3:             "#e3b341", // Gold
		Color4:             "#f78166", // Red
	}
}
