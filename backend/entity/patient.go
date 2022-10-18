package entity

import (
	"gorm.io/gorm"
)

// คนไข้
type Patient struct {
	gorm.Model

	PID     string
	Name    string
	Surname string
	Age     uint
	Gender  string
	Allergy string

	Prescriptions []Prescription `gorm:"foreignKey:PatientID"`
}
