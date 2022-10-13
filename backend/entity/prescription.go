package entity

import (
	"time"

	"gorm.io/gorm"
)

// Entity หลัก
type Prescription struct {
	gorm.Model
	PrescriptionID string
	Symptom        string //อาการป่วย
	Case_Time      time.Time

	EmployeeID *uint
	Employee   Employee

	MedicineID *uint
	Medicine   Medicine

	PatientID *uint
	Patient   Patient
}
