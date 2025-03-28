package controllers

import (
	"time"

	req "github.com/elginbrian/ELDERWISE-BE/pkg/dto/request"
	res "github.com/elginbrian/ELDERWISE-BE/pkg/dto/response"
	"github.com/gofiber/fiber/v2"
)

func RegisterHandler(c *fiber.Ctx) error {
	var req req.RegisterRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
	}

	resData := res.RegisterResponseDTO{
		UserID:    "dummy-user-id",
		Email:     req.Email,
		CreatedAt: time.Now(),
	}

	return c.Status(fiber.StatusCreated).JSON(res.ResponseWrapper{
		Success: true,
		Message: "Registration successful",
		Data:    resData,
	})
}

func LoginHandler(c *fiber.Ctx) error {
	var req req.LoginRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
	}

	resData := res.LoginResponseDTO{
		Token: "dummy-jwt-token",
	}

	return c.Status(fiber.StatusOK).JSON(res.ResponseWrapper{
		Success: true,
		Message: "Login successful",
		Data:    resData,
	})
}
