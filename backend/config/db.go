package config

import (
	"fmt"
	"github.com/cpe-nuntawut/assignment-6-bobox/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func ConnectionDB() {
	database, err := gorm.Open(sqlite.Open("ServerCloud.db?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connected database")
	db = database
}

func SetupDatabase() {
	db.AutoMigrate(
		&entity.Member{},
		&entity.Ledger{},
		&entity.Money_Type{},
		&entity.Category_Type{},
	)
	
	Income := entity.Money_Type{Type: "รายรับ"}
	Outcome := entity.Money_Type{Type: "รายจ่าย"}

	Food_drink := entity.Category_Type{Category_Name: "มื้ออาหาร/เครื่องดื่ม"}
	Shopping := entity.Category_Type{Category_Name: "ช้อปปิ้ง"}
	Family_Personal := entity.Category_Type{Category_Name: "ครอบครัว/ส่วนตัว"}
	Invest := entity.Category_Type{Category_Name: "ออมเงิน/ลงทุน"}
	Bill := entity.Category_Type{Category_Name: "ชำระบิล"}
	Entertainment := entity.Category_Type{Category_Name: "บันเทิง"}
	Gift_give := entity.Category_Type{Category_Name: "ของขวัญ/บริจาค"}
	Transport := entity.Category_Type{Category_Name: "ค่าเดินทาง"}
	Education := entity.Category_Type{Category_Name: "การศึกษา"}
	Hotel_Travel := entity.Category_Type{Category_Name: "โรงแรม/ท่องเที่ยว"}
	Insure := entity.Category_Type{Category_Name: "ประกัน"}
	Withdraw := entity.Category_Type{Category_Name: "ถอนเงิน"}
	Healty_Beauty := entity.Category_Type{Category_Name: "สุขภาพ/ความงาม"}
	Pet := entity.Category_Type{Category_Name: "สัตว์เลี้ยง"}
	Home := entity.Category_Type{Category_Name: "บ้าน/ที่พักอาศัย"}
	Repay := entity.Category_Type{Category_Name: "โอนตืนเงิน"}
	Parcel := entity.Category_Type{Category_Name: "พัสดุ"}
	Etc := entity.Category_Type{Category_Name: "อื่นๆ"}

	Testledger := entity.Ledger{Description: "เงินจากทางบ้าน",Amount: 1111,MemberID: 1,Money_Type_ID: 1}

	db.FirstOrCreate(&Income, &entity.Money_Type{Type: "รายรับ"})
	db.FirstOrCreate(&Outcome, &entity.Money_Type{Type: "รายจ่าย"})

	db.FirstOrCreate(&Food_drink, &entity.Category_Type{Category_Name: "มื้ออาหาร/เครื่องดื่ม"})
	db.FirstOrCreate(&Shopping, &entity.Category_Type{Category_Name: "ช้อปปิ้ง"})
	db.FirstOrCreate(&Family_Personal, &entity.Category_Type{Category_Name: "ครอบครัว/ส่วนตัว"})
	db.FirstOrCreate(&Invest, &entity.Category_Type{Category_Name: "ออมเงิน/ลงทุน"})
	db.FirstOrCreate(&Bill, &entity.Category_Type{Category_Name: "ชำระบิล"})
	db.FirstOrCreate(&Entertainment, &entity.Category_Type{Category_Name: "บันเทิง"})
	db.FirstOrCreate(&Gift_give, &entity.Category_Type{Category_Name: "ของขวัญ/บริจาค"})
	db.FirstOrCreate(&Transport, &entity.Category_Type{Category_Name: "ค่าเดินทาง"})
	db.FirstOrCreate(&Education, &entity.Category_Type{Category_Name: "การศึกษา"})
	db.FirstOrCreate(&Hotel_Travel, &entity.Category_Type{Category_Name: "โรงแรม/ท่องเที่ยว"})
	db.FirstOrCreate(&Insure, &entity.Category_Type{Category_Name: "ประกัน"})
	db.FirstOrCreate(&Withdraw, &entity.Category_Type{Category_Name: "ถอนเงิน"})
	db.FirstOrCreate(&Healty_Beauty, &entity.Category_Type{Category_Name: "สุขภาพ/ความงาม"})
	db.FirstOrCreate(&Pet, &entity.Category_Type{Category_Name: "สัตว์เลี้ยง"})
	db.FirstOrCreate(&Home, &entity.Category_Type{Category_Name: "บ้าน/ที่พักอาศัย"})
	db.FirstOrCreate(&Repay, &entity.Category_Type{Category_Name: "โอนเงินคืน"})
	db.FirstOrCreate(&Parcel, &entity.Category_Type{Category_Name: "พัสดุ"})
	db.FirstOrCreate(&Etc, &entity.Category_Type{Category_Name: "อื่นๆ"})

	db.FirstOrCreate(&Testledger, &entity.Ledger{Description: "เงินจากทางบ้าน",Amount: 1111,MemberID: 1,Money_Type_ID: 1})

	hashedPassword, _ := HashPassword("12345")

	member := &entity.Member{
		FirstName:  "IT",
		LastName:   "Admin",
		Email:      "ITAdmin@stayease.com",
		Password:   hashedPassword,
	}

	db.FirstOrCreate(member, &entity.Member{
		Email: "ITAdmin@stayease.com",
	})
}