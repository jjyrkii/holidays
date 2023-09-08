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
	err = db.AutoMigrate(&model.Employee{}, &model.Address{})
	if err != nil {
		log.Fatal("Database migration failed")
	}
	db.Create(&model.Employee{
		FirstName: "aösjkdlhfö",
		LastName:  "aösdfösd",
		Address: model.Address{
			Street:      "ddddd",
			HouseNumber: 47,
			ZipCode:     53299,
			City:        "berlin",
		},
	})
}

func main() {
	// create

	// web server instance
	app := fiber.New()

	e := app.Group("/employee")
	e.Get("/", controller.GetAllEmployees)
	e.Get("/:id", controller.GetEmployeeById)
	e.Delete("/:id", controller.DeleteEmployee)
	// // /employee endpoints
	// app.Get("/employee", controller.GetAllEmployees)
	// app.Get("/employee/:id", controller.GetEmployeeById)
	// app.Get("/employee/:id", controller.DeleteEmployee)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
