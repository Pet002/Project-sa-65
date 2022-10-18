package entity

import (
	"gorm.io/gorm"
)

type Storage struct {
	gorm.Model
	Name string

	Medicine []Medicine `gorm:"foreignKey:StorageID"`
}
