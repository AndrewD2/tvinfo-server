package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	Name string `gorm:"not null;unique_index"`
}

type RoleService interface {
	ByID(id uint) *Role
	Create(ep *Role) error
	Update(ep *Role) error
	Delete(id uint) error
	AutoMigrate()
	DestructiveReset()
}

type RoleGorm struct {
	*gorm.DB
}

func NewRoleGorm(db *gorm.DB) *RoleGorm {
	return &RoleGorm{db}
}

func (eg *RoleGorm) ByID(id uint) *Role {
	return eg.byQuery(eg.DB.Where("id=?", id))
}

func (eg *RoleGorm) Create(role *Role) error {
	return eg.DB.Create(role).Error
}

func (eg *RoleGorm) Update(role *Role) error {
	return eg.DB.Save(role).Error
}

func (eg *RoleGorm) Delete(id uint) error {
	role := &Role{Model: gorm.Model{ID: id}}
	return eg.DB.Delete(role).Error
}
func (eg *RoleGorm) byQuery(query *gorm.DB) *Role {
	ret := &Role{}
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

func (eg *RoleGorm) DestructiveReset() {
	eg.DropTableIfExists(&Role{})
	eg.AutoMigrate()
}

func (eg *RoleGorm) AutoMigrate() {
	eg.DB.AutoMigrate(&Role{})
}
