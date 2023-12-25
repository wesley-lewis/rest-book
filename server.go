package main

import (
	"log"
	"os"
	"rest-book/model"
	"rest-book/storage"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	address			string 
	store			storage.Store // Storage Dependency
}

func NewServer(address string ) *Server {
	uri := os.Getenv("MONGO_URI")
	store := storage.NewMongoStore(uri)
	store.RestaurantCollection("rest-book", "restaurant_details")

	return &Server {
		address: address, 
		store: store,
		// store: storage.NewMemoryStore(),
	}
}

func (s *Server) Start() {
	app := fiber.New()
	v1 := app.Group("api/v1")

	// Handlers
	v1.Get("/restaurant", s.GetAllRestaurantDetails)
	v1.Get("/restaurant/:id", s.GetRestaurantDetailsById)
	v1.Put("/restaurant/:id", s.UpdateRestaurantDetails)
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

func(s *Server) GetAllRestaurantDetails(c *fiber.Ctx) error {
	rests, err := s.store.GetAllRestaurantDetails()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Send([]byte(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(rests)
}

func(s *Server) AddRestaurantDetails(c *fiber.Ctx) error {
	rest := &model.Restaurant{}

	if err := c.BodyParser(rest); err != nil {
		return c.Status(fiber.StatusBadRequest).Send([]byte(err.Error()))
	}
	if err := s.store.AddRestaurantDetails(rest); err != nil {
		return c.Status(fiber.StatusBadRequest).Send([]byte(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(rest)
}

func(s *Server) UpdateRestaurantDetails(c *fiber.Ctx) error {
	id := c.Params("id")

	rest := &model.Restaurant{}
	if err := c.BodyParser(rest); err != nil {
		return c.Status(fiber.StatusBadRequest).Send([]byte(err.Error()))
	}

	if err := s.store.UpdateRestaurantDetails(id, rest); err != nil {
		return c.Status(fiber.StatusBadRequest).Send([]byte(err.Error()))
	}

	return c.Status(fiber.StatusOK).Send([]byte("success"))
}
