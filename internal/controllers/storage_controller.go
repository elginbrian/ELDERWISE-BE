package controllers

import (
	"github.com/elginbrian/ELDERWISE-BE/config"
	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"github.com/elginbrian/ELDERWISE-BE/internal/services"
	res "github.com/elginbrian/ELDERWISE-BE/pkg/dto/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type StorageController struct {
	service        services.StorageService
	supabaseConfig *config.SupabaseConfig
}

func NewStorageController(service services.StorageService, supabaseConfig *config.SupabaseConfig) *StorageController {
	return &StorageController{
		service:        service,
		supabaseConfig: supabaseConfig,
	}
}

func (c *StorageController) ProcessEntityImage(ctx *fiber.Ctx) error {
	var upload models.StorageUpload
	if err := ctx.BodyParser(&upload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
	}
	
	if upload.URL == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Image URL is required",
		})
	}
	
	if upload.ID == "" {
		upload.ID = uuid.New().String()
	}
	
	if err := c.service.ProcessImageUpload(&upload); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to process image",
			Error:   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(res.ResponseWrapper{
		Success: true,
		Message: "Image processed successfully",
		Data:    upload,
	})
}
