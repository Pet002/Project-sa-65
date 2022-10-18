package entity
import (
	"time"
	"gorm.io/gorm"
)
type MedicineOrder struct {
	gorm.Model
	OrderID uint
	OrderAmount uint
	OrderTime time.Time
	EmployeeID *uint
	Employee Employee
	
	CompanyID *uint
	MedicineCompany MedicineCompany
}
