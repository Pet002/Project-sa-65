package entity

import (
	"github.com/Pet002/Project-sa-65/services"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("MedicineRoom.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(
		&Login{},
		&Role{},
		&Employee{},
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
		Name: "intendant",
	}
	db.Model(&Role{}).Create(&role)

	role = Role{
		Name: "pharmacist",
	}
	db.Model(&Role{}).Create(&role)

	role = Role{
		Name: "payment",
	}
	db.Model(&Role{}).Create(&role)
	db.Model(&Role{}).Create(&role)
	login1 := Login{
		User:     "GODCY",
		Password: string(password),
	}

	Loginerr1 := db.Model(&Login{}).Create(&login1)
	if Loginerr1.Error == nil {
		db.Model(&Employee{}).Create(&Employee{
			Name:    "GOD",
			Surname: "CY",
			Login:   login1,
			Role:    role,
		})
	}


	tp1 := Types{
		TypeName: "ชำระเงินสด",
	}
	db.Model(&Types{}).Create(&tp1)

	tp2 := Types{
		TypeName: "ชำระด้วยการโอน",
	}
	db.Model(&Types{}).Create(&tp2)

	db.Model(&Receipt{}).Create(&Receipt{
		//Employee: 
		//PayMedicine:
		Types: tp1,

		TotalPrice: 100,
		Receive: 500,
		Refund: 400,

	})
}
