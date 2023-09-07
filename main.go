package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jjyrkii/holidays/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func getProduct(ctx *fiber.Ctx) error {
	var product model.Product
	db.First(&product)
	return ctx.SendString(product.Format())
	//return ctx.JSON(product)
}

func main() {
	// db connection
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not establish database connection")
	}

	// db migration
	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		panic(err)
	}

	// create
	db.Create(&model.Product{
		Code:  "D42",
		Price: 100,
	})

	// read
	var product model.Product
	db.First(&product, 1)
	db.First(&product, "code = ?", "D42")

	// update
	db.Model(&product).Update("Price", 200)

	db.Model(&product).Updates(model.Product{Price: 200, Code: product.Code})
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	//delete
	db.Delete(&product, 1)
	app := fiber.New()

	app.Get("/", getProduct)

	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
