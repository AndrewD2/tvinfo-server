package models

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

// Series - the basic structure of a series.
type Series struct {
	gorm.Model
	Name            string `gorm:"not null"`
	Description     string
	StartDate       time.Time
	EndDate         time.Time
	EpisodeDuration uint
	NetworkID       uint
	Genre           []Genre
	Creator         []People
	Actors          []People
	SeasonID        uint
	Backgrounds     []Image
	Banners         []Image
	Poster          []Image
	EditedBy        uint `gorm:"ForeignKey:UserID"`
	IsLocked        bool
}

type SeriesService interface {
	ByID(id uint) *Image
	Create(series *Series) error
	Update(series *Series) error
	Delete(id uint) error
	DestructiveReset()
	AutoMigrate()
}

type SeriesGorm struct {
	*gorm.DB
}

func NewSeriesGorm(db *gorm.DB) *SeriesGorm {
	return &SeriesGorm{db}
}

func (sg *SeriesGorm) ByID(id uint) *Series {
	return sg.byQuery(sg.DB.Where("id=?", id))
}

func (sg *SeriesGorm) Create(series *Series) error {
	return sg.DB.Create(series).Error
}

func (sg *SeriesGorm) Update(series *Series) error {
	return sg.DB.Save(series).Error
}

func (sg *SeriesGorm) Delete(id uint) error {
	series := &Image{Model: gorm.Model{ID: id}}
	return sg.DB.Delete(series).Error
}

func (sg *SeriesGorm) byQuery(query *gorm.DB) *Series {
	ret := &Series{}
	err := sg.DB.First(ret).Error
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
func (sg *SeriesGorm) DestructiveReset() {
	sg.DropTableIfExists(&Series{})
	sg.AutoMigrate()
}

func (sg *SeriesGorm) AutoMigrate() {
	sg.DB.AutoMigrate(&Series{})
}
