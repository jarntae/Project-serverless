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
		// &entity.Employee{},
		// &entity.Gender{},
		// &entity.Position{},
		// &entity.Room{},
		// &entity.PetAllow{},
		// &entity.RoomType{},
		// &entity.Booking{},
		// &entity.Facility{},
		// &entity.FacilityStatus{},
		// &entity.FacilityType{},
	)
	
	// GenderMale := entity.Gender{Name: "ชาย"}
	// GenderFemale := entity.Gender{Name: "หญิง"}
	// GenderOther := entity.Gender{Name: "อื่นๆ"}
	
	// PetAllow := entity.PetAllow{Name: "อนุญาต"}
	// PetNotAllow := entity.PetAllow{Name: "ไม่อนุญาต"}
	
	// RoomAvaliable := entity.RoomType{Name: "พร้อมใช้งาน"}
	// RoomNotAvaliable := entity.RoomType{Name: "มีผู้พักอาศัย"}
	// RoomMaintenance := entity.RoomType{Name: "ซ่อมบำรุง"}
	
	// PositionDormitoryAdmin := entity.Position{Name: "ไอที"}
	// PositionDormitoryCaretaker := entity.Position{Name: "ผู้ดูแลหอพัก"}
	// PositionDormitorySecurityGuard := entity.Position{Name: "ผู้ดูแลรักษาความปลอดภัยหอพัก"}
	// PositionDormitoryMaid := entity.Position{Name: "แม่บ้าน"}
	// PositionDormitoryMaintenanceTechnician := entity.Position{Name: "ช่างซ่อมบำรุง"}
	// PositionDormitoryMember := entity.Position{Name: "ผู้พักอาศัย"}

	// FacilityOpen := entity.FacilityStatus{Name: "เปิดใช้งาน"}
	// FacilityClose := entity.FacilityStatus{Name: "ปิดใช้งาน"}
	// FacilityMaintenace := entity.FacilityStatus{Name: "ซ่อมบำรุง"}

	// FacilityGym := entity.FacilityType{Name: "ฟิตเนส"}
	// FacilityMeetRoom := entity.FacilityType{Name: "ห้องประชุม"}

	// db.FirstOrCreate(&GenderMale, &entity.Gender{Name: "ชาย"})
	// db.FirstOrCreate(&GenderFemale, &entity.Gender{Name: "หญิง"})
	// db.FirstOrCreate(&GenderOther, &entity.Gender{Name: "อื่นๆ"})

	// db.FirstOrCreate(&PetAllow, &entity.PetAllow{Name: "อนุญาต"})
	// db.FirstOrCreate(&PetNotAllow, &entity.PetAllow{Name: "ไม่อนุญาต"})

	// db.FirstOrCreate(&RoomAvaliable, &entity.RoomType{Name: "พร้อมใช้งาน"})
	// db.FirstOrCreate(&RoomNotAvaliable, &entity.RoomType{Name: "มีผู้พักอาศัย"})
	// db.FirstOrCreate(&RoomMaintenance, &entity.RoomType{Name: "ซ่อมบำรุง"})

	// db.FirstOrCreate(&PositionDormitoryAdmin, &entity.Position{Name: "ไอที"})
	// db.FirstOrCreate(&PositionDormitoryCaretaker, &entity.Position{Name: "ผู้ดูแลหอพัก"})
	// db.FirstOrCreate(&PositionDormitorySecurityGuard, &entity.Position{Name: "ผู้ดูแลรักษาความปลอดภัยหอพัก"})
	// db.FirstOrCreate(&PositionDormitoryMaid, &entity.Position{Name: "แม่บ้าน"})
	// db.FirstOrCreate(&PositionDormitoryMaintenanceTechnician, &entity.Position{Name: "ช่างซ่อมบำรุง"})
	// db.FirstOrCreate(&PositionDormitoryMember, &entity.Position{Name: "ผู้พักอาศัย"})

	// db.FirstOrCreate(&FacilityOpen, &entity.FacilityStatus{Name: "เปิดใช้งาน"})
	// db.FirstOrCreate(&FacilityClose, &entity.FacilityStatus{Name: "ปิดใช้งาน"})
	// db.FirstOrCreate(&FacilityMaintenace, &entity.FacilityStatus{Name: "ซ่อมบำรุง"})

	// db.FirstOrCreate(&FacilityGym, &entity.FacilityType{Name: "ฟิตเนส"})
	// db.FirstOrCreate(&FacilityMeetRoom, &entity.FacilityType{Name: "ห้องประชุม"})

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