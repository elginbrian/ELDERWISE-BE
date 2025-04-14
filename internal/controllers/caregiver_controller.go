package controllers

import (
	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"github.com/elginbrian/ELDERWISE-BE/internal/services"
	res "github.com/elginbrian/ELDERWISE-BE/pkg/dto/response"
	"github.com/gofiber/fiber/v2"
)

type CaregiverController struct {
	service services.CaregiverService
}

func NewCaregiverController(service services.CaregiverService) *CaregiverController {
	return &CaregiverController{service}
}

// GetCaregiverByID godoc
// @Summary Get caregiver by ID
// @Description Get a caregiver's details by their ID
// @Tags caregivers
// @Accept json
// @Produce json
// @Param caregiver_id path string true "Caregiver ID"
// @Success 200 {object} res.ResponseWrapper{data=res.CaregiverResponseDTO} "Caregiver retrieved successfully"
// @Failure 404 {object} res.ResponseWrapper "Caregiver not found"
// @Router /caregivers/{caregiver_id} [get]
// @Security Bearer
func (ctr *CaregiverController) GetCaregiverByID(c *fiber.Ctx) error {
	caregiverID := c.Params("caregiver_id")
	caregiver, err := ctr.service.GetCaregiverByID(caregiverID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Caregiver not found",
			Error:   err.Error(),
		})
	}

	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Caregiver retrieved successfully",
		Data:    res.CaregiverResponseDTO{Caregiver: *caregiver},
	})
}

// CreateCaregiver godoc
// @Summary Create a new caregiver
// @Description Create a new caregiver with the provided information
// @Tags caregivers
// @Accept json
// @Produce json
// @Param caregiver body models.Caregiver true "Caregiver information"
// @Success 201 {object} res.ResponseWrapper{data=res.CaregiverResponseDTO} "Caregiver created successfully"
// @Failure 400 {object} res.ResponseWrapper "Invalid request payload"
// @Failure 500 {object} res.ResponseWrapper "Failed to create caregiver"
// @Router /caregivers [post]
// @Security Bearer
func (ctr *CaregiverController) CreateCaregiver(c *fiber.Ctx) error {
	var caregiver models.Caregiver
	if err := c.BodyParser(&caregiver); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
	}

	if err := ctr.service.CreateCaregiver(&caregiver); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to create caregiver",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(res.ResponseWrapper{
		Success: true,
		Message: "Caregiver created successfully",
		Data:    res.CaregiverResponseDTO{Caregiver: caregiver},
	})
}

// UpdateCaregiver godoc
// @Summary Update a caregiver
// @Description Update a caregiver with the provided information
// @Tags caregivers
// @Accept json
// @Produce json
// @Param caregiver_id path string true "Caregiver ID"
// @Param caregiver body models.Caregiver true "Caregiver information"
// @Success 200 {object} res.ResponseWrapper{data=res.CaregiverResponseDTO} "Caregiver updated successfully"
// @Failure 400 {object} res.ResponseWrapper "Invalid request payload"
// @Failure 404 {object} res.ResponseWrapper "Failed to update caregiver"
// @Router /caregivers/{caregiver_id} [put]
// @Security Bearer
func (ctr *CaregiverController) UpdateCaregiver(c *fiber.Ctx) error {
	caregiverID := c.Params("caregiver_id")
	var caregiver models.Caregiver
	if err := c.BodyParser(&caregiver); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
	}

	if err := ctr.service.UpdateCaregiver(caregiverID, &caregiver); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to update caregiver",
			Error:   err.Error(),
		})
	}

	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Caregiver updated successfully",
		Data:    res.CaregiverResponseDTO{Caregiver: caregiver},
	})
}




