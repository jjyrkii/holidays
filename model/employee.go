package model

type Employee struct {
	ID          uint `gorm:"primarykey"`
	FirstName   string
	LastName    string
	Street      string
	HouseNumber uint
	ZipCode     uint
	City        string
	Holidays    []Holiday
}
