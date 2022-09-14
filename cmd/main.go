package main

import (
	"log"

	"github.com/YanYanUcstt/go-fiber-api/pkg/books"
	"github.com/YanYanUcstt/go-fiber-api/pkg/common/config"
	"github.com/YanYanUcstt/go-fiber-api/pkg/common/db"
	"github.com/gofiber/fiber/v2"
)

func main() {

	// Load Config file
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	// Create Fiber
	app := fiber.New()

	// Initialize database
	db := db.Init(c.DBUrl)

	//register routes
	books.RegisterRoutes(app, db)
	app.Listen(c.Port)

}
