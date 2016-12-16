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
	EditedBy  uint `gorm:"ForeignKey:UserID"`
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

func NewSeasonGorm(db *gorm.DB) *SeasonGorm {
	return &SeasonGorm{db}
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

func (sg *SeasonGorm) DestructiveReset() {
	sg.DropTableIfExists(&Season{})
	sg.AutoMigrate()
}

func (sg *SeasonGorm) AutoMigrate() {
	sg.DB.AutoMigrate(&Season{})
}
