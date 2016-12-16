package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Network struct {
	gorm.Model
	Name         string `gorm:"not null;unique_index"`
	Abbreviation string
}

type NetworkService interface {
	ByID(id uint) *Network
	Create(ep *Network) error
	Update(ep *Network) error
	Delete(id uint) error
	AutoMigrate()
	DestructiveReset()
}

type NetworkGorm struct {
	*gorm.DB
}

func NewNetworkGorm(db *gorm.DB) *NetworkGorm {
	return &NetworkGorm{db}
}

func (eg *NetworkGorm) ByID(id uint) *Network {
	return eg.byQuery(eg.DB.Where("id=?", id))
}

func (eg *NetworkGorm) Create(network *Network) error {
	return eg.DB.Create(network).Error
}

func (eg *NetworkGorm) Update(network *Network) error {
	return eg.DB.Save(network).Error
}

func (eg *NetworkGorm) Delete(id uint) error {
	network := &Network{Model: gorm.Model{ID: id}}
	return eg.DB.Delete(network).Error
}
func (eg *NetworkGorm) byQuery(query *gorm.DB) *Network {
	ret := &Network{}
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

func (eg *NetworkGorm) DestructiveReset() {
	eg.DropTableIfExists(&Network{})
	eg.AutoMigrate()
}

func (eg *NetworkGorm) AutoMigrate() {
	eg.DB.AutoMigrate(&Network{})
}
