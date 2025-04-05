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
	emailService services.EmailService
	authService services.AuthService
}

func NewEmergencyAlertController(service services.EmergencyAlertService, emailService services.EmailService, authService services.AuthService) *EmergencyAlertController {
	return &EmergencyAlertController{
		service: service,
		emailService: emailService,
		authService: authService,
	}
}

// GetEmergencyAlertByID godoc
// @Summary Get emergency alert by ID
// @Description Get an emergency alert's details by its ID
// @Tags emergency-alerts
// @Accept json
// @Produce json
// @Param emergency_alert_id path string true "Emergency Alert ID"
// @Success 200 {object} res.ResponseWrapper{data=res.EmergencyAlertResponseDTO} "Emergency alert retrieved successfully"
// @Failure 404 {object} res.ResponseWrapper "Emergency alert not found"
// @Router /emergency-alerts/{emergency_alert_id} [get]
// @Security Bearer
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

// CreateEmergencyAlert godoc
// @Summary Create a new emergency alert
// @Description Create a new emergency alert and send notifications
// @Tags emergency-alerts
// @Accept json
// @Produce json
// @Param alert body models.EmergencyAlert true "Emergency alert information"
// @Success 201 {object} res.ResponseWrapper{data=res.EmergencyAlertResponseDTO} "Emergency alert created successfully"
// @Failure 400 {object} res.ResponseWrapper "Invalid request payload"
// @Failure 500 {object} res.ResponseWrapper "Failed to create emergency alert"
// @Router /emergency-alerts [post]
// @Security Bearer
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

// UpdateEmergencyAlert godoc
// @Summary Update an emergency alert
// @Description Update an emergency alert status (e.g., dismissing it)
// @Tags emergency-alerts
// @Accept json
// @Produce json
// @Param emergency_alert_id path string true "Emergency Alert ID"
// @Param alert body models.EmergencyAlert true "Emergency alert information"
// @Success 200 {object} res.ResponseWrapper{data=res.EmergencyAlertResponseDTO} "Emergency alert updated successfully"
// @Failure 400 {object} res.ResponseWrapper "Invalid request payload"
// @Failure 500 {object} res.ResponseWrapper "Failed to update emergency alert"
// @Router /emergency-alerts/{emergency_alert_id} [put]
// @Security Bearer
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

// MockEmergencyAlert godoc
// @Summary Send a test emergency alert SMS
// @Description Send a test SMS to the specified phone number
// @Tags emergency-alerts
// @Accept json
// @Produce json
// @Param phone query string false "Phone number to send the test SMS to"
// @Success 200 {object} res.ResponseWrapper "Test emergency alert SMS sent"
// @Failure 500 {object} res.ResponseWrapper "Failed to send test SMS notification"
// @Router /mock/emergency-alert [get]
func (c *EmergencyAlertController) MockEmergencyAlert(ctx *fiber.Ctx) error {
	userID := ctx.Query("user_id")
	email := "elginbrian49@gmail.com" 
	
	if userID != "" {
		
		user, err := c.authService.GetUserByID(userID)
		if err == nil && user.Email != "" {
			email = user.Email
		}
	} else if emailParam := ctx.Query("email"); emailParam != "" {
		
		email = emailParam
	}
	
	mockTime := time.Now()
	mockLat := -6.200000  
	mockLong := 106.816666 
	
	subject := "⚠️ TEST EMERGENCY ALERT"
	message := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
  <style>
    body { font-family: Arial, sans-serif; }
    .alert { background-color: #FFE0E0; padding: 15px; border-radius: 5px; }
    .alert-header { color: #D00000; font-size: 20px; font-weight: bold; }
    .map-link { margin-top: 15px; }
    .map-link a { background-color: #0066CC; color: white; padding: 10px 15px; text-decoration: none; border-radius: 5px; }
  </style>
</head>
<body>
  <div class="alert">
    <div class="alert-header">⚠️ TEST EMERGENCY ALERT</div>
    <p><strong>This is a test alert. No action required.</strong></p>
    <p>Alert time: %s</p>
    <div class="map-link">
      <a href="https://maps.google.com/?q=%f,%f" target="_blank">VIEW LOCATION ON MAP</a>
    </div>
  </div>
</body>
</html>
`, mockTime.Format("02/01 15:04"), mockLat, mockLong)
	
	if err := c.emailService.SendMessage(email, subject, message); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to send test email notification",
			Error:   err.Error(),
			Data: map[string]interface{}{
				"recipient": email,
			},
		})
	}
	
	return ctx.Status(fiber.StatusOK).JSON(res.ResponseWrapper{
		Success: true,
		Message: "Test emergency alert email sent to " + email,
		Data: map[string]interface{}{
			"recipient": email,
			"sent_at": mockTime,
		},
	})
}
