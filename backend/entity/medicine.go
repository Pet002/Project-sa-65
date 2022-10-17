package entity

import (
	"time"

	"gorm.io/gorm"
)

type Medicine struct {
	gorm.Model
	Name    string
	Type    string
	MFD     time.Time
	EXP     time.Time
	Amount  uint
	Storage string

	Prescriptions []Prescription `gorm:"foreignKey:MedicineID"`
}
