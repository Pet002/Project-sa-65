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

	// Migrate the schema

	database.AutoMigrate(
		&MedicineUse{},
		&Warning{},
		&MedicineLabel{},
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

}
