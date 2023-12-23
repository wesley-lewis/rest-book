package main

import (
	"log"
	"rest-book/model"
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
		// store : storage.NewMongoStore(""),
		store: storage.NewMemoryStore(),
	}
}

func (s *Server) Start() {
	app := fiber.New()
	v1 := app.Group("api/v1")

	// Handlers
	v1.Get("/restaurant/:id", s.GetRestaurantDetailsById)
	v1.Post("/restaurant", s.AddRestaurantDetails)

	if err := app.Listen(":8000"); err != nil {
		log.Fatal(err)
	}
}

func helloWorldHandler(c *fiber.Ctx) error {
	return c.SendString("Hello world\n")
}

func ( s *Server) GetRestaurantDetailsById(c *fiber.Ctx) error {
	id := c.Params("id")
	rest, err := s.store.GetRestaurantDetails(id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(rest)
}

func(s *Server) AddRestaurantDetails(c *fiber.Ctx) error {
	rest := &model.Restaurant{}

	if err := c.BodyParser(rest); err != nil {
		return c.Status(fiber.StatusBadRequest).Send([]byte(err.Error()))
	}
	s.store.AddRestaurantDetails(rest)

	return c.Status(fiber.StatusOK).JSON(rest)
}
