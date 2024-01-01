package api

import "github.com/gofiber/fiber/v2"

func (s *Server) AddProducts(c *fiber.Ctx) error {
    return c.Status(fiber.StatusOK).SendString("add products")
}

func(s *Server) GetProducts(c *fiber.Ctx) error {
    return c.Status(fiber.StatusOK).SendString("products")
}
