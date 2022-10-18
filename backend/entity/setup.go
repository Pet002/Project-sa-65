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
		//Medicine label
		&MedicineUse{},
		&Warning{},
		&MedicineLabel{},
		//Employee system
		&Login{},
		&Role{},
		&Employee{},
		//Prescription
		&Prescription{},
		&Patient{},
		//Medicine
		&Type{},
		&Storage{},
		&Medicine{},
		//PayMedicine
		&PayMedicine{},
		//Receipt
		&Types{},
		&Receipt{},
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

	// role = Role{
	// 	Name: "intendant",
	// }
	// db.Model(&Role{}).Create(&role)
	// login_2 := Login{
	// 	User:     "B6217174",
	// 	Password: string(password),
	// }
	// Loginerr2 := db.Model(&Login{}).Create(&login_2)

	// emp := Employee{
	// 	Name:    "Palida",
	// 	Surname: "Suwannarat",
	// 	Login:   login_2,
	// 	Role:    role,
	// }
	// if Loginerr2.Error == nil {
	// 	db.Model(&Employee{}).Create(&emp)
	// }

	// login_3 := Login{
	// 	User:     "B6217162",
	// 	Password: string(password),
	// }
	// Loginerr3 := db.Model(&Login{}).Create(&login_3)

	// emp2 := Employee{
	// 	Name:    "Pichanon",
	// 	Surname: "Suwannarat",
	// 	Login:   login_3,
	// 	Role:    role,
	// }
	// if Loginerr3.Error == nil {
	// 	db.Model(&Employee{}).Create(&emp2)
	// }

	// role = Role{
	// 	Name: "payment",
	// }
	// db.Model(&Role{}).Create(&role)

	// login_4 := Login{
	// 	User:     "pharmacist",
	// 	Password: string(password),
	// }
	// Loginerr4 := db.Model(&Login{}).Create(&login_4)

	// emp4 := Employee{
	// 	Name:    "Pichanon",
	// 	Surname: "Suwannarat",
	// 	Login:   login_4,
	// 	Role:    role,
	// }
	// if Loginerr4.Error == nil {
	// 	db.Model(&Employee{}).Create(&emp4)
	// }
	// login_6 := Login{
	// 	User:     "U6215576",
	// 	Password: string(password),
	// }
	// Loginer_6 := db.Model(&Login{}).Create(&login_6)
	// if Loginer_6.Error == nil {
	// 	e1 := Employee{
	// 		Name:    "Butsakorn",
	// 		Surname: "Kharom",
	// 		Login:   login_6,
	// 		Role:    role,
	// 	}
	// 	db.Model(&Employee{}).Create(&e1)
	// }

	// role = Role{
	// 	Name: "pharmacist",
	// }
	// db.Model(&Role{}).Create(&role)

	// login_5 := Login{
	// 	User:     "pharmacist",
	// 	Password: string(password),
	// }
	// Loginerr5 := db.Model(&Login{}).Create(&login_5)

	// emp5 := Employee{
	// 	Name:    "pharmacist",
	// 	Surname: "Test",
	// 	Login:   login_5,
	// 	Role:    role,
	// }
	// if Loginerr5.Error == nil {
	// 	db.Model(&Employee{}).Create(&emp5)
	// }

	// //-------------------------------------- medicine ----------------------------------------------

	// //----------Type-------------------
	// tha := Type{
	// 	Tmedicine:  "ยาใช้ภายนอก",
	// 	Utilzation: "ทา",
	// }
	// db.Model(&Type{}).Create(&tha)

	// med := Type{
	// 	Tmedicine:  "ยาใช้ภายใน",
	// 	Utilzation: "เม็ด",
	// }
	// db.Model(&Type{}).Create(&med)

	// cheed := Type{
	// 	Tmedicine:  "ยาใช้ภายใน",
	// 	Utilzation: "ฉีด",
	// }
	// db.Model(&Type{}).Create(&cheed)

	// nam := Type{
	// 	Tmedicine:  "ยาใช้ภายใน",
	// 	Utilzation: "น้ำ",
	// }
	// db.Model(&Type{}).Create(&nam)

	// //-------------Storage----------------
	// b1 := Storage{
	// 	Name: "B1",
	// }
	// db.Model(&Storage{}).Create(&b1)

	// b2 := Storage{
	// 	Name: "B2",
	// }
	// db.Model(&Storage{}).Create(&b2)

	// //-------------Medicine---------------

	// db.Model(&Medicine{}).Create(&Medicine{
	// 	Employee: emp,
	// 	Name:     "Paracetamol",
	// 	Type:     med,
	// 	MFD:      time.Date(2022, 8, 28, 12, 0, 0, 0, time.UTC),
	// 	EXP:      time.Date(2023, 8, 28, 12, 0, 0, 0, time.UTC),
	// 	Storage:  b1,
	// 	Amount:   100,
	// })
	// /*db.Model(&Medicine{}).Create(&Medicine{
	// 	Employee: emp,
	// 	Name:     "Menopain",
	// 	Type:     med,
	// 	MFD:      time.Date(2022, 8, 30, 12, 0, 0, 0, time.UTC),
	// 	EXP:      time.Date(2023, 8, 30, 12, 0, 0, 0, time.UTC),
	// 	Storage:  b1,
	// 	Amount:   100,
	// })
	// db.Model(&Medicine{}).Create(&Medicine{
	// 	Employee: emp,
	// 	Name:     "b-derm",
	// 	Type:     tha,
	// 	MFD:      time.Date(2022, 8, 30, 12, 0, 0, 0, time.UTC),
	// 	EXP:      time.Date(2023, 8, 30, 12, 0, 0, 0, time.UTC),
	// 	Storage:  b2,
	// 	Amount:   100,
	// })
	// db.Model(&Medicine{}).Create(&Medicine{
	// 	Employee: emp,
	// 	Name:     "Cetirizine",
	// 	Type:     med,
	// 	MFD:      time.Date(2022, 10, 24, 12, 0, 0, 0, time.UTC),
	// 	EXP:      time.Date(2023, 10, 24, 12, 0, 0, 0, time.UTC),
	// 	Storage:  b2,
	// 	Amount:   100,
	// })
	// db.Model(&Medicine{}).Create(&Medicine{
	// 	Employee: emp,
	// 	Name:     "ฺBromhexine", //ยาละลายเสมหะ
	// 	Type:     med,
	// 	MFD:      time.Date(2022, 10, 24, 12, 0, 0, 0, time.UTC),
	// 	EXP:      time.Date(2023, 10, 24, 12, 0, 0, 0, time.UTC),
	// 	Storage:  b2,
	// 	Amount:   100,
	// })
	// db.Model(&Medicine{}).Create(&Medicine{
	// 	Employee: emp,
	// 	Name:     "Cenor", //ยาต้านเชื้อแบคทีเรีย
	// 	Type:     med,
	// 	MFD:      time.Date(2022, 10, 24, 12, 0, 0, 0, time.UTC),
	// 	EXP:      time.Date(2023, 10, 24, 12, 0, 0, 0, time.UTC),
	// 	Storage:  b1,
	// 	Amount:   100,
	// })
	// db.Model(&Medicine{}).Create(&Medicine{
	// 	Employee: emp,
	// 	Name:     "Tramadol", //ยาบรรเทาอาการปวดรุนแรง
	// 	Type:     med,
	// 	MFD:      time.Date(2022, 10, 24, 12, 0, 0, 0, time.UTC),
	// 	EXP:      time.Date(2023, 10, 24, 12, 0, 0, 0, time.UTC),
	// 	Storage:  b1,
	// 	Amount:   100,
	// })
	// db.Model(&Medicine{}).Create(&Medicine{
	// 	Employee: emp,
	// 	Name:     "Salol et Menthol Mixture", //ยาธาตุน้ำขาว
	// 	Type:     med,
	// 	MFD:      time.Date(2022, 10, 24, 12, 0, 0, 0, time.UTC),
	// 	EXP:      time.Date(2023, 10, 24, 12, 0, 0, 0, time.UTC),
	// 	Storage:  b2,
	// 	Amount:   100,
	// })
	// db.Model(&Medicine{}).Create(&Medicine{
	// 	Employee: emp,
	// 	Name:     "Atorvastatin", //ยา]f++
	// 	Type:     med,
	// 	MFD:      time.Date(2022, 10, 24, 12, 0, 0, 0, time.UTC),
	// 	EXP:      time.Date(2023, 10, 24, 12, 0, 0, 0, time.UTC),
	// 	Storage:  b2,
	// 	Amount:   100,
	// })*/

	// //--------------------------------------------------- Medicine Label part -------------------------------------

	// //MedicineUse Data
	// mu1 := MedicineUse{
	// 	How_To_Use: "ครั้งละ 1 เม็ด ทุก 4-6 ชั่วโง เวลาปวดหรือมีไข้",
	// }
	// db.Model(&MedicineUse{}).Create(&mu1)

	// mu2 := MedicineUse{
	// 	How_To_Use: "ครั้งละ 1 เม็ด หลังอาหาร เช้า-กลางวัน-เย็น",
	// }
	// db.Model(&MedicineUse{}).Create(&mu2)

	// mu3 := MedicineUse{
	// 	How_To_Use: "ครั้งละ 1 เม็ด ก่อนอาหาร เช้า-กลางวัน-เย็น",
	// }
	// db.Model(&MedicineUse{}).Create(&mu3)

	// // Warning Data
	// w1 := Warning{
	// 	Medicine_Warning: "ห้ามใช้เกิน 8 เม็ดต่อวัน",
	// }
	// db.Model(&Warning{}).Create(&w1)

	// w2 := Warning{
	// 	Medicine_Warning: "ทานยาแล้วอาจรู้สึกง่วงซึม",
	// }
	// db.Model(&Warning{}).Create(&w2)

	// w3 := Warning{
	// 	Medicine_Warning: "ไม่ควรใช้เกินกว่าขนาดที่ระบุ",
	// }
	// db.Model(&Warning{}).Create(&w3)

	// //MedicineLabel Data
	// db.Model(&MedicineLabel{}).Create(&MedicineLabel{
	// 	RecordingDate: time.Now(),
	// 	Warning:       w1,
	// 	MedicineUse:   mu1,
	// 	Employee:      emp5,
	// })

	// db.Model(&MedicineLabel{}).Create(&MedicineLabel{
	// 	RecordingDate: time.Now(),
	// 	Warning:       w2,
	// 	MedicineUse:   mu2,
	// 	Employee:      emp5,
	// })

	// db.Model(&MedicineLabel{}).Create(&MedicineLabel{
	// 	RecordingDate: time.Now(),
	// 	Warning:       w3,
	// 	MedicineUse:   mu3,
	// 	Employee:      emp5,
	// })

	// var e1 Employee
	// db.Raw("SELECT * FROM employees WHERE name = ?", "Butsakorn").Scan(&e1)
	// role = Role{
	// 	Name: "intendant",
	// }
	// db.Model(&Role{}).Create(&role)
	// role = Role{
	// 	Name: "payment",
	// }
	// db.Model(&Role{}).Create(&role)
	// db.Model(&Role{}).Create(&role)
	// login1 := Login{
	// 	User:     "GODCY",
	// 	Password: string(password),
	// }

	// Loginerr1 := db.Model(&Login{}).Create(&login1)
	// if Loginerr1.Error == nil {
	// 	db.Model(&Employee{}).Create(&Employee{
	// 		Name:    "GOD",
	// 		Surname: "CY",
	// 		Login:   login1,
	// 		Role:    role,
	// 	})
	// }

	// tp1 := Types{
	// 	TypeName: "ชำระเงินสด",
	// }
	// db.Model(&Types{}).Create(&tp1)

	// tp2 := Types{
	// 	TypeName: "ชำระด้วยการโอน",
	// }
	// db.Model(&Types{}).Create(&tp2)

	// db.Model(&Receipt{}).Create(&Receipt{
	// 	//Employee:
	// 	//PayMedicine:
	// 	Types: tp1,

	// 	TotalPrice: 100,
	// 	Receive: 500,
	// 	Refund: 400,

	// //Medicine
	// medicine_1 := Medicine{
	// 	Name:    "Paracetamol",
	// 	Type:    med,
	// 	MFD:     time.Date(2022, 8, 28, 12, 0, 0, 0, time.UTC),
	// 	EXP:     time.Date(2023, 8, 28, 12, 0, 0, 0, time.UTC),
	// 	Storage: b1,
	// 	Amount:  100,
	// }
	// db.Model(&Medicine{}).Create(&medicine_1)

	// medicine_2 := Medicine{
	// 	Name:    "Menopain",
	// 	Type:    med,
	// 	MFD:     time.Date(2022, 8, 30, 12, 0, 0, 0, time.UTC),
	// 	EXP:     time.Date(2023, 8, 30, 12, 0, 0, 0, time.UTC),
	// 	Storage: b1,
	// 	Amount:  100,
	// }
	// db.Model(&Medicine{}).Create(&medicine_2)

	// medicine_3 := Medicine{
	// 	Name:    "b-derm",
	// 	Type:    med,
	// 	MFD:     time.Date(2022, 8, 30, 12, 0, 0, 0, time.UTC),
	// 	EXP:     time.Date(2023, 8, 30, 12, 0, 0, 0, time.UTC),
	// 	Storage: b2,
	// 	Amount:  100,
	// }

	// db.Model(&Medicine{}).Create(&medicine_3)

	// medicine_4 := Medicine{
	// 	Name:    "Cetirizine",
	// 	Type:    nam,
	// 	MFD:     time.Date(2022, 10, 24, 12, 0, 0, 0, time.UTC),
	// 	EXP:     time.Date(2023, 10, 24, 12, 0, 0, 0, time.UTC),
	// 	Storage: b2,
	// 	Amount:  100,
	// }
	// db.Model(&Medicine{}).Create(&medicine_4)

	// patient_1 := Patient{
	// 	PID:     "P0001",
	// 	Name:    "กิตติมากร",
	// 	Surname: "สอนแก้ว",
	// 	Age:     21,
	// 	Gender:  "หญิง",
	// 	Allergy: "ยาแอสไพริน",
	// }
	// db.Model(&Patient{}).Create(&patient_1)

	// patient_2 := Patient{
	// 	PID:     "P0002",
	// 	Name:    "ยศพล",
	// 	Surname: "จันทะนาม",
	// 	Age:     22,
	// 	Gender:  "ชาย",
	// 	Allergy: "ไม่แพ้ยาใดๆ",
	// }

	// db.Model(&Patient{}).Create(&patient_2)
	// patient_3 := Patient{
	// 	PID:     "P0003",
	// 	Name:    "กฤษฎา",
	// 	Surname: "น้อยผา",
	// 	Age:     22,
	// 	Gender:  "ชาย",
	// 	Allergy: "ไม่แพ้ยาใดๆ",
	// }

	// db.Model(&Patient{}).Create(&patient_3)

	// prescription_1 := Prescription{
	// 	PrescriptionID: "P00001",
	// 	Symptom:        "ไข้หวัด",
	// 	Case_Time:      time.Now(),
	// 	Employee:       e1,
	// 	Medicine:       medicine_1,
	// 	Patient:        patient_1,
	// }
	// db.Model(&Prescription{}).Create(&prescription_1)

	// prescription_2 := Prescription{
	// 	PrescriptionID: "P00002",
	// 	Symptom:        "ปวดหัว",
	// 	Case_Time:      time.Now(),
	// 	Employee:       e1,
	// 	Medicine:       medicine_1,
	// 	Patient:        patient_2,
	// }
	// db.Model(&Prescription{}).Create(&prescription_2)

	// prescription_3 := Prescription{
	// 	PrescriptionID: "P00003",
	// 	Symptom:        "แผลถลอก",
	// 	Case_Time:      time.Now(),
	// 	Employee:       e1,
	// 	Medicine:       medicine_3,
	// 	Patient:        patient_2,
	// }
	// db.Model(&Prescription{}).Create(&prescription_3)

	// //Medicine
	// medicine_1 := Medicine{
	// 	Name:    "Paracetamol",
	// 	Type:    med,
	// 	MFD:     time.Date(2022, 8, 28, 12, 0, 0, 0, time.UTC),
	// 	EXP:     time.Date(2023, 8, 28, 12, 0, 0, 0, time.UTC),
	// 	Storage: b1,
	// 	Amount:  100,
	// }
	// db.Model(&Medicine{}).Create(&medicine_1)

	// medicine_2 := Medicine{
	// 	Name:    "Menopain",
	// 	Type:    med,
	// 	MFD:     time.Date(2022, 8, 30, 12, 0, 0, 0, time.UTC),
	// 	EXP:     time.Date(2023, 8, 30, 12, 0, 0, 0, time.UTC),
	// 	Storage: b1,
	// 	Amount:  100,
	// }
	// db.Model(&Medicine{}).Create(&medicine_2)

	// medicine_3 := Medicine{
	// 	Name:    "b-derm",
	// 	Type:    med,
	// 	MFD:     time.Date(2022, 8, 30, 12, 0, 0, 0, time.UTC),
	// 	EXP:     time.Date(2023, 8, 30, 12, 0, 0, 0, time.UTC),
	// 	Storage: b2,
	// 	Amount:  100,
	// }

	// db.Model(&Medicine{}).Create(&medicine_3)

	// medicine_4 := Medicine{
	// 	Name:    "Cetirizine",
	// 	Type:    nam,
	// 	MFD:     time.Date(2022, 10, 24, 12, 0, 0, 0, time.UTC),
	// 	EXP:     time.Date(2023, 10, 24, 12, 0, 0, 0, time.UTC),
	// 	Storage: b2,
	// 	Amount:  100,
	// }
	// db.Model(&Medicine{}).Create(&medicine_4)

	// patient_1 := Patient{
	// 	PID:     "P0001",
	// 	Name:    "กิตติมากร",
	// 	Surname: "สอนแก้ว",
	// 	Age:     21,
	// 	Gender:  "หญิง",
	// 	Allergy: "ยาแอสไพริน",
	// }
	// db.Model(&Patient{}).Create(&patient_1)

	// patient_2 := Patient{
	// 	PID:     "P0002",
	// 	Name:    "ยศพล",
	// 	Surname: "จันทะนาม",
	// 	Age:     22,
	// 	Gender:  "ชาย",
	// 	Allergy: "ไม่แพ้ยาใดๆ",
	// }

	// db.Model(&Patient{}).Create(&patient_2)
	// patient_3 := Patient{
	// 	PID:     "P0003",
	// 	Name:    "กฤษฎา",
	// 	Surname: "น้อยผา",
	// 	Age:     22,
	// 	Gender:  "ชาย",
	// 	Allergy: "ไม่แพ้ยาใดๆ",
	// }

	// db.Model(&Patient{}).Create(&patient_3)

	// prescription_1 := Prescription{
	// 	PrescriptionID: "P00001",
	// 	Symptom:        "ไข้หวัด",
	// 	Case_Time:      time.Now(),
	// 	Employee:       e1,
	// 	Medicine:       medicine_1,
	// 	Patient:        patient_1,
	// }
	// db.Model(&Prescription{}).Create(&prescription_1)

	// prescription_2 := Prescription{
	// 	PrescriptionID: "P00002",
	// 	Symptom:        "ปวดหัว",
	// 	Case_Time:      time.Now(),
	// 	Employee:       e1,
	// 	Medicine:       medicine_1,
	// 	Patient:        patient_2,
	// }
	// db.Model(&Prescription{}).Create(&prescription_2)

	// prescription_3 := Prescription{
	// 	PrescriptionID: "P00003",
	// 	Symptom:        "แผลถลอก",
	// 	Case_Time:      time.Now(),
	// 	Employee:       e1,
	// 	Medicine:       medicine_3,
	// 	Patient:        patient_2,
	// }
	// db.Model(&Prescription{}).Create(&prescription_3)

	// 	TotalPrice: 100,
	// 	Receive: 500,
	// 	Refund: 400,

	// })
}
