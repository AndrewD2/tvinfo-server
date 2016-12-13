package models

// Episode - the basic structure of an episode.
type Episode struct {
	ID             int
	Name           string
	FirstAired     string
	Guest          []string
	Director       []string
	Writer         []string
	ProductionCode string
	Description    string
	Poster         []Image
	DVDID          int
	DVDSeason      int
	DVDEpisode     int
	DVDChapter     int
	BDID           int
	BDSeason       int
	BDEpisode      int
	BDChapter      int
	Absolute       int
	IMDBID         string
	IsMovie        bool
}
