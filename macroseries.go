package tvinfo

// MacroSeries - the basic structure of a MacroSeries.
type MacroSeries struct {
	ID          int
	Name        string
	Series      []Series
	Backgrounds []Image
	Posters     []Image
	Banners     []Image
}
