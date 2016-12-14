package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

// Series - the basic structure of a series.
type Series struct {
	gorm.Model
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
	EditedBy    User
	IsLocked    bool
}

type SeriesService interface {
	ByID(id uint) *Image
	Create(series *Series) error
	Update(series *Series) error
	Delete(id uint) error
}

type SeriesGorm struct {
	*gorm.DB
}

func NewSeriesGorm(connectionInfo string) (*SeriesGorm, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		log.Println(err)
	}
	return &SeriesGorm{db}, nil
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
