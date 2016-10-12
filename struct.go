package tvinfo

type MacroSeries struct {
	id          int
	name        string
	series      []Series
	backgrounds []string //images not sure what type to put
	posters     []string //images not sure what type to put
	banners     []string //images not sure what type to put
}

type Series struct {
	id          int
	name        string
	desc        string
	firstaired  string
	airtime     string
	runtime     int
	network     string
	genre       []string
	stars       []string
	seasons     []Season
	backgrounds []string //images not sure what type to put
	banners     []string //images not sure what type to put
	poster      []string //images not sure what type to put
}

type Season struct {
	id        int
	metatitle string
	metadesc  string
	power     []string //images not sure what type to put
	episodes  []Episode
}

type Episode struct {
	id          int
	name        string
	first_aired string
	guest       []string
	director    []string
	writer      []string
	prodcode    string
	description string
	poster      []string //images not sure what type to put
	dvd_id      int
	dvd_season  int
	dvd_episode int
	dvd_chapter int
	bd_id       int
	bd_season   int
	bd_episode  int
	bd_chapter  int
	absolute    int
	imdb_id     string
	is_movie    bool
}

func init() {

}
