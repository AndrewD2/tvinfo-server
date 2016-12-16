package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Episode - the basic structure of an episode.
type Episode struct {
	gorm.Model
	Season         uint   `gorm:"not null"` //Season
	EpisodeNumber  uint   `gorm:"not null"`
	Title          string `gorm:"not null"`
	FirstAired     string
	Guest          []People
	Director       []People
	Writer         []People
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
	EditedBy       uint `gorm:"ForeignKey:UserID"`
	IsLocked       bool
}

type EpisodeService interface {
	ByID(id uint) *Episode
	Create(ep *Episode) error
	Update(ep *Episode) error
	Delete(id uint) error
	AutoMigrate()
	DestructiveReset()
}

type EpisodeGorm struct {
	*gorm.DB
}

func NewEpisodeGorm(db *gorm.DB) *EpisodeGorm {
	return &EpisodeGorm{db}
}

func (eg *EpisodeGorm) ByID(id uint) *Episode {
	return eg.byQuery(eg.DB.Where("id=?", id))
}

func (eg *EpisodeGorm) Create(episode *Episode) error {
	return eg.DB.Create(episode).Error
}

func (eg *EpisodeGorm) Update(episode *Episode) error {
	return eg.DB.Save(episode).Error
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

func (eg *EpisodeGorm) DestructiveReset() {
	eg.DropTableIfExists(&Episode{})
	eg.AutoMigrate()
}

func (eg *EpisodeGorm) AutoMigrate() {
	eg.DB.AutoMigrate(&Episode{})
}
