package models

// Season - the basic structure of a season.
type Season struct {
	ID        int
	Metatitle string
	Metadesc  string
	Poster    []Image
	Episodes  []Episode
}
