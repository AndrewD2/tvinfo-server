package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

// Season - the basic structure of a season.
type Season struct {
	gorm.Model
	Metatitle string
	Metadesc  string
	Poster    []Image
	Episodes  []Episode
	EditedBy  User
	IsLocked  bool
}

type SeasonService interface {
	ByID(id uint) *Season
	Create(image *Season) error
	Update(image *Season) error
	Delete(id uint) error
}

type SeasonGorm struct {
	*gorm.DB
}

func NewSeasonGorm(connectionInfo string) (*SeasonGorm, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		log.Println(err)
	}
	return &SeasonGorm{db}, nil
}

func (sg *SeasonGorm) ByID(id uint) *Season {
	return sg.byQuery(sg.DB.Where("id=?", id))
}

func (sg *SeasonGorm) Create(image *Season) error {
	return sg.DB.Create(image).Error
}

func (sg *SeasonGorm) Update(image *Season) error {
	return sg.DB.Save(image).Error
}

func (sg *SeasonGorm) Delete(id uint) error {
	image := &Season{Model: gorm.Model{ID: id}}
	return sg.DB.Delete(image).Error
}

func (sg *SeasonGorm) byQuery(query *gorm.DB) *Season {
	ret := &Season{}
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
