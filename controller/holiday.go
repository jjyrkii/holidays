package controller

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jjyrkii/holidays/model"
	"github.com/jjyrkii/holidays/utils"
)

// CreateHoliday godoc
// @Summary Create holiday
// @Description Create a new holiday
// @Tags create
// @Accept json
// @Param request body model.Holiday true "JSON data to create a new holiday"
// @Success 200
// @Router /holiday [post]
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

	if db.Find(&model.Employee{}, holiday.EmployeeID).RowsAffected == 0 {
		return c.SendStatus(http.StatusBadRequest)
	}

	err = db.Save(&holiday).Error
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.SendStatus(http.StatusOK)
}

// GetAllHolidays godoc
// @Summary Get all holidays
// @Description Get a list of all holidays
// @Tags read
// @Produce json
// @Success 200
// @Router /holiday [get]
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

// GetHolidayById godoc
// @Summary Get a single holiday
// @Description Get data for a single holiday
// @Tags read
// @Param id path int true "The holidays id"
// @Produce json
// @Success 200 {object}
// @Router /holiday/{id} [get]
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

// GetHolidaysForEmployee godoc
// @Summary Get a holidays for employee
// @Description Get all the associated holidays for a specific employee
// @Tags read
// @Param id path int true "The employees id"
// @Produce json
// @Success 200
// @Router /holiday/employee/{id} [get]
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

// DeleteHoliday godoc
// @Summary Delete a holiday
// @Description Delete data for a single holiday
// @Tags delete
// @Param id path int true "The holidays id"
// @Success 200
// @Router /holiday/{id} [delete]
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

// UpdateHoliday godoc
// @Summary Update holiday
// @Description Update the values for a holiday
// @Tags update
// @Accept json
// @Param id path int true "Holidays id"
// @Param request body model.Holiday true "Updated JSON data for an employee"
// @Success 200
// @Router /holiday/{id} [put]
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
