package controllers

import (
	"time"

	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	res "github.com/elginbrian/ELDERWISE-BE/pkg/dto/response"
	"github.com/gofiber/fiber/v2"
)

func GetEmergencyAlertByID(c *fiber.Ctx) error {
	alertID := c.Params("emergency_alert_id")
	alert := models.EmergencyAlert{
		EmergencyAlertID: alertID,
		ElderID:          "dummy-elder-id",
		CaregiverID:      "dummy-caregiver-id",
		Datetime:         time.Now(),
		ElderLat:         -6.200000,
		ElderLong:        106.816666,
		IsDismissed:      false,
	}

	responseData := res.EmergencyAlertResponseDTO{
		EmergencyAlert: alert,
	}
	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Emergency alert retrieved successfully",
		Data:    responseData,
	})
}

func CreateEmergencyAlert(c *fiber.Ctx) error {
	var alert models.EmergencyAlert
	if err := c.BodyParser(&alert); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
	}

	alert.EmergencyAlertID = "dummy-alert-id"
	alert.Datetime = time.Now()

	responseData := res.EmergencyAlertResponseDTO{
		EmergencyAlert: alert,
	}
	return c.Status(fiber.StatusCreated).JSON(res.ResponseWrapper{
		Success: true,
		Message: "Emergency alert created successfully",
		Data:    responseData,
	})
}

func UpdateEmergencyAlert(c *fiber.Ctx) error {
	alertID := c.Params("emergency_alert_id")
	var alert models.EmergencyAlert
	if err := c.BodyParser(&alert); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
	}

	alert.EmergencyAlertID = alertID
	alert.Datetime = time.Now()

	responseData := res.EmergencyAlertResponseDTO{
		EmergencyAlert: alert,
	}
	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Emergency alert updated successfully",
		Data:    responseData,
	})
}
