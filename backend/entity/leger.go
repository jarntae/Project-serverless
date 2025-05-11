package entity

import (
	"gorm.io/gorm"
)

type Ledger struct {
	gorm.Model
	Description string 
	Amount  float64 `valid:"required~Amount is required"`
	Date    string 	`valid:"required~Date is required,dateFormat~Invalid date format"`

	MemberID	uint		`valid:"required~MemberID is required"`
	Member		Member		`gorm:"foreignKey: member_id"`

	Money_Type_ID	uint			`valid:"required~MoneyTypeID is required"`
	Money_Type		Money_Type		`gorm:"foreignKey: money_type_id"`

	Category_Type_ID	uint	
	Category_Type		Category_Type		`gorm:"foreignKey: category_type_id"`
}