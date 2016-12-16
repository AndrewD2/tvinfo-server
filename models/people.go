package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

type People struct {
	gorm.Model
	Name  string
	Roles []Role
}

type PeopleService interface {
	ByID(id uint) *People
	Create(ep *People) error
	Update(ep *People) error
	Delete(id uint) error
	AutoMigrate()
	DestructiveReset()
}

type PeopleGorm struct {
	*gorm.DB
}

func (eg *PeopleGorm) ByID(id uint) *People {
	return eg.byQuery(eg.DB.Where("id=?", id))
}

func (eg *PeopleGorm) Create(people *People) error {
	return eg.DB.Create(people).Error
}

func (eg *PeopleGorm) Update(people *People) error {
	return eg.DB.Save(people).Error
}

func (eg *PeopleGorm) Delete(id uint) error {
	people := &People{Model: gorm.Model{ID: id}}
	return eg.DB.Delete(people).Error
}
func (eg *PeopleGorm) byQuery(query *gorm.DB) *People {
	ret := &People{}
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

func NewPeopleGorm(db *gorm.DB) *PeopleGorm {
	return &PeopleGorm{db}
}

func (eg *PeopleGorm) DestructiveReset() {
	eg.DropTableIfExists(&People{})
	eg.AutoMigrate()
}

func (eg *PeopleGorm) AutoMigrate() {
	eg.DB.AutoMigrate(&People{})
}
