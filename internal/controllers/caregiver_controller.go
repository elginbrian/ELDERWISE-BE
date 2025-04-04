package controllers

import (
	"time"

	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	res "github.com/elginbrian/ELDERWISE-BE/pkg/dto/response"
	"github.com/gofiber/fiber/v2"
)

func GetCaregiverByID(c *fiber.Ctx) error {
	caregiverID := c.Params("caregiver_id")

	caregiver := models.Caregiver{
		CaregiverID:  caregiverID,
		UserID:       "dummy-user-id",
		Name:         "Dummy Caregiver",
		Birthdate:    time.Now().AddDate(-30, 0, 0),
		Gender:       "M",
		PhoneNumber:  "081234567890",
		ProfileURL:   "https://example.com/caregiver",
		Relationship: "Child",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	responseData := res.CaregiverResponseDTO{
		Caregiver: caregiver,
	}

	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Caregiver retrieved successfully",
		Data:    responseData,
	})
}

func CreateCaregiver(c *fiber.Ctx) error {
	var caregiver models.Caregiver
	if err := c.BodyParser(&caregiver); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
	}

	caregiver.CaregiverID = "dummy-caregiver-id"
	caregiver.CreatedAt = time.Now()
	caregiver.UpdatedAt = time.Now()

	responseData := res.CaregiverResponseDTO{
		Caregiver: caregiver,
	}

	return c.Status(fiber.StatusCreated).JSON(res.ResponseWrapper{
		Success: true,
		Message: "Caregiver created successfully",
		Data:    responseData,
	})
}

func UpdateCaregiver(c *fiber.Ctx) error {
	caregiverID := c.Params("caregiver_id")
	var caregiver models.Caregiver
	if err := c.BodyParser(&caregiver); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
	}

	caregiver.CaregiverID = caregiverID
	caregiver.UpdatedAt = time.Now()

	responseData := res.CaregiverResponseDTO{
		Caregiver: caregiver,
	}

	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Caregiver updated successfully",
		Data:    responseData,
	})
}
