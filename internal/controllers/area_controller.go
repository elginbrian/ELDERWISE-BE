package controllers

import (
	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"github.com/elginbrian/ELDERWISE-BE/internal/services"
	res "github.com/elginbrian/ELDERWISE-BE/pkg/dto/response"
	"github.com/gofiber/fiber/v2"
)

type AreaController struct {
	service services.AreaService
}

func NewAreaController(service services.AreaService) *AreaController {
	return &AreaController{service}
}

// GetAreaByID godoc
// @Summary Get area by ID
// @Description Get an area's details by its ID
// @Tags areas
// @Accept json
// @Produce json
// @Param area_id path string true "Area ID"
// @Success 200 {object} res.ResponseWrapper{data=res.AreaResponseDTO} "Area retrieved successfully"
// @Failure 404 {object} res.ResponseWrapper "Area not found"
// @Router /areas/{area_id} [get]
// @Security Bearer
func (ac *AreaController) GetAreaByID(c *fiber.Ctx) error {
	areaID := c.Params("area_id")
	area, err := ac.service.GetAreaByID(areaID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Area not found",
			Error:   err.Error(),
		})
	}

	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Area retrieved successfully",
		Data:    res.AreaResponseDTO{Area: *area},
	})
}

// CreateArea godoc
// @Summary Create a new area
// @Description Create a new area with the provided information
// @Tags areas
// @Accept json
// @Produce json
// @Param area body models.Area true "Area information"
// @Success 201 {object} res.ResponseWrapper{data=res.AreaResponseDTO} "Area created successfully"
// @Failure 400 {object} res.ResponseWrapper "Invalid request payload"
// @Failure 500 {object} res.ResponseWrapper "Failed to create area"
// @Router /areas [post]
// @Security Bearer
func (ac *AreaController) CreateArea(c *fiber.Ctx) error {
	var area models.Area
	if err := c.BodyParser(&area); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
	}

	area.AreaID = ""

	if err := ac.service.CreateArea(&area); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to create area",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(res.ResponseWrapper{
		Success: true,
		Message: "Area created successfully",
		Data:    res.AreaResponseDTO{Area: area},
	})
}

// UpdateArea godoc
// @Summary Update an area
// @Description Update an area with the provided information
// @Tags areas
// @Accept json
// @Produce json
// @Param area_id path string true "Area ID"
// @Param area body models.Area true "Area information"
// @Success 200 {object} res.ResponseWrapper{data=res.AreaResponseDTO} "Area updated successfully"
// @Failure 400 {object} res.ResponseWrapper "Invalid request payload"
// @Failure 500 {object} res.ResponseWrapper "Failed to update area"
// @Router /areas/{area_id} [put]
// @Security Bearer
func (ac *AreaController) UpdateArea(c *fiber.Ctx) error {
	areaID := c.Params("area_id")
	var area models.Area
	if err := c.BodyParser(&area); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
	}

	if err := ac.service.UpdateArea(areaID, &area); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to update area",
			Error:   err.Error(),
		})
	}

	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Area updated successfully",
		Data:    res.AreaResponseDTO{Area: area},
	})
}

// DeleteArea godoc
// @Summary Delete an area
// @Description Delete an area by its ID
// @Tags areas
// @Accept json
// @Produce json
// @Param area_id path string true "Area ID"
// @Success 200 {object} res.ResponseWrapper{data=string} "Area deleted successfully"
// @Failure 500 {object} res.ResponseWrapper "Failed to delete area"
// @Router /areas/{area_id} [delete]
// @Security Bearer
func (ac *AreaController) DeleteArea(c *fiber.Ctx) error {
	areaID := c.Params("area_id")
	if err := ac.service.DeleteArea(areaID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to delete area",
			Error:   err.Error(),
		})
	}

	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Area deleted successfully",
		Data:    areaID,
	})
}

// GetAreasByCaregiver godoc
// @Summary Get areas by caregiver ID
// @Description Get all areas associated with a caregiver
// @Tags areas
// @Accept json
// @Produce json
// @Param caregiver_id path string true "Caregiver ID"
// @Success 200 {object} res.ResponseWrapper{data=res.AreasResponseDTO} "Areas retrieved successfully"
// @Failure 500 {object} res.ResponseWrapper "Failed to retrieve areas"
// @Router /caregivers/{caregiver_id}/areas [get]
// @Security Bearer
func (ac *AreaController) GetAreasByCaregiver(c *fiber.Ctx) error {
	caregiverID := c.Params("caregiver_id")
	areas, err := ac.service.GetAreasByCaregiver(caregiverID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to retrieve areas",
			Error:   err.Error(),
		})
	}

	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Areas retrieved successfully for caregiver",
		Data:    res.AreasResponseDTO{Areas: areas},
	})
}



