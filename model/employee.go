package model

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	FirstName   string
	LastName    string
	Street      string
	HouseNumber uint
	ZipCode     uint
	City        string
}
