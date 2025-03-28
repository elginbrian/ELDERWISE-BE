package controllers

import (
	"time"

	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	res "github.com/elginbrian/ELDERWISE-BE/pkg/dto/response"
	"github.com/gofiber/fiber/v2"
)

func GetElderByID(c *fiber.Ctx) error {
	elderID := c.Params("elder_id")
	elder := models.Elder{
		ElderID:    elderID,
		UserID:     "dummy-user-id",
		Name:       "Dummy Elder",
		Birthdate:  time.Now().AddDate(-70, 0, 0),
		Gender:     "M",
		BodyHeight: 160.0,
		BodyWeight: 60.0,
		PhotoURL:   "https://example.com/elder.jpg",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	responseData := res.ElderResponseDTO{Elder: elder}
	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Elder retrieved successfully",
		Data:    responseData,
	})
}

func CreateElder(c *fiber.Ctx) error {
	var elder models.Elder
	if err := c.BodyParser(&elder); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
	}

	elder.ElderID = "dummy-elder-id"
	elder.CreatedAt = time.Now()
	elder.UpdatedAt = time.Now()

	responseData := res.ElderResponseDTO{Elder: elder}
	return c.Status(fiber.StatusCreated).JSON(res.ResponseWrapper{
		Success: true,
		Message: "Elder created successfully",
		Data:    responseData,
	})
}

func UpdateElder(c *fiber.Ctx) error {
	elderID := c.Params("elder_id")
	var elder models.Elder
	if err := c.BodyParser(&elder); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
	}

	elder.ElderID = elderID
	elder.UpdatedAt = time.Now()

	responseData := res.ElderResponseDTO{Elder: elder}
	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Elder updated successfully",
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
