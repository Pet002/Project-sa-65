package entity
import (
	"gorm.io/gorm"
)
type MedicineCompany struct {
	gorm.Model
	Company_Name	string
	Location		string
	
}
