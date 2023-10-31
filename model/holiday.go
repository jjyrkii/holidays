package model

import (
	"time"
)

type Holiday struct {
	ID         uint `gorm:"primarykey"`
	Start      time.Time
	End        time.Time
	Approved   bool
	EmployeeID uint
}
