package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Genre struct {
	gorm.Model
	Name string `gorm:"not null;unique_index"`
}

type GenreService interface {
	ByID(id uint) *Genre
	Create(ep *Genre) error
	Update(ep *Genre) error
	Delete(id uint) error
	AutoMigrate()
	DestructiveReset()
}

type GenreGorm struct {
	*gorm.DB
}

func NewGenreGorm(db *gorm.DB) *GenreGorm {
	return &GenreGorm{db}
}

func (eg *GenreGorm) ByID(id uint) *Genre {
	return eg.byQuery(eg.DB.Where("id=?", id))
}

func (eg *GenreGorm) Create(genre *Genre) error {
	return eg.DB.Create(genre).Error
}

func (eg *GenreGorm) Update(genre *Genre) error {
	return eg.DB.Save(genre).Error
}

func (eg *GenreGorm) Delete(id uint) error {
	genre := &Genre{Model: gorm.Model{ID: id}}
	return eg.DB.Delete(genre).Error
}
func (eg *GenreGorm) byQuery(query *gorm.DB) *Genre {
	ret := &Genre{}
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

func (eg *GenreGorm) DestructiveReset() {
	eg.DropTableIfExists(&Genre{})
	eg.AutoMigrate()
}

func (eg *GenreGorm) AutoMigrate() {
	eg.DB.AutoMigrate(&Genre{})
}
