package entity

import (
	"time"

	"github.com/tonphaii/Project-sa-65/services"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("SA-65.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(
		&Login{},
		&Role{},
		&Employee{},
		&Prescription{},
		&Medicine{},
		&Patient{},
	)

	db = database

	//เราจะทำการสร้าง Admin account ไว้สำหรับการสร้าง Employee และ role ต่างๆ

	password, _ := services.Hash("123456")

	role := Role{
		Name: "admin",
	}
	db.Model(&Role{}).Create(&role)

	login := Login{
		User:     "Admin",
		Password: string(password),
	}

	Loginerr := db.Model(&Login{}).Create(&login)
	if Loginerr.Error == nil {
		db.Model(&Employee{}).Create(&Employee{
			Name:    "Admin",
			Surname: "Admin",
			Login:   login,
			Role:    role,
		})
	}

	role = Role{
		Name: "pharmacist",
	}
	db.Model(&Role{}).Create(&role)

	login_2 := Login{
		User:     "U6215576",
		Password: string(password),
	}
	Loginer_1 := db.Model(&Login{}).Create(&login_2)
	if Loginer_1.Error == nil {
		e1 := Employee{
			Name:    "Butsakorn",
			Surname: "Kharom",
			Login:   login_2,
			Role:    role,
		}
		db.Model(&Employee{}).Create(&e1)
	}

	var e1 Employee
	db.Raw("SELECT * FROM employees WHERE name = ?", "Butsakorn").Scan(&e1)
	role = Role{
		Name: "intendant",
	}
	db.Model(&Role{}).Create(&role)
	role = Role{
		Name: "payment",
	}
	db.Model(&Role{}).Create(&role)

	//Medicine
	medicine_1 := Medicine{
		Name:    "Paracetamol",
		Type:    "med",
		MFD:     time.Date(2022, 8, 28, 12, 0, 0, 0, time.UTC),
		EXP:     time.Date(2023, 8, 28, 12, 0, 0, 0, time.UTC),
		Storage: "B1",
		Amount:  100,
	}
	db.Model(&Medicine{}).Create(&medicine_1)

	medicine_2 := Medicine{
		Name:    "Menopain",
		Type:    "med",
		MFD:     time.Date(2022, 8, 30, 12, 0, 0, 0, time.UTC),
		EXP:     time.Date(2023, 8, 30, 12, 0, 0, 0, time.UTC),
		Storage: "B1",
		Amount:  100,
	}
	db.Model(&Medicine{}).Create(&medicine_2)

	medicine_3 := Medicine{
		Name:    "b-derm",
		Type:    "tha",
		MFD:     time.Date(2022, 8, 30, 12, 0, 0, 0, time.UTC),
		EXP:     time.Date(2023, 8, 30, 12, 0, 0, 0, time.UTC),
		Storage: "B2",
		Amount:  100,
	}

	db.Model(&Medicine{}).Create(&medicine_3)

	medicine_4 := Medicine{
		Name:    "Cetirizine",
		Type:    "tha",
		MFD:     time.Date(2022, 10, 24, 12, 0, 0, 0, time.UTC),
		EXP:     time.Date(2023, 10, 24, 12, 0, 0, 0, time.UTC),
		Storage: "B2",
		Amount:  100,
	}
	db.Model(&Medicine{}).Create(&medicine_4)

	patient_1 := Patient{
		PID:     "P0001",
		Name:    "กิตติมากร",
		Surname: "สอนแก้ว",
		Age:     21,
		Gender:  "หญิง",
		Allergy: "ยาแอสไพริน",
	}
	db.Model(&Patient{}).Create(&patient_1)

	patient_2 := Patient{
		PID:     "P0002",
		Name:    "ยศพล",
		Surname: "จันทะนาม",
		Age:     22,
		Gender:  "ชาย",
		Allergy: "ไม่แพ้ยาใดๆ",
	}

	db.Model(&Patient{}).Create(&patient_2)
	patient_3 := Patient{
		PID:     "P0003",
		Name:    "กฤษฎา",
		Surname: "น้อยผา",
		Age:     22,
		Gender:  "ชาย",
		Allergy: "ไม่แพ้ยาใดๆ",
	}

	db.Model(&Patient{}).Create(&patient_3)

	prescription_1 := Prescription{
		PrescriptionID: "P00001",
		Symptom:        "ไข้หวัด",
		Case_Time:      time.Now(),
		Employee:       e1,
		Medicine:       medicine_1,
		Patient:        patient_1,
	}
	db.Model(&Prescription{}).Create(&prescription_1)

	prescription_2 := Prescription{
		PrescriptionID: "P00002",
		Symptom:        "ปวดหัว",
		Case_Time:      time.Now(),
		Employee:       e1,
		Medicine:       medicine_1,
		Patient:        patient_2,
	}
	db.Model(&Prescription{}).Create(&prescription_2)

	prescription_3 := Prescription{
		PrescriptionID: "P00003",
		Symptom:        "แผลถลอก",
		Case_Time:      time.Now(),
		Employee:       e1,
		Medicine:       medicine_3,
		Patient:        patient_2,
	}
	db.Model(&Prescription{}).Create(&prescription_3)

}
