package api

import (
	"rest-book/model"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) AddItem(c *fiber.Ctx) error {
    product := &model.Product{}

    if err := c.BodyParser(product); err != nil {
        return c.Status(fiber.StatusBadRequest).SendString(err.Error())
    }

    if err := s.Store.AddProduct(product); err != nil {
        return c.Status(fiber.StatusBadRequest).SendString(err.Error())
    }

    return c.Status(fiber.StatusOK).JSON(struct{
        message string
    }{
        message: "added product",
    })
}

func(s *Server) GetAllItems(c *fiber.Ctx) error {
    return c.Status(fiber.StatusOK).SendString("products")
}
