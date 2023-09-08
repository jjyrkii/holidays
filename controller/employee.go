package controller

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jjyrkii/holidays/model"
	"github.com/jjyrkii/holidays/utils"
)

func GetAllEmployees(ctx *fiber.Ctx) error {
	db, err := utils.GetDB()
	if err != nil {
		return ctx.SendStatus(http.StatusInternalServerError)
	}
	var employees []model.Employee
	err = db.Model(&model.Employee{}).Preload("Address").Find(&employees).Error
	if err != nil {
		return ctx.SendStatus(http.StatusInternalServerError)
	}
	return ctx.JSON(employees)
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
	err = db.Preload("Address").First(&employee, id).Error
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
		return c.SendString(err.Error())
	}
	return c.SendStatus(http.StatusOK)
}
