package controllers

import (
	"time"

	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"github.com/elginbrian/ELDERWISE-BE/internal/services"
	req "github.com/elginbrian/ELDERWISE-BE/pkg/dto/request"
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

func (c *LocationHistoryController) CreateLocationHistory(ctx *fiber.Ctx) error {
	var historyReq req.LocationHistoryRequestDTO
	if err := ctx.BodyParser(&historyReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
	}

	history := models.LocationHistory{
		ElderID:     historyReq.ElderID,
		CaregiverID: historyReq.CaregiverID,
		CreatedAt:   historyReq.CreatedAt,
	}

	if history.CreatedAt.IsZero() {
		history.CreatedAt = time.Now()
	}

	createdHistory, err := c.service.CreateLocationHistory(history)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to create location history",
			Error:   err.Error(),
		})
	}

	responseData := res.LocationHistoryResponseDTO{
		LocationHistory: createdHistory,
	}

	return ctx.Status(fiber.StatusCreated).JSON(res.ResponseWrapper{
		Success: true,
		Message: "Location history created successfully",
		Data:    responseData,
	})
}

func (c *LocationHistoryController) AddLocationPoint(ctx *fiber.Ctx) error {
	var pointReq req.LocationHistoryPointRequestDTO
	if err := ctx.BodyParser(&pointReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
	}

	historyID := ctx.Params("location_history_id")
	if historyID != "" && pointReq.LocationHistoryID == "" {
		pointReq.LocationHistoryID = historyID
	}

	point := models.LocationHistoryPoint{
		LocationHistoryID: pointReq.LocationHistoryID,
		Latitude:          pointReq.Latitude,
		Longitude:         pointReq.Longitude,
		Timestamp:         pointReq.Timestamp,
	}

	if point.Timestamp.IsZero() {
		point.Timestamp = time.Now()
	}

	err := c.service.AddLocationPoint(point)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to add location point",
			Error:   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(res.ResponseWrapper{
		Success: true,
		Message: "Location point added successfully",
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





