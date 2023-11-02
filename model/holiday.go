package model

import (
	"encoding/json"
	"os"
	"time"
)

type Holiday struct {
	ID         uint      `gorm:"primarykey"`
	Start      time.Time `binding:"required"`
	End        time.Time `binding:"required"`
	Approved   bool      `binding:"required"`
	EmployeeID uint      `binding:"required"`
}

func (h Holiday) Export() (*os.File, error) {
	var (
		file    *os.File
		content []byte
		err     error
	)

	if _, err := os.Stat("holiday_export.json"); err == nil {
		if err := os.Remove("holiday_export.json"); err != nil {
			return nil, err
		}
	}

	if content, err = json.Marshal(h); err != nil {
		return nil, err
	}

	if file, err = os.Create("holiday_export.json"); err != nil {
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
