package controllers

import (
	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"github.com/elginbrian/ELDERWISE-BE/internal/services"
	req "github.com/elginbrian/ELDERWISE-BE/pkg/dto/request"
	res "github.com/elginbrian/ELDERWISE-BE/pkg/dto/response"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (ac *AuthController) RegisterHandler(c *fiber.Ctx) error {
	var reqDTO req.RegisterRequestDTO
	if err := c.BodyParser(&reqDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
	}

	user := &models.User{
		Email:    reqDTO.Email,
		Password: reqDTO.Password,
	}

	registeredUser, err := ac.authService.Register(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Registration failed",
			Error:   err.Error(),
		})
	}

	resDTO := res.RegisterResponseDTO{
		UserID:    registeredUser.UserID,
		Email:     registeredUser.Email,
		CreatedAt: registeredUser.CreatedAt,
	}

	return c.Status(fiber.StatusCreated).JSON(res.ResponseWrapper{
		Success: true,
		Message: "Registration successful",
		Data:    resDTO,
	})
}

func (ac *AuthController) LoginHandler(c *fiber.Ctx) error {
	var reqDTO req.LoginRequestDTO
	if err := c.BodyParser(&reqDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
	}

	token, err := ac.authService.Login(reqDTO.Email, reqDTO.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Login failed",
			Error:   err.Error(),
		})
	}

	resDTO := res.LoginResponseDTO{
		Token: token,
	}
	return c.Status(fiber.StatusOK).JSON(res.ResponseWrapper{
		Success: true,
		Message: "Login successful",
		Data:    resDTO,
	})
}
