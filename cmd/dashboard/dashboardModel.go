package main

import "github.com/Rtarun3606k/TakaTime/internal/types"

type Model struct {
	Loading bool
	Err     error
	//mongo uri

	MongoURI string

	//data model
	LanguageListStats []types.ListStats
	ProjectListStats  []types.ListStats
	OsListStats       []types.ListStats
	editorListStats   []types.ListStats
}
