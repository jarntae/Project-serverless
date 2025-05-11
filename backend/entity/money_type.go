package entity

import (
	"gorm.io/gorm"
)

type Money_Type struct {
	gorm.Model
	Type string `valid:"required~Type is required"`

	Ledger []Ledger `gorm:"foreignKey:money_type_id"`
}