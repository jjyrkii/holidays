package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jjyrkii/holidays/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Flo lügt

var db *gorm.DB
var err error

func getProduct(ctx *fiber.Ctx) error {
	var employee model.Employee
	db.First(&employee)
	return ctx.JSON(employee)
}

func main() {
	// db connection
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not establish database connection")
	}

	// db migration
	err = db.AutoMigrate(&model.Employee{}, &model.Address{})
	if err != nil {
		panic(err)
	}

	// create
	db.Create(&model.Employee{
		FirstName: "Tim",
		LastName:  "Pfeiffer",
		Address: model.Address{
			Street:      "Baentschstraße",
			HouseNumber: 11,
			ZipCode:     55122,
			City:        "Mainz",
		},
	})

	// read
	app := fiber.New()

	app.Get("/", getProduct)

	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
