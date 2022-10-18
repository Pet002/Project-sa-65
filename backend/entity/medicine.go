package entity

import (
	"time"

	"gorm.io/gorm"
)

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
