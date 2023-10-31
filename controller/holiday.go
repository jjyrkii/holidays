package controller

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jjyrkii/holidays/model"
	"github.com/jjyrkii/holidays/utils"
)

func CreateHoliday(c *fiber.Ctx) error {
	db, err := utils.GetDB()
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	var holiday model.Holiday
	err = c.BodyParser(&holiday)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	err = db.Create(&holiday).Error
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.SendStatus(http.StatusOK)
}

func GetAllHolidays(c *fiber.Ctx) error {
	db, err := utils.GetDB()
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	var holidays []model.Holiday
	err = db.Model(&model.Holiday{}).Find(&holidays).Error
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.JSON(holidays)
}

func GetHolidayById(c *fiber.Ctx) error {
	db, err := utils.GetDB()
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	var holiday model.Holiday
	err = db.First(&holiday, id).Error
	if err != nil {
		return c.SendStatus(http.StatusNotFound)
	}

	return c.JSON(holiday)
}

func GetHolidaysForEmployee(c *fiber.Ctx) error {
	db, err := utils.GetDB()
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	var holidays []model.Holiday
	err = db.Where("employee_id = ?", id).Find(&holidays).Error
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.JSON(holidays)
}

func DeleteHoliday(c *fiber.Ctx) error {
	db, err := utils.GetDB()
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	err = db.Delete(&model.Holiday{}, id).Error
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.SendStatus(http.StatusOK)
}

func UpdateHoliday(c *fiber.Ctx) error {
	db, err := utils.GetDB()
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	var holiday model.Holiday
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	err = c.BodyParser(&holiday)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	err = db.Where("id = ?", id).Updates(&holiday).Error
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.SendStatus(http.StatusOK)
}
