package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Image - the basic structure of an Image.
type Image struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	URL         string
	AddedBy     uint `gorm:"ForeignKey:UserID"`
}

type ImageService interface {
	ByID(id uint) *Image
	Create(image *Image) error
	Update(image *Image) error
	Delete(id uint) error
}

type ImageGorm struct {
	*gorm.DB
}

func NewImageGorm(db *gorm.DB) *ImageGorm {
	return &ImageGorm{db}
}

func (ig *ImageGorm) ByID(id uint) *Image {
	return ig.byQuery(ig.DB.Where("id=?", id))
}

func (ig *ImageGorm) Create(image *Image) error {
	return ig.DB.Create(image).Error
}

func (ig *ImageGorm) Update(image *Image) error {
	return ig.DB.Save(image).Error
}

func (ig *ImageGorm) Delete(id uint) error {
	image := &Image{Model: gorm.Model{ID: id}}
	return ig.DB.Delete(image).Error
}

func (ig *ImageGorm) byQuery(query *gorm.DB) *Image {
	ret := &Image{}
	err := ig.DB.First(ret).Error
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

func (eg *ImageGorm) DestructiveReset() {
	eg.DropTableIfExists(&Image{})
	eg.AutoMigrate()
}

func (eg *ImageGorm) AutoMigrate() {
	eg.DB.AutoMigrate(&Image{})
}