package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

// Episode - the basic structure of an episode.
type Episode struct {
	gorm.Model
	Season         Season
	EpisodeNumber  uint
	Title          string
	FirstAired     string
	Guest          []string
	Director       []string
	Writer         []string
	ProductionCode string
	Description    string
	Poster         []Image
	DVDID          int
	DVDSeason      int
	DVDEpisode     int
	DVDChapter     int
	BDID           int
	BDSeason       int
	BDEpisode      int
	BDChapter      int
	Absolute       int
	IMDBID         string
	IsMovie        bool
	EditedBy       User
	IsLocked       bool
}

type EpisodeService interface {
	ByID(id uint) *Episode
	Create(ep *Episode) error
	Update(ep *Episode) error
	Delete(id uint) error
}

type EpisodeGorm struct {
	*gorm.DB
}

func (eg *EpisodeGorm) ByID(id uint) *Episode {
	return eg.byQuery(eg.DB.Where("id=?", id))
}

func (eg *EpisodeGorm) Create(image *Image) error {
	return eg.DB.Create(image).Error
}

func (eg *EpisodeGorm) Update(image *Image) error {
	return eg.DB.Save(image).Error
}

func (eg *EpisodeGorm) Delete(id uint) error {
	episode := &Episode{Model: gorm.Model{ID: id}}
	return eg.DB.Delete(episode).Error
}
func (eg *EpisodeGorm) byQuery(query *gorm.DB) *Episode {
	ret := &Episode{}
	err := eg.DB.First(ret).Error
	switch err {
	case nil:
		return ret
	case gorm.ErrRecordNotFound:
		return nil
	default:
		log.Println(err)
	}
	return nil
}
