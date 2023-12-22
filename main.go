package main 

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", helloWorldHandler)
	fmt.Println("rest book")
	if err := app.Listen(":8000"); err != nil {
		log.Fatal(err)
	}
}

func helloWorldHandler(c *fiber.Ctx) error {
	return c.SendString("Hello world\n")
}
