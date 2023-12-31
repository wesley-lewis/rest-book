package api

import (
	"github.com/gofiber/fiber/v2"
	"rest-book/model"
)

func (s *Server) CreateUser(c *fiber.Ctx) error {
	user := &model.User{}

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	id, err := s.Store.AddUser(user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).SendString(id.Hex())
}

func(s *Server) GetUsers(c *fiber.Ctx) error {
	users, err := s.Store.GetUsers()	
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(users)
}
