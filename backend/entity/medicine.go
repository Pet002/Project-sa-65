package entity

import (
	"time"

	"gorm.io/gorm"
)

// types
type Type struct {
	gorm.Model
	Tmedicine  string
	Utilzation string

	Medicine []Medicine `gorm:"foreignKey:TypeID"`
}

type Storage struct {
	gorm.Model
	Name string

	Medicine []Medicine `gorm:"foreignKey:StorageID"`
}

type Medicine struct {
	gorm.Model
	Name   string
	MFD    time.Time
	EXP    time.Time
	Amount int

	EmployeeID *uint
	Employee   Employee

	TypeID *uint
	Type   Type

	StorageID *uint
	Storage   Storage
}
