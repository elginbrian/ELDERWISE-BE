package controllers

import (
	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"github.com/elginbrian/ELDERWISE-BE/internal/services"
	res "github.com/elginbrian/ELDERWISE-BE/pkg/dto/response"
	"github.com/gofiber/fiber/v2"
)

type AreaController struct {
	service services.AreaService
}

func NewAreaController(service services.AreaService) *AreaController {
	return &AreaController{service}
}

func (ac *AreaController) GetAreaByID(c *fiber.Ctx) error {
	areaID := c.Params("area_id")
	area, err := ac.service.GetAreaByID(areaID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Area not found",
			Error:   err.Error(),
		})
	}

	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Area retrieved successfully",
		Data:    res.AreaResponseDTO{Area: *area},
	})
}

func (ac *AreaController) CreateArea(c *fiber.Ctx) error {
	var area models.Area
	if err := c.BodyParser(&area); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
	}

	if err := ac.service.CreateArea(&area); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to create area",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(res.ResponseWrapper{
		Success: true,
		Message: "Area created successfully",
		Data:    res.AreaResponseDTO{Area: area},
	})
}

func (ac *AreaController) UpdateArea(c *fiber.Ctx) error {
	areaID := c.Params("area_id")
	var area models.Area
	if err := c.BodyParser(&area); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
	}

	if err := ac.service.UpdateArea(areaID, &area); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to update area",
			Error:   err.Error(),
		})
	}

	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Area updated successfully",
		Data:    res.AreaResponseDTO{Area: area},
	})
}

func (ac *AreaController) DeleteArea(c *fiber.Ctx) error {
	areaID := c.Params("area_id")
	if err := ac.service.DeleteArea(areaID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to delete area",
			Error:   err.Error(),
		})
	}

	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Area deleted successfully",
		Data:    areaID,
	})
}

func (ac *AreaController) GetAreasByCaregiver(c *fiber.Ctx) error {
	caregiverID := c.Params("caregiver_id")
	areas, err := ac.service.GetAreasByCaregiver(caregiverID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to retrieve areas",
			Error:   err.Error(),
		})
	}

	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Areas retrieved successfully for caregiver",
		Data:    res.AreasResponseDTO{Areas: areas},
	})
}
