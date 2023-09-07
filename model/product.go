package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func (p Product) Format() string {
	return fmt.Sprintf("ID: %d, Code: %s, Price: %d", p.ID, p.Code, p.Price)
}
