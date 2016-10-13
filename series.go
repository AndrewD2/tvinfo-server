package tvinfo

// Series - the basic structure of a series.
type Series struct {
	ID          int
	Name        string
	Description string
	FirstAired  string
	AirTime     string
	RunTime     int
	Network     string
	Genre       []string
	Stars       []string
	Seasons     []Season
	Backgrounds []Image
	Banners     []Image
	Poster      []Image
}
