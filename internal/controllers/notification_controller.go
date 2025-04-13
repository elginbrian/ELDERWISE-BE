package controllers

import (
	"github.com/elginbrian/ELDERWISE-BE/internal/services"
	res "github.com/elginbrian/ELDERWISE-BE/pkg/dto/response"
	"github.com/gofiber/fiber/v2"
)

type NotificationController struct {
	service *services.NotificationService
}

func NewNotificationController(service *services.NotificationService) *NotificationController {
	return &NotificationController{service: service}
}

// GetNotifications godoc
// @Summary Get all notifications for an elder
// @Description Get all notifications for a specific elder
// @Tags notifications
// @Accept json
// @Produce json
// @Param elder_id path string true "Elder ID"
// @Success 200 {object} res.ResponseWrapper{data=res.NotificationsResponseDTO} "Notifications retrieved successfully"
// @Failure 500 {object} res.ResponseWrapper "Failed to retrieve notifications"
// @Router /elders/{elder_id}/notifications [get]
// @Security Bearer
func (c *NotificationController) GetNotifications(ctx *fiber.Ctx) error {
	elderID := ctx.Params("elder_id")
	
	notifications, err := c.service.GetNotificationsByElderID(elderID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to retrieve notifications",
			Error:   err.Error(),
		})
	}
	
	responseData := res.NotificationsResponseDTO{Notifications: notifications}
	return ctx.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Notifications retrieved successfully",
		Data:    responseData,
	})
}

// CheckNotifications godoc
// @Summary Check for new notifications
// @Description Check for new notifications for a specific elder
// @Tags notifications
// @Accept json
// @Produce json
// @Param elder_id path string true "Elder ID"
// @Success 200 {object} res.ResponseWrapper{data=res.NotificationsResponseDTO} "Notifications checked successfully"
// @Failure 500 {object} res.ResponseWrapper "Failed to check notifications"
// @Router /elders/{elder_id}/notifications/check [get]
// @Security Bearer
func (c *NotificationController) CheckNotifications(ctx *fiber.Ctx) error {
	elderID := ctx.Params("elder_id")
	
	notifications, err := c.service.CheckForNotifications(elderID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to check for notifications",
			Error:   err.Error(),
		})
	}
	
	responseData := res.NotificationsResponseDTO{Notifications: notifications}
	return ctx.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Notifications checked successfully",
		Data:    responseData,
	})
}

// MarkNotificationAsRead godoc
// @Summary Mark a notification as read
// @Description Mark a specific notification as read
// @Tags notifications
// @Accept json
// @Produce json
// @Param notification_id path string true "Notification ID"
// @Success 200 {object} res.ResponseWrapper "Notification marked as read"
// @Failure 500 {object} res.ResponseWrapper "Failed to mark notification as read"
// @Router /notifications/{notification_id}/read [put]
// @Security Bearer
func (c *NotificationController) MarkNotificationAsRead(ctx *fiber.Ctx) error {
	notificationID := ctx.Params("notification_id")
	
	err := c.service.MarkNotificationAsRead(notificationID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to mark notification as read",
			Error:   err.Error(),
		})
	}
	
	return ctx.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Notification marked as read",
	})
}

// GetUnreadCount godoc
// @Summary Get count of unread notifications
// @Description Get count of unread notifications for a specific elder
// @Tags notifications
// @Accept json
// @Produce json
// @Param elder_id path string true "Elder ID"
// @Success 200 {object} res.ResponseWrapper{data=res.UnreadCountResponseDTO} "Unread count retrieved successfully"
// @Failure 500 {object} res.ResponseWrapper "Failed to retrieve unread count"
// @Router /elders/{elder_id}/notifications/unread [get]
// @Security Bearer
func (c *NotificationController) GetUnreadCount(ctx *fiber.Ctx) error {
	elderID := ctx.Params("elder_id")
	
	count, err := c.service.CountUnreadNotifications(elderID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to retrieve unread count",
			Error:   err.Error(),
		})
	}
	
	responseData := res.UnreadCountResponseDTO{Count: count}
	return ctx.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Unread count retrieved successfully",
		Data:    responseData,
	})
}

