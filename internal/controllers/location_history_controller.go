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

func GetLocationHistoryByID(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).JSON(res.ResponseWrapper{
		Success: false,
		Message: "This method is deprecated, please update your routes",
	})
}

func GetLocationHistoryPoints(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).JSON(res.ResponseWrapper{
		Success: false,
		Message: "This method is deprecated, please update your routes",
	})
}



