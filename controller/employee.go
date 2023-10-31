package controller

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jjyrkii/holidays/model"
	"github.com/jjyrkii/holidays/utils"
)

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
