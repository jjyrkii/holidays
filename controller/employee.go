package controller

import (
	"net/http"

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
