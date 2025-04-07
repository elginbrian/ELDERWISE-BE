package controllers

import (
	"github.com/elginbrian/ELDERWISE-BE/internal/services"
	res "github.com/elginbrian/ELDERWISE-BE/pkg/dto/response"
	"github.com/gofiber/fiber/v2"
)

type LocationHistoryController struct {
	service *services.LocationHistoryService
}

func NewLocationHistoryController(service *services.LocationHistoryService) *LocationHistoryController {
	return &LocationHistoryController{service: service}
}

// GetLocationHistoryByID fetches a location history by its ID
func (c *LocationHistoryController) GetLocationHistoryByID(ctx *fiber.Ctx) error {
	historyID := ctx.Params("location_history_id")

	history, err := c.service.GetLocationHistoryByID(historyID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Location history not found",
			Error:   err.Error(),
		})
	}

	responseData := res.LocationHistoryResponseDTO{
		LocationHistory: history,
	}

	return ctx.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Location history retrieved successfully",
		Data:    responseData,
	})
}

// GetLocationHistoryPoints fetches points for a location history
func (c *LocationHistoryController) GetLocationHistoryPoints(ctx *fiber.Ctx) error {
	historyID := ctx.Params("location_history_id")

	points, err := c.service.GetLocationHistoryPoints(historyID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Location history points not found",
			Error:   err.Error(),
		})
	}

	responseData := res.LocationHistoryPointsResponseDTO{
		Points: points,
	}

	return ctx.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Location history points retrieved successfully",
		Data:    responseData,
	})
}

// GetElderLocationHistory fetches location history for an elder
func (c *LocationHistoryController) GetElderLocationHistory(ctx *fiber.Ctx) error {
	elderID := ctx.Params("elder_id")

	history, err := c.service.GetElderLocationHistory(elderID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Elder location history not found",
			Error:   err.Error(),
		})
	}

	responseData := res.LocationHistoryResponseDTO{
		LocationHistory: history,
	}

	return ctx.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Location history retrieved successfully",
		Data:    responseData,
	})
}

// Legacy function signatures to maintain compatibility with current routes
func GetLocationHistoryByID(c *fiber.Ctx) error {
	// This will be removed after routes are updated
	return c.Status(fiber.StatusNotImplemented).JSON(res.ResponseWrapper{
		Success: false,
		Message: "This method is deprecated, please update your routes",
	})
}

func GetLocationHistoryPoints(c *fiber.Ctx) error {
	// This will be removed after routes are updated
	return c.Status(fiber.StatusNotImplemented).JSON(res.ResponseWrapper{
		Success: false,
		Message: "This method is deprecated, please update your routes",
	})
}

