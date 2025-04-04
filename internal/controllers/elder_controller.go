package controllers

import (
	"time"

	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"github.com/elginbrian/ELDERWISE-BE/internal/services"
	res "github.com/elginbrian/ELDERWISE-BE/pkg/dto/response"
	"github.com/gofiber/fiber/v2"
)

type ElderController struct {
	service services.ElderService
}

func NewElderController(service services.ElderService) *ElderController {
	return &ElderController{service: service}
}

func (ec *ElderController) GetElderByID(c *fiber.Ctx) error {
	elderID := c.Params("elder_id")
	elder, err := ec.service.GetElderByID(elderID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Elder not found",
			Error:   err.Error(),
		})
	}

	responseData := res.ElderResponseDTO{Elder: *elder}
	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Elder retrieved successfully",
		Data:    responseData,
	})
}

func (ec *ElderController) CreateElder(c *fiber.Ctx) error {
	var elder models.Elder
	if err := c.BodyParser(&elder); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
	}

	createdElder, err := ec.service.CreateElder(&elder)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to create elder",
			Error:   err.Error(),
		})
	}

	responseData := res.ElderResponseDTO{Elder: *createdElder}
	return c.Status(fiber.StatusCreated).JSON(res.ResponseWrapper{
		Success: true,
		Message: "Elder created successfully",
		Data:    responseData,
	})
}

func (ec *ElderController) UpdateElder(c *fiber.Ctx) error {
	elderID := c.Params("elder_id")
	var elder models.Elder
	if err := c.BodyParser(&elder); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
	}

	updatedElder, err := ec.service.UpdateElder(elderID, &elder)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to update elder",
			Error:   err.Error(),
		})
	}

	responseData := res.ElderResponseDTO{Elder: *updatedElder}
	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Elder updated successfully",
		Data:    responseData,
	})
}

func (ec *ElderController) GetEldersByUserID(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	elders, err := ec.service.GetEldersByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to retrieve elders",
			Error:   err.Error(),
		})
	}

	responseData := res.EldersResponseDTO{Elders: elders}
	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Elders retrieved successfully",
		Data:    responseData,
	})
}

func GetElderAreas(c *fiber.Ctx) error {
	elderID := c.Params("elder_id")
	areas := []models.Area{
		{
			AreaID:          "area1",
			ElderID:         elderID,
			CaregiverID:     "dummy-caregiver-id",
			CenterLat:       -6.200000,
			CenterLong:      106.816666,
			FreeAreaRadius:  100,
			WatchAreaRadius: 50,
			IsActive:        true,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
	}
	responseData := res.AreasResponseDTO{Areas: areas}
	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Areas retrieved successfully",
		Data:    responseData,
	})
}

func GetElderLocationHistory(c *fiber.Ctx) error {
	elderID := c.Params("elder_id")
	histories := models.LocationHistory{
		LocationHistoryID: "history1",
		ElderID:           elderID,
		CaregiverID:       "dummy-caregiver-id",
		CreatedAt:         time.Now(),
		Points: []models.LocationHistoryPoint{
			{
				PointID:   "point1",
				Latitude:  -6.200000,
				Longitude: 106.816666,
				Timestamp: time.Now(),
			},
			{
				PointID:   "point2",
				Latitude:  -6.200001,
				Longitude: 106.816667,
				Timestamp: time.Now().Add(1 * time.Hour),
			},
		},
	}
	responseData := res.LocationHistoryResponseDTO{LocationHistory: histories}
	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Location history retrieved successfully",
		Data:    responseData,
	})
}

func GetElderAgendas(c *fiber.Ctx) error {
	elderID := c.Params("elder_id")
	agendas := []models.Agenda{
		{
			AgendaID:    "agenda1",
			ElderID:     elderID,
			CaregiverID: "dummy-caregiver-id",
			Category:    "Medical",
			Content1:    "Doctor appointment",
			Content2:    "Check blood pressure",
			Datetime:    time.Now().Add(24 * time.Hour),
			IsFinished:  false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	responseData := res.AgendasResponseDTO{Agendas: agendas}
	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Agendas retrieved successfully",
		Data:    responseData,
	})
}

func GetElderEmergencyAlerts(c *fiber.Ctx) error {
	elderID := c.Params("elder_id")
	alerts := models.EmergencyAlert{
		EmergencyAlertID: "alert1",
		ElderID:          elderID,
		CaregiverID:      "dummy-caregiver-id",
		Datetime:         time.Now(),
		ElderLat:         -6.200000,
		ElderLong:        106.816666,
		IsDismissed:      false,
	}
	responseData := res.EmergencyAlertResponseDTO{EmergencyAlert: alerts}
	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Emergency alerts retrieved successfully",
		Data:    responseData,
	})
}
