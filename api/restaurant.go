package api

import (
	"github.com/gofiber/fiber/v2"
	"rest-book/model"
)

func helloWorldHandler(c *fiber.Ctx) error {
	return c.SendString("Hello world\n")
}

func (s *Server) GetRestaurantDetailsById(c *fiber.Ctx) error {
	id := c.Params("id")
	rest, err := s.Store.GetRestaurantDetails(id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(rest)
}

func (s *Server) GetAllRestaurantDetails(c *fiber.Ctx) error {
	rests, err := s.Store.GetAllRestaurantDetails()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Send([]byte(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(rests)
}

func (s *Server) AddRestaurantDetails(c *fiber.Ctx) error {
	rest := &model.Restaurant{}

	if err := c.BodyParser(rest); err != nil {
		return c.Status(fiber.StatusBadRequest).Send([]byte(err.Error()))
	}
	if err := s.Store.AddRestaurantDetails(rest); err != nil {
		return c.Status(fiber.StatusBadRequest).Send([]byte(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(rest)
}

func (s *Server) UpdateRestaurantDetails(c *fiber.Ctx) error {
	id := c.Params("id")

	rest := &model.Restaurant{}
	if err := c.BodyParser(rest); err != nil {
		return c.Status(fiber.StatusBadRequest).Send([]byte(err.Error()))
	}

	if err := s.Store.UpdateRestaurantDetails(id, rest); err != nil {
		return c.Status(fiber.StatusBadRequest).Send([]byte(err.Error()))
	}

	return c.Status(fiber.StatusOK).Send([]byte("success"))
}

func (s *Server) DeleteRestaurantDetails(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := s.Store.DeleteRestaurantDetails(id); err != nil {
		return c.Status(fiber.StatusOK).Send([]byte(err.Error()))
	}

	return c.Status(fiber.StatusOK).Send([]byte("deleted " + id))
}
