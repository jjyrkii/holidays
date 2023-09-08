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
}

func main() {
	// create
	// db.Create(&model.Employee{
	//	FirstName: "aösjkdlhfö",
	//	LastName:  "aösdfösd",
	//	Address: model.Address{
	//		Street:      "ddddd",
	//		HouseNumber: 48,
	//		ZipCode:     53300,
	//		City:        "berlin",
	//	},
	// })

	// web server instance
	app := fiber.New()

	// /employee endpoints
	employee := app.Group("/employee", controller.GetAllEmployees)
	employee.Get("/:id", controller.GetEmployeeById)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
