package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Item struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Time struct {
	ID     uint   `json:"id"`
	Nome   string `json:"nome"`
	Cor    string `json:"cor"`
	Pontos int    `json:"pontos"`
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	db1, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db1.AutoMigrate(&Item{})

	db2, err := gorm.Open(sqlite.Open("times.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db2.AutoMigrate(&Time{})

	app.Get("/items", func(c *fiber.Ctx) error {
		var items []Item
		db1.Find(&items)
		return c.JSON(items)
	})

	app.Get("/itemspar", func(c *fiber.Ctx) error {
		var items []Item
		var itemsPar []Item

		db1.Find(&items)

		for _, item := range items {
			if item.ID%2 == 0 {
				itemsPar = append(itemsPar, item)
			}
		}

		return c.JSON(itemsPar)
	})

	app.Get("/times", func(c *fiber.Ctx) error{
		var times []Time
		var times10 []Time
		db2.Find(&times)
		for _, time:= range times {
			if time.ID < 11 {
				times10 = append(times10, time)
			}
		}
		return c.JSON(times10)
	})

	log.Fatal(app.Listen(":8080"))
}
