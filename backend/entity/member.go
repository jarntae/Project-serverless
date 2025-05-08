package entity

import (
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	FirstName string `valid:"required~FirstName is required"`
	LastName  string `valid:"required~LastName is required"`
	Email     string `valid:"required~Email is required, email~Email is invalid"`
	Password  string `valid:"required~Password is required"`
}