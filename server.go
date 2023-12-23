package main

import (
	"log"
	"rest-book/storage"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	address			string 
	store			storage.Store
}

func NewServer(address string ) *Server {
	return &Server {
		address: address, 
		store : storage.NewMongoStore(""),
	}
}

func (s *Server) Start() {
	app := fiber.New()

	// Handlers
	app.Get("/", helloWorldHandler)
	app.Get("/restaurant/:id", GetRestaurantDetailsById)

	if err := app.Listen(":8000"); err != nil {
		log.Fatal(err)
	}
}

func helloWorldHandler(c *fiber.Ctx) error {
	return c.SendString("Hello world\n")
}

func GetRestaurantDetailsById(c *fiber.Ctx) error {
	id := c.Params("id")

	return c.Status(fiber.StatusOK).Send([]byte(id + " details"))
}
