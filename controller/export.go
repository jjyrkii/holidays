package controller

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jjyrkii/holidays/model"
	"github.com/jjyrkii/holidays/utils"
)

func FullExport(c *fiber.Ctx) error {
	if _, err := os.Stat("export.json"); err == nil {
		if err := os.Remove("export.json"); err != nil {
			return c.SendStatus(http.StatusInternalServerError)
		}
	}

	db, err := utils.GetDB()
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	var employees []model.Employee
	err = db.Model(&model.Employee{}).Preload("Holidays").Find(&employees).Error
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	content, err := json.Marshal(employees)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	err = os.WriteFile("export.json", content, 0644)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.SendFile("export.json")
}

func ExportEmployee(c *fiber.Ctx) error {
	db, err := utils.GetDB()
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	var employee model.Employee
	err = db.Model(&employee).Preload("Holidays").First(&employee, id).Error
	if err != nil {
		return c.SendStatus(http.StatusNotFound)
	}

	var file *os.File
	if file, err = employee.Export(); err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	} else {

		return c.SendFile(file.Name())
	}
}

func ExportHoliday(c *fiber.Ctx) error {
	db, err := utils.GetDB()
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	var holiday model.Holiday
	err = db.Model(&holiday).First(&holiday, id).Error
	if err != nil {
		return c.SendStatus(http.StatusNotFound)
	}

	var file *os.File
	if file, err = holiday.Export(); err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	} else {

		return c.SendFile(file.Name())
	}
}
