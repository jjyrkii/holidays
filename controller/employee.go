package controller

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jjyrkii/holidays/model"
	"github.com/jjyrkii/holidays/utils"
)

// GetAllEmployees godoc
// @Summary Get all employees
// @Description Get a list of all employees with their associated holiday records
// @Tags read
// @Produce json
// @Success 200
// @Router /employee [get]
func GetAllEmployees(c *fiber.Ctx) error {
	db, err := utils.GetDB()
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	var employees []model.Employee
	err = db.Model(&model.Employee{}).Preload("Holidays").Find(&employees).Error
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.JSON(employees)
}

// GetEmployeeById godoc
// @Summary Get a single employee
// @Description Get data for a single employee with their associated holiday records
// @Tags read
// @Param id path int true "The employees id"
// @Produce json
// @Success 200
// @Router /employee/{id} [get]
func GetEmployeeById(c *fiber.Ctx) error {
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

	return c.JSON(employee)
}

// DeleteEmployee godoc
// @Summary Delete an employee
// @Description Delete data for a single employee
// @Tags delete
// @Param id path int true "The employees id"
// @Success 200
// @Router /employee/{id} [delete]
func DeleteEmployee(c *fiber.Ctx) error {
	db, err := utils.GetDB()
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	err = db.Delete(&model.Employee{}, id).Error
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.SendStatus(http.StatusOK)
}

// CreateEmployee godoc
// @Summary Create employee
// @Description Create a new employee
// @Tags create
// @Accept json
// @Param request body model.Employee true "JSON data to create a new employee"
// @Success 200
// @Router /employee [post]
func CreateEmployee(c *fiber.Ctx) error {
	db, err := utils.GetDB()
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	var employee model.Employee
	err = c.BodyParser(&employee)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	db.Save(&employee)

	return c.SendStatus(http.StatusOK)
}

// UpdateEmployee godoc
// @Summary Update employee
// @Description Update the values for an employee
// @Tags update
// @Accept json
// @Param id path int true "Employees id"
// @Param request body model.Employee true "Updated JSON data for an employee"
// @Success 200
// @Router /employee/{id} [put]
func UpdateEmployee(c *fiber.Ctx) error {
	db, err := utils.GetDB()
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	var employee model.Employee
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	err = c.BodyParser(&employee)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	db.Where("id = ?", id).Updates(&employee)

	return c.SendStatus(http.StatusOK)
}
