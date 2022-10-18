package entity

import (
	"gorm.io/gorm"
)

type Receipt struct {
	gorm.Model
	TotalPrice int
	Receive    int
	Refund     int

	EmployeeID *uint
	Employee   Employee

	TypesID *uint
	Types   Types

	PayMedicineID *uint
	PayMedicine   PayMedicine
}
