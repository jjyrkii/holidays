package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jjyrkii/holidays/controller"
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

	h := model.Holiday{
		Start:      time.Time{},
		End:        time.Time{},
		EmployeeID: 1,
	}
	db.Save(&h)
}

func main() {
	// web server instance
	app := fiber.New()

	// route definitions
	e := app.Group("/employee")
	e.Get("/", controller.GetAllEmployees)
	e.Get("/:id", controller.GetEmployeeById)
	e.Delete("/:id", controller.DeleteEmployee)
	e.Post("/", controller.CreateEmployee)
	e.Put("/:id", controller.UpdateEmployee)

	h := app.Group("/holiday")
	h.Get("/", controller.GetAllHolidays)
	h.Get("/:id", controller.GetHolidayById)
	h.Post("/", controller.CreateHoliday)
	h.Delete("/:id", controller.DeleteHoliday)
	h.Get("/employee/:id", controller.GetHolidaysForEmployee)
	e.Put("/:id", controller.UpdateHoliday)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
