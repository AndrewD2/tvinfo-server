package tvinfo

type Series struct {
	Id          int
	Name        string
	Description string
	FirstAired  string
	AirTime     string
	RunTime     int
	Network     string
	Genre       []string
	Stars       []string
	Seasons     []Season
	Backgrounds []string //images not sure what type to put
	Banners     []string //images not sure what type to put
	Poster      []string //images not sure what type to put
}
