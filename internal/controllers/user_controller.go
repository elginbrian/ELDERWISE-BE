package controllers

import (
	"time"

	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	res "github.com/elginbrian/ELDERWISE-BE/pkg/dto/response"
	"github.com/gofiber/fiber/v2"
)

func GetUserByID(c *fiber.Ctx) error {
	userID := c.Params("user_id")

	user := models.User{
		UserID:    userID,
		Email:     "dummy@example.com",
		CreatedAt: time.Now(),
	}

	responseData := res.UserResponseDTO{
		User: user,
	}

	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "User retrieved successfully",
		Data:    responseData,
	})
}

func GetUserCaregivers(c *fiber.Ctx) error {
	userID := c.Params("user_id")

	caregivers := []models.Caregiver{
		{
			CaregiverID:  "cg1",
			UserID:       userID,
			Name:         "Dummy Caregiver One",
			Birthdate:    time.Now().AddDate(-30, 0, 0),
			Gender:       "M",
			PhoneNumber:  "081234567890",
			ProfileURL:   "https://example.com/cg1",
			Relationship: "Son",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			CaregiverID:  "cg2",
			UserID:       userID,
			Name:         "Dummy Caregiver Two",
			Birthdate:    time.Now().AddDate(-32, 0, 0),
			Gender:       "F",
			PhoneNumber:  "089876543210",
			ProfileURL:   "https://example.com/cg2",
			Relationship: "Daughter",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	responseData := res.CaregiversResponseDTO{
		Caregivers: caregivers,
	}

	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Caregivers retrieved successfully",
		Data:    responseData,
	})
}

func GetUserElders(c *fiber.Ctx) error {
	userID := c.Params("user_id")

	elders := []models.Elder{
		{
			ElderID:    "elder1",
			UserID:     userID,
			Name:       "Dummy Elder One",
			Birthdate:  time.Now().AddDate(-70, 0, 0),
			Gender:     "F",
			BodyHeight: 150.5,
			BodyWeight: 55.0,
			PhotoURL:   "https://example.com/elder1",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			ElderID:    "elder2",
			UserID:     userID,
			Name:       "Dummy Elder Two",
			Birthdate:  time.Now().AddDate(-75, 0, 0),
			Gender:     "M",
			BodyHeight: 160.0,
			BodyWeight: 60.0,
			PhotoURL:   "https://example.com/elder2",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	responseData := res.EldersResponseDTO{
		Elders: elders,
	}

	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Elders retrieved successfully",
		Data:    responseData,
	})
}
