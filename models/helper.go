package models

import "github.com/jinzhu/gorm"

type DBReset interface {
	DestructiveReset()
	AutoMigrate()
}

type Models struct {
	*UserGorm
	*ImageGorm
	*PeopleGorm
	*RoleGorm
	*GenreGorm
	*NetworkGorm
	*MacroSeriesGorm
	*SeriesGorm
	*SeasonGorm
	*EpisodeGorm
}

func AllModels(connectionInfo string) (*Models, *gorm.DB, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	return &Models{
		UserGorm:        NewUserGorm(db),
		ImageGorm:       NewImageGorm(db),
		PeopleGorm:      NewPeopleGorm(db),
		RoleGorm:        NewRoleGorm(db),
		GenreGorm:       NewGenreGorm(db),
		NetworkGorm:     NewNetworkGorm(db),
		MacroSeriesGorm: NewMacroSeriesGorm(db),
		SeriesGorm:      NewSeriesGorm(db),
		SeasonGorm:      NewSeasonGorm(db),
		EpisodeGorm:     NewEpisodeGorm(db),
	}, db, err
}
