package entity

import (
	"gorm.io/gorm"
)

type Type struct {
	gorm.Model
	Tmedicine  string
	Utilzation string

	Medicine []Medicine `gorm:"foreignKey:TypeID"`
}
