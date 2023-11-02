package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/jjyrkii/holidays/controller"
	_ "github.com/jjyrkii/holidays/docs"
	"github.com/jjyrkii/holidays/model"
	"github.com/jjyrkii/holidays/utils"
)

func init() {
	// db connection
	db, err := utils.GetDB()
	if err != nil {
		log.Fatal("Could not establish database connection")
	}

	// db migration
	err = db.AutoMigrate(&model.Employee{}, &model.Holiday{})
	if err != nil {
		log.Fatal("Database migration failed")
	}
}

// @title Holidays
// @version 1.0
// @description This is a service to manage Employees and their holidays
// @host 127.0.0.1:3000
// @BasePath /
func main() {
	// web server instance
	app := fiber.New()

	// route definitions for managing employees
	e := app.Group("/employee")
	e.Get("/", controller.GetAllEmployees)
	e.Get("/:id", controller.GetEmployeeById)
	e.Delete("/:id", controller.DeleteEmployee)
	e.Post("/", controller.CreateEmployee)
	e.Put("/:id", controller.UpdateEmployee)

	// route definitions for managing holidays
	h := app.Group("/holiday")
	h.Get("/", controller.GetAllHolidays)
	h.Get("/:id", controller.GetHolidayById)
	h.Post("/", controller.CreateHoliday)
	h.Delete("/:id", controller.DeleteHoliday)
	h.Get("/employee/:id", controller.GetHolidaysForEmployee)
	e.Put("/:id", controller.UpdateHoliday)

	// route definitions for managing file exports
	x := app.Group("export")
	x.Get("/", controller.FullExport)
	x.Get("employee/:id", controller.ExportEmployee)
	x.Get("holiday/:id", controller.ExportHoliday)

	// route for the documentation
	app.Get("/docs/*", swagger.HandlerDefault)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
