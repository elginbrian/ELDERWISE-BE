package controllers

import (
	"github.com/elginbrian/ELDERWISE-BE/internal/services"
	res "github.com/elginbrian/ELDERWISE-BE/pkg/dto/response"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) GetUserByID(c *fiber.Ctx) error {
	userID := c.Params("user_id")

	user, err := uc.userService.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(res.ResponseWrapper{
			Success: false,
			Message: "User not found",
			Error:   err.Error(),
		})
	}

	user.Password = ""

	responseData := res.UserResponseDTO{
		User: *user,
	}

	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "User retrieved successfully",
		Data:    responseData,
	})
}

func (uc *UserController) GetUserCaregivers(c *fiber.Ctx) error {
	userID := c.Params("user_id")

	caregivers, err := uc.userService.GetCaregiversByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to retrieve caregivers",
			Error:   err.Error(),
		})
	}

	responseData := res.CaregiversResponseDTO{
		Caregivers: caregivers,
	}

	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Caregivers retrieved successfully",
		Data:    responseData,
	})
}

func (uc *UserController) GetUserElders(c *fiber.Ctx) error {
	userID := c.Params("user_id")

	elders, err := uc.userService.GetEldersByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to retrieve elders",
			Error:   err.Error(),
		})
	}

	responseData := res.EldersResponseDTO{
		Elders: elders,
	}

	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Elders retrieved successfully",
		Data:    responseData,
	})
}

