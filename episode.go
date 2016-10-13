package tvinfo

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
	DvdID          int
	DvdSeason      int
	DvdEpisode     int
	DvdChapter     int
	BdID           int
	BdSeason       int
	BdEpisode      int
	BdChapter      int
	Absolute       int
	ImdbID         string
	IsMovie        bool
}
