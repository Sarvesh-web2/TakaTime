package types

import "time"

type LogEntry struct {
	FileName  string    `bson:"name"`
	Project   string    `bson:"project"`
	TimeStamp time.Time `bson:"timestamp"`
	Duration  float64   `bson:"duration"`
	Date      string    `bson:"date"`
	Language  string    `bson:"language"`
}

type StatItem struct {
	Name       string
	Duration   float64
	Percentage float64
}
