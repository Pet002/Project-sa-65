package entity

import (
	"time"

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
		&Type{},
		&Storage{},
		&Medicine{},
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
	login_2 := Login{
		User:     "B6217174",
		Password: string(password),
	}
	Loginerr2 := db.Model(&Login{}).Create(&login_2)

	emp := Employee{
		Name:    "Palida",
		Surname: "Suwannarat",
		Login:   login_2,
		Role:    role,
	}
	if Loginerr2.Error == nil {
		db.Model(&Employee{}).Create(&emp)
	}

	login_3 := Login{
		User:     "B6217162",
		Password: string(password),
	}
	Loginerr3 := db.Model(&Login{}).Create(&login_3)

	emp2 := Employee{
		Name:    "Pichanon",
		Surname: "Suwannarat",
		Login:   login_3,
		Role:    role,
	}
	if Loginerr3.Error == nil {
		db.Model(&Employee{}).Create(&emp2)
	}

	role = Role{
		Name: "pharmacist",
	}
	db.Model(&Role{}).Create(&role)

	login_4 := Login{
		User:     "pharmacist",
		Password: string(password),
	}
	Loginerr4 := db.Model(&Login{}).Create(&login_4)

	emp4 := Employee{
		Name:    "Pichanon",
		Surname: "Suwannarat",
		Login:   login_4,
		Role:    role,
	}
	if Loginerr4.Error == nil {
		db.Model(&Employee{}).Create(&emp4)
	}

	role = Role{
		Name: "payment",
	}
	db.Model(&Role{}).Create(&role)

	login_5 := Login{
		User:     "pharmacist",
		Password: string(password),
	}
	Loginerr5 := db.Model(&Login{}).Create(&login_5)

	emp5 := Employee{
		Name:    "Pichanon",
		Surname: "Suwannarat",
		Login:   login_5,
		Role:    role,
	}
	if Loginerr5.Error == nil {
		db.Model(&Employee{}).Create(&emp5)
	}

	//-------------------------------------- medicine ----------------------------------------------

	//----------Type-------------------
	tha := Type{
		Tmedicine:  "ยาใช้ภายนอก",
		Utilzation: "ทา",
	}
	db.Model(&Type{}).Create(&tha)

	med := Type{
		Tmedicine:  "ยาใช้ภายใน",
		Utilzation: "เม็ด",
	}
	db.Model(&Type{}).Create(&med)

	cheed := Type{
		Tmedicine:  "ยาใช้ภายใน",
		Utilzation: "ฉีด",
	}
	db.Model(&Type{}).Create(&cheed)

	nam := Type{
		Tmedicine:  "ยาใช้ภายใน",
		Utilzation: "น้ำ",
	}
	db.Model(&Type{}).Create(&nam)

	//-------------Storage----------------
	b1 := Storage{
		Name: "B1",
	}
	db.Model(&Storage{}).Create(&b1)

	b2 := Storage{
		Name: "B2",
	}
	db.Model(&Storage{}).Create(&b2)

	//-------------Medicine---------------

	db.Model(&Medicine{}).Create(&Medicine{
		Employee: emp,
		Name:     "Paracetamol",
		Type:     med,
		MFD:      time.Date(2022, 8, 28, 12, 0, 0, 0, time.UTC),
		EXP:      time.Date(2023, 8, 28, 12, 0, 0, 0, time.UTC),
		Storage:  b1,
		Amount:   100,
	})
	/*db.Model(&Medicine{}).Create(&Medicine{
		Employee: emp,
		Name:     "Menopain",
		Type:     med,
		MFD:      time.Date(2022, 8, 30, 12, 0, 0, 0, time.UTC),
		EXP:      time.Date(2023, 8, 30, 12, 0, 0, 0, time.UTC),
		Storage:  b1,
		Amount:   100,
	})
	db.Model(&Medicine{}).Create(&Medicine{
		Employee: emp,
		Name:     "b-derm",
		Type:     tha,
		MFD:      time.Date(2022, 8, 30, 12, 0, 0, 0, time.UTC),
		EXP:      time.Date(2023, 8, 30, 12, 0, 0, 0, time.UTC),
		Storage:  b2,
		Amount:   100,
	})
	db.Model(&Medicine{}).Create(&Medicine{
		Employee: emp,
		Name:     "Cetirizine",
		Type:     med,
		MFD:      time.Date(2022, 10, 24, 12, 0, 0, 0, time.UTC),
		EXP:      time.Date(2023, 10, 24, 12, 0, 0, 0, time.UTC),
		Storage:  b2,
		Amount:   100,
	})
	db.Model(&Medicine{}).Create(&Medicine{
		Employee: emp,
		Name:     "ฺBromhexine", //ยาละลายเสมหะ
		Type:     med,
		MFD:      time.Date(2022, 10, 24, 12, 0, 0, 0, time.UTC),
		EXP:      time.Date(2023, 10, 24, 12, 0, 0, 0, time.UTC),
		Storage:  b2,
		Amount:   100,
	})
	db.Model(&Medicine{}).Create(&Medicine{
		Employee: emp,
		Name:     "Cenor", //ยาต้านเชื้อแบคทีเรีย
		Type:     med,
		MFD:      time.Date(2022, 10, 24, 12, 0, 0, 0, time.UTC),
		EXP:      time.Date(2023, 10, 24, 12, 0, 0, 0, time.UTC),
		Storage:  b1,
		Amount:   100,
	})
	db.Model(&Medicine{}).Create(&Medicine{
		Employee: emp,
		Name:     "Tramadol", //ยาบรรเทาอาการปวดรุนแรง
		Type:     med,
		MFD:      time.Date(2022, 10, 24, 12, 0, 0, 0, time.UTC),
		EXP:      time.Date(2023, 10, 24, 12, 0, 0, 0, time.UTC),
		Storage:  b1,
		Amount:   100,
	})
	db.Model(&Medicine{}).Create(&Medicine{
		Employee: emp,
		Name:     "Salol et Menthol Mixture", //ยาธาตุน้ำขาว
		Type:     med,
		MFD:      time.Date(2022, 10, 24, 12, 0, 0, 0, time.UTC),
		EXP:      time.Date(2023, 10, 24, 12, 0, 0, 0, time.UTC),
		Storage:  b2,
		Amount:   100,
	})
	db.Model(&Medicine{}).Create(&Medicine{
		Employee: emp,
		Name:     "Atorvastatin", //ยา]f++
		Type:     med,
		MFD:      time.Date(2022, 10, 24, 12, 0, 0, 0, time.UTC),
		EXP:      time.Date(2023, 10, 24, 12, 0, 0, 0, time.UTC),
		Storage:  b2,
		Amount:   100,
	})*/

	//--------------------------------------------------- Medicine Label part -------------------------------------

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
		Employee:      emp5,
	})

	db.Model(&MedicineLabel{}).Create(&MedicineLabel{
		RecordingDate: time.Now(),
		Warning:       w2,
		MedicineUse:   mu2,
		Employee:      emp5,
	})

	db.Model(&MedicineLabel{}).Create(&MedicineLabel{
		RecordingDate: time.Now(),
		Warning:       w3,
		MedicineUse:   mu3,
		Employee:      emp5,
	})

}
