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

// RegisterHandler godoc
// @Summary Register a new user
// @Description Register a new user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param registerRequest body req.RegisterRequestDTO true "Register credentials"
// @Success 201 {object} res.ResponseWrapper{data=res.RegisterResponseDTO} "Registration successful"
// @Failure 400 {object} res.ResponseWrapper "Bad request"
// @Router /auth/register [post]
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

// LoginHandler godoc
// @Summary Login user
// @Description Login with email and password to get a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param loginRequest body req.LoginRequestDTO true "Login credentials"
// @Success 200 {object} res.ResponseWrapper{data=res.LoginResponseDTO} "Login successful"
// @Failure 400 {object} res.ResponseWrapper "Bad request"
// @Failure 401 {object} res.ResponseWrapper "Unauthorized"
// @Router /auth/login [post]
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

// GetCurrentUser godoc
// @Summary Get current user information
// @Description Get the current user information based on JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} res.ResponseWrapper{data=models.User} "User retrieved successfully"
// @Failure 401 {object} res.ResponseWrapper "Unauthorized"
// @Router /auth/me [get]
// @Security Bearer
func (ac *AuthController) GetCurrentUser(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Authorization header is required",
			Error:   "unauthorized",
		})
	}

	user, err := ac.authService.GetUserFromToken(authHeader)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to authenticate",
			Error:   err.Error(),
		})
	}

	user.Password = ""

	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "User retrieved successfully",
		Data:    user,
	})
}



