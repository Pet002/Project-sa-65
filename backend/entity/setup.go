// package entity

// import (
// 	"github.com/Thanaporn4226/Project-sa-65/services"
// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"
// )

// var db *gorm.DB

// func DB() *gorm.DB {
// 	return db
// }

// func SetupDatabase() {
// 	database, err := gorm.Open(sqlite.Open("MedicineRoom.db"), &gorm.Config{})

// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	database.AutoMigrate(
// 		&Login{},
// 		&Role{},
// 		&Employee{},
// 	)

// 	db = database

// 	//เราจะทำการสร้าง Admin account ไว้สำหรับการสร้าง Employee และ role ต่างๆ

// 	password, _ := services.Hash("123456")

// 	role := Role{
// 		Name: "admin",
// 	}

// 	login := Login{
// 		User:     "Admin",
// 		Password: string(password),
// 	}

// 	db.Model(&Role{}).Create(&role)
// 	Loginerr := db.Model(&Login{}).Create(&login)
// 	if Loginerr.Error == nil {
// 		db.Model(&Employee{}).Create(&Employee{
// 			Name:    "Admin",
// 			Surname: "Admin",
// 			Login:   login,
// 			Role:    role,
// 		})
// 	}

// }

package entity

import (
	"time"

	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema

	database.AutoMigrate(
		&MedicineUse{},
		&Warning{},
		&MedicineLabel{},
	)

	db = database

	//Employee Data
	emp1 := Employee{
		Name:    "Thanaporn",
		Surname: "Jansap",
		//Role:    "Pharmacist",
	}
	db.Model(&Employee{}).Create(&emp1)

	emp2 := Employee{
		Name:    "Name",
		Surname: "Surname",
		//Role:    "Pharmacist",
	}
	db.Model(&Employee{}).Create(&emp2)

	emp3 := Employee{
		Name:    "Manee",
		Surname: "Jaidee",
		//Role:    "Pharmacist",
	}
	db.Model(&Employee{}).Create(&emp3)

	//MedicineUse Data
	mu1 := MedicineUse{
		How_To_Use: "ครั้งละ 1 เม็ด ทุก 4-6 ชั่วโง เวลาปวดหรือมีไข้",
	}
	db.Model(&MedicineUse{}).Create(&mu1)

	mu2 := MedicineUse{
		How_To_Use: "ครั้งละ 1 เม็ด หลังอาหาร เช้า-กลางวัน-เย็น",
	}
	db.Model(&MedicineUse{}).Create(&mu2)

	mu3 := MedicineUse{
		How_To_Use: "ครั้งละ 1 เม็ด ก่อนอาหาร เช้า-กลางวัน-เย็น",
	}
	db.Model(&MedicineUse{}).Create(&mu3)

	// Warning Data
	w1 := Warning{
		Medicine_Warning: "ห้ามใช้เกิน 8 เม็ดต่อวัน",
	}
	db.Model(&Warning{}).Create(&w1)

	w2 := Warning{
		Medicine_Warning: "ทานยาแล้วอาจรู้สึกง่วงซึม",
	}
	db.Model(&Warning{}).Create(&w2)

	w3 := Warning{
		Medicine_Warning: "ไม่ควรใช้เกินกว่าขนาดที่ระบุ",
	}
	db.Model(&Warning{}).Create(&w3)

	//MedicineLabel Data
	db.Model(&MedicineLabel{}).Create(&MedicineLabel{
		RecordingDate: time.Now(),
		Warning:       w1,
		MedicineUse:   mu1,
		Employee:      emp1,
	})

	db.Model(&MedicineLabel{}).Create(&MedicineLabel{
		RecordingDate: time.Now(),
		Warning:       w2,
		MedicineUse:   mu2,
		Employee:      emp1,
	})

	db.Model(&MedicineLabel{}).Create(&MedicineLabel{
		RecordingDate: time.Now(),
		Warning:       w3,
		MedicineUse:   mu3,
		Employee:      emp2,
	})
}
