package main

import (
	"log"

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
	err = db.AutoMigrate(&model.Employee{})
	if err != nil {
		log.Fatal("Database migration failed")
	}
}

func main() {
	// create

	// web server instance
	app := fiber.New()

	e := app.Group("/employee")
	e.Get("/", controller.GetAllEmployees)
	e.Get("/:id", controller.GetEmployeeById)
	e.Delete("/:id", controller.DeleteEmployee)
	e.Post("/", controller.CreateEmployee)
	e.Put("/:id", controller.UpdateEmployee)
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
