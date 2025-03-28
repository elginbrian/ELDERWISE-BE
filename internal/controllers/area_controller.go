package controllers

import (
	"time"

	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	res "github.com/elginbrian/ELDERWISE-BE/pkg/dto/response"
	"github.com/gofiber/fiber/v2"
)

func GetAreaByID(c *fiber.Ctx) error {
	areaID := c.Params("area_id")
	area := models.Area{
		AreaID:          areaID,
		ElderID:         "dummy-elder-id",
		CaregiverID:     "dummy-caregiver-id",
		CenterLat:       -6.200000,
		CenterLong:      106.816666,
		FreeAreaRadius:  100,
		WatchAreaRadius: 50,
		IsActive:        true,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	responseData := res.AreaResponseDTO{Area: area}
	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Area retrieved successfully",
		Data:    responseData,
	})
}

func CreateArea(c *fiber.Ctx) error {
	var area models.Area
	if err := c.BodyParser(&area); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
	}

	area.AreaID = "dummy-area-id"
	area.CreatedAt = time.Now()
	area.UpdatedAt = time.Now()

	responseData := res.AreaResponseDTO{Area: area}
	return c.Status(fiber.StatusCreated).JSON(res.ResponseWrapper{
		Success: true,
		Message: "Area created successfully",
		Data:    responseData,
	})
}

func UpdateArea(c *fiber.Ctx) error {
	areaID := c.Params("area_id")
	var area models.Area
	if err := c.BodyParser(&area); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
	}

	area.AreaID = areaID
	area.UpdatedAt = time.Now()

	responseData := res.AreaResponseDTO{Area: area}
	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Area updated successfully",
		Data:    responseData,
	})
}

func DeleteArea(c *fiber.Ctx) error {
	areaID := c.Params("area_id")
	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Area deleted successfully",
		Data:    areaID,
	})
}

func GetAreasByCaregiver(c *fiber.Ctx) error {
	caregiverID := c.Params("caregiver_id")
	areas := []models.Area{
		{
			AreaID:          "area1",
			ElderID:         "dummy-elder-id",
			CaregiverID:     caregiverID,
			CenterLat:       -6.200000,
			CenterLong:      106.816666,
			FreeAreaRadius:  100,
			WatchAreaRadius: 50,
			IsActive:        true,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
		{
			AreaID:          "area2",
			ElderID:         "dummy-elder-id",
			CaregiverID:     caregiverID,
			CenterLat:       -6.210000,
			CenterLong:      106.826666,
			FreeAreaRadius:  150,
			WatchAreaRadius: 75,
			IsActive:        true,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
	}

	responseData := res.AreasResponseDTO{Areas: areas}
	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Areas retrieved successfully for caregiver",
		Data:    responseData,
	})
}
