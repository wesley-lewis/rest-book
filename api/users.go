package api

import (
	"rest-book/model"

	"github.com/gofiber/fiber/v2"
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

func(s *Server) UpdateUser(c *fiber.Ctx) error { user := &model.User{}
    id := c.Params("id")	

    if err := c.BodyParser(user); err != nil {
        return c.Status(fiber.StatusBadRequest).SendString(err.Error())
    }

    if err := s.Store.UpdateUser(id, user); err != nil {
        return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
    }

    return c.Status(fiber.StatusOK).JSON(struct{
        message string
    }{
        message: "all good",
    })
}

func(s *Server) LoginUser(c *fiber.Ctx) error {
    user := &model.UserLogin{}

    if err := c.BodyParser(user); err != nil {
        return c.Status(fiber.StatusBadRequest).SendString(err.Error())
    }

    if user.Email == "" {
        return c.Status(fiber.StatusBadRequest).SendString("bad request")
    }
    dbUser, err := s.Store.GetUserByEmail(user.Email)
    if err != nil {
        return err
    }

    if user.Email != dbUser.Email {
        return c.Status(fiber.StatusNonAuthoritativeInformation).JSON(struct {
            Message string `json:"message"`
            Success bool `json:"success"`
        }{
                Message: "failed",
                Success: false,
        })
    }

    return c.Status(fiber.StatusOK).JSON(struct{
        Message string `json:"message"`
        Success bool `json:"success"`
    }{
        Message: "logged in user",
        Success: true,
    })
}

func(s *Server) DeleteUser(c *fiber.Ctx) error {
    id := c.Params("id")

    if err := s.Store.DeleteUser(id); err != nil {
        return err
    }
    return c.Status(fiber.StatusOK).SendString("Deleted user: " + id)
}
