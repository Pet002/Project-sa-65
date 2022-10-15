package entity

import (
	"time"

	"gorm.io/gorm"
)

// Main Entity
type MedicineLabel struct {
	gorm.Model

	RecordingDate time.Time

	EmployeeID *uint
	Employee   Employee

	WarningID *uint
	Warning   Warning

	MedicineUseID *uint
	MedicineUse   MedicineUse
}

type MedicineUse struct {
	gorm.Model
	How_To_Use    string
	MedicineLabel []MedicineLabel `gorm:"foreignKey:MedicineUseID"`
}

type Warning struct {
	gorm.Model
	Medicine_Warning string
	MedicineLabel    []MedicineLabel `gorm:"foreignKey:WarningID"`
}
