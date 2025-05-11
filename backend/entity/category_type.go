package entity

import (
	"gorm.io/gorm"
)

type Category_Type struct {
	gorm.Model
	Category_Name string `valid:"required~Category Name is required"`

	Ledger []Ledger `gorm:"foreignKey:category_type_id"`
}