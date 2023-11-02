package model

import (
	"encoding/json"
	"os"
)

type Employee struct {
	ID          uint   `gorm:"primarykey"`
	FirstName   string `binding:"required"`
	LastName    string `binding:"required"`
	Street      string `binding:"required"`
	HouseNumber uint   `binding:"required"`
	ZipCode     uint   `binding:"required"`
	City        string `binding:"required"`
	Holidays    []Holiday
}

func (e Employee) Export() (*os.File, error) {
	var (
		file    *os.File
		content []byte
		err     error
	)

	if _, err := os.Stat("employee_export.json"); err == nil {
		if err := os.Remove("employee_export.json"); err != nil {
			return nil, err
		}
	}

	if content, err = json.Marshal(e); err != nil {
		return nil, err
	}

	if file, err = os.Create("employee_export.json"); err != nil {
		return nil, err
	}

	if _, err = file.Write(content); err != nil {
		return nil, err
	}

	if err = file.Close(); err != nil {
		return nil, err
	}

	return file, nil
}
