package entity

import (
	"time"

	"gorm.io/gorm"
)

type PayMedicine struct {
	gorm.Model
	Amount  int
	Price   float32
	Paydate time.Time

	PrescriptionID uint
	Prescription   Prescription

	Receipt []Receipt `gorm:"foreignKey:PayMedicineID"`
}
