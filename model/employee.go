package model

import (
	"encoding/json"
	"os"
)

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

func (e Employee) Export() (*os.File, error) {
	var (
		file    *os.File
		content []byte
		err     error
	)

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
