package tvinfo

type Episode struct {
	Id             int
	Name           string
	FirstAired     string
	Guest          []string
	Director       []string
	Writer         []string
	ProductionCode string
	Description    string
	Poster         []string //images not sure what type to put
	DvdId          int
	DvdSeason      int
	DvdEpisode     int
	DvdChapter     int
	BdId           int
	BdSeason       int
	BdEpisode      int
	BdChapter      int
	Absolute       int
	ImdbId         string
	IsMovie        bool
}
