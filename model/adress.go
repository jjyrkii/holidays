package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Street      string
	HouseNumber uint
	ZipCode     uint
	City        string
	EmployeeID  uint
}
