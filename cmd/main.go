package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/YanYanUcstt/go-fiber-api/pkg/common/config"
	"github.com/YanYanUcstt/go-fiber-api/pkg/common/db"
)

func main(){

	// Load Config file
	c,err := config.LoadConfig()

	if err != nil{
		log.Fatalln("Failed at config",err)
	}
	// Create Fiber
	app := fiber.New()

	// Initialize database
	db.Init(c.DBUrl)

	app.Get("/",func (ctx *fiber.Ctx) error {

		return ctx.Status(fiber.StatusOK).SendString("Hello Go!!")
	})

	app.Listen(c.Port)

}