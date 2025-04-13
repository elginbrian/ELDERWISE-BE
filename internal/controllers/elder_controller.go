package controllers

import (
	"time"

	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"github.com/elginbrian/ELDERWISE-BE/internal/services"
	res "github.com/elginbrian/ELDERWISE-BE/pkg/dto/response"
	"github.com/gofiber/fiber/v2"
)

type ElderController struct {
	service     services.ElderService
	areaService services.AreaService
}

func NewElderController(service services.ElderService, areaService services.AreaService) *ElderController {
	return &ElderController{
		service:     service,
		areaService: areaService,
	}
}

// GetElderByID godoc
// @Summary Get elder by ID
// @Description Get an elder's details by their ID
// @Tags elders
// @Accept json
// @Produce json
// @Param elder_id path string true "Elder ID"
// @Success 200 {object} res.ResponseWrapper{data=res.ElderResponseDTO} "Elder retrieved successfully"
// @Failure 404 {object} res.ResponseWrapper "Elder not found"
// @Router /elders/{elder_id} [get]
// @Security Bearer
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

// CreateElder godoc
// @Summary Create a new elder
// @Description Create a new elder with the provided information
// @Tags elders
// @Accept json
// @Produce json
// @Param elder body models.Elder true "Elder information"
// @Success 201 {object} res.ResponseWrapper{data=res.ElderResponseDTO} "Elder created successfully"
// @Failure 400 {object} res.ResponseWrapper "Invalid request payload"
// @Failure 500 {object} res.ResponseWrapper "Failed to create elder"
// @Router /elders [post]
// @Security Bearer
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

// UpdateElder godoc
// @Summary Update an elder
// @Description Update an elder with the provided information
// @Tags elders
// @Accept json
// @Produce json
// @Param elder_id path string true "Elder ID"
// @Param elder body models.Elder true "Elder information"
// @Success 200 {object} res.ResponseWrapper{data=res.ElderResponseDTO} "Elder updated successfully"
// @Failure 400 {object} res.ResponseWrapper "Invalid request payload"
// @Failure 404 {object} res.ResponseWrapper "Elder not found"
// @Router /elders/{elder_id} [put]
// @Security Bearer
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

// GetEldersByUserID godoc
// @Summary Get elders by user ID
// @Description Get all elders associated with a user
// @Tags elders
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} res.ResponseWrapper{data=res.EldersResponseDTO} "Elders retrieved successfully"
// @Failure 500 {object} res.ResponseWrapper "Failed to retrieve elders"
// @Router /users/{user_id}/elders [get]
// @Security Bearer
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

// GetElderAreas godoc
// @Summary Get areas by elder ID
// @Description Get all areas associated with an elder
// @Tags elders
// @Accept json
// @Produce json
// @Param elder_id path string true "Elder ID"
// @Success 200 {object} res.ResponseWrapper{data=res.AreasResponseDTO} "Areas retrieved successfully"
// @Failure 404 {object} res.ResponseWrapper "Elder not found"
// @Failure 500 {object} res.ResponseWrapper "Failed to retrieve areas"
// @Router /elders/{elder_id}/areas [get]
// @Security Bearer
func (ec *ElderController) GetElderAreas(c *fiber.Ctx) error {
	elderID := c.Params("elder_id")
	
	_, err := ec.service.GetElderByID(elderID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Elder not found",
			Error:   err.Error(),
		})
	}
	
	areas, err := ec.areaService.GetAreasByElder(elderID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to retrieve areas",
			Error:   err.Error(),
		})
	}
	
	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Areas retrieved successfully",
		Data:    res.AreasResponseDTO{Areas: areas},
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


