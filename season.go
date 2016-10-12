package tvinfo

type Season struct {
	Id        int
	Metatitle string
	Metadesc  string
	Poster    []string //images not sure what type to put
	Episodes  []Episode
}
