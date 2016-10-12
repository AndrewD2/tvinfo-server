package tvinfo

type MacroSeries struct {
	Id          int
	Name        string
	Series      []Series
	Backgrounds []string //images not sure what type to put
	Posters     []string //images not sure what type to put
	Banners     []string //images not sure what type to put
}
