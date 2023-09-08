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
	//// db connection
	//db, err := utils.GetDB()
	//if err != nil {
	//	log.Fatal("Could not establish database connection")
	//}
	//
	//// db migration
	//err = db.AutoMigrate(&model.Employee{}, &model.Address{})
	//if err != nil {
	//	panic(err)
	//}

	// create
	//db.Create(&model.Employee{
	//	FirstName: "aösjkdlhfö",
	//	LastName:  "aösdfösd",
	//	Address: model.Address{
	//		Street:      "ddddd",
	//		HouseNumber: 48,
	//		ZipCode:     53300,
	//		City:        "berlin",
	//	},
	//})

	app := fiber.New()

	app.Get("/", controller.GetAllEmployees)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
