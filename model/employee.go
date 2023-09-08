package model

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	FirstName string
	LastName  string
	Address   Address
}

type Address struct {
	gorm.Model
	Street      string
	HouseNumber uint
	ZipCode     uint
	City        string
	EmployeeID  uint
}
