package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

// MacroSeries - the basic structure of a MacroSeries.
type MacroSeries struct {
	gorm.Model
	Name        string
	Series      []Series
	Backgrounds []Image
	Posters     []Image
	Banners     []Image
	EditedBy    uint `gorm:"ForeignKey:UserID"`
	IsLocked    bool
}

type MacroSeriesService interface {
	ByID(id uint) *Episode
	Create(ep *Episode) error
	Update(ep *Episode) error
	Delete(id uint) error
	DestructiveReset()
	AutoMigrate()
}

type MacroSeriesGorm struct {
	*gorm.DB
}

func NewMacroSeriesGorm(db *gorm.DB) *MacroSeriesGorm {
	return &MacroSeriesGorm{db}
}

func (mg *MacroSeriesGorm) ByID(id uint) *Episode {
	return mg.byQuery(mg.DB.Where("id=?", id))
}

func (mg *MacroSeriesGorm) Create(ms *MacroSeries) error {
	return mg.DB.Create(ms).Error
}

func (mg *MacroSeriesGorm) Update(ms *MacroSeries) error {
	return mg.DB.Save(ms).Error
}

func (mg *MacroSeriesGorm) Delete(id uint) error {
	episode := &Episode{Model: gorm.Model{ID: id}}
	return mg.DB.Delete(episode).Error
}
func (mg *MacroSeriesGorm) byQuery(query *gorm.DB) *Episode {
	ret := &Episode{}
	err := mg.DB.First(ret).Error
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

func (mg *MacroSeriesGorm) DestructiveReset() {
	mg.DropTableIfExists(&MacroSeries{})
	mg.AutoMigrate()
}

func (mg *MacroSeriesGorm) AutoMigrate() {
	mg.DB.AutoMigrate(&MacroSeries{})
}
