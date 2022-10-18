package entity

import (
	"gorm.io/gorm"
)

type Types struct {
	gorm.Model

	TypeName string
	Receip   []Receipt `gorm:"foreignKey:TypesID"`
}
