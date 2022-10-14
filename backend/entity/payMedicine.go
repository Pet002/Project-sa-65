package entity

import (
	"time"

	"gorm.io/gorm"
)

type MedicineLabel struct {
	gorm.Model
	RecordingDate time.Time `valid:"past"`

	MedicineUse string
	Warning     string

	PayMedicines []PayMedicine `gorm:"foreignKey:MedicineLabelID"`
}

type Perscription struct {
	gorm.Model
	CaseTime time.Time

	Symptom  string
	Medicine string
	Patient  string

	PayMedicine []PayMedicine `gorm:"foreignKey:PerscriptionID"`
}

type PayMedicine struct {
	gorm.Model

	Amount  uint
	Price   float64
	PayDate time.Time

	MedicineLabelID *uint
	MedicineLabel   MedicineLabel

	PerscriptionID *uint `gorm:"uniqueIndex"` //set Unique for 1 to 1 relational database
	Perscription   Perscription

	EmployeeID *uint
	Employee   Employee
}
