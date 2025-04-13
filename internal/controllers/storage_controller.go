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

// ProcessEntityImage godoc
// @Summary Process an entity image
// @Description Process an uploaded image and associate it with an entity
// @Tags storage
// @Accept json
// @Produce json
// @Param upload body models.StorageUpload true "Upload information"
// @Success 201 {object} res.ResponseWrapper "Image processed successfully"
// @Failure 400 {object} res.ResponseWrapper "Invalid request payload"
// @Failure 500 {object} res.ResponseWrapper "Failed to process image"
// @Router /storage/images [post]
// @Security Bearer
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
	
	if upload.Path == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Image path is required",
		})
	}
	
	if !upload.EntityType.IsValid() {
		return ctx.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Invalid entity type",
		})
	}
	
	if upload.EntityID == nil || *upload.EntityID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Entity ID is required when entity type is provided",
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


