package controllers

import (
	"fmt"
	"time"

	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"github.com/elginbrian/ELDERWISE-BE/internal/services"
	res "github.com/elginbrian/ELDERWISE-BE/pkg/dto/response"
	"github.com/gofiber/fiber/v2"
)

type EmergencyAlertController struct {
	service services.EmergencyAlertService
	smsService services.SMSService
}

func NewEmergencyAlertController(service services.EmergencyAlertService, smsService services.SMSService) *EmergencyAlertController {
	return &EmergencyAlertController{
		service: service,
		smsService: smsService,
	}
}

func (c *EmergencyAlertController) GetEmergencyAlertByID(ctx *fiber.Ctx) error {
	alertID := ctx.Params("emergency_alert_id")
	alert, err := c.service.GetEmergencyAlertByID(alertID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Emergency alert not found",
			Error:   err.Error(),
		})
	}

	return ctx.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Emergency alert retrieved successfully",
		Data:    res.EmergencyAlertResponseDTO{EmergencyAlert: *alert},
	})
}

func (c *EmergencyAlertController) CreateEmergencyAlert(ctx *fiber.Ctx) error {
	var alert models.EmergencyAlert
	if err := ctx.BodyParser(&alert); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
	}

	if err := c.service.CreateEmergencyAlert(&alert); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to create emergency alert",
			Error:   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(res.ResponseWrapper{
		Success: true,
		Message: "Emergency alert created successfully and notification sent",
		Data:    res.EmergencyAlertResponseDTO{EmergencyAlert: alert},
	})
}

func (c *EmergencyAlertController) UpdateEmergencyAlert(ctx *fiber.Ctx) error {
	alertID := ctx.Params("emergency_alert_id")
	var alert models.EmergencyAlert
	if err := ctx.BodyParser(&alert); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
	}

	if err := c.service.UpdateEmergencyAlert(alertID, &alert); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to update emergency alert",
			Error:   err.Error(),
		})
	}

	return ctx.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Emergency alert updated successfully",
		Data:    res.EmergencyAlertResponseDTO{EmergencyAlert: alert},
	})
}

func (c *EmergencyAlertController) MockEmergencyAlert(ctx *fiber.Ctx) error {
	phoneNumber := "+6285749806571"
	
	if phone := ctx.Query("phone"); phone != "" {
		phoneNumber = phone
	}
	
	mockTime := time.Now()
	mockLat := -6.200000  
	mockLong := 106.816666 
	
	message := fmt.Sprintf("⚠️TEST ALERT! Time: %s. Map: https://maps.google.com/?q=%f,%f No action needed.",
		mockTime.Format("02/01 15:04"),
		mockLat,
		mockLong,
	)
	
	fmt.Printf("Sending test SMS to %s\n", phoneNumber)
	if err := c.smsService.SendMessage(phoneNumber, message); err != nil {
		fmt.Printf("Error sending SMS: %v\n", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to send test SMS notification",
			Error:   err.Error(),
			Data: map[string]interface{}{
				"recipient": phoneNumber,
			},
		})
	}
	
	return ctx.Status(fiber.StatusOK).JSON(res.ResponseWrapper{
		Success: true,
		Message: "Test emergency alert SMS sent to " + phoneNumber,
		Data: map[string]interface{}{
			"recipient": phoneNumber,
			"sent_at": mockTime,
		},
	})
}
