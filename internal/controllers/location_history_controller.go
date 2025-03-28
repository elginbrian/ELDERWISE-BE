package controllers

import (
	"time"

	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	res "github.com/elginbrian/ELDERWISE-BE/pkg/dto/response"
	"github.com/gofiber/fiber/v2"
)

func GetLocationHistoryByID(c *fiber.Ctx) error {
	historyID := c.Params("location_history_id")

	history := models.LocationHistory{
		LocationHistoryID: historyID,
		ElderID:           "dummy-elder-id",
		CaregiverID:       "dummy-caregiver-id",
		CreatedAt:         time.Now(),
		Points: []models.LocationHistoryPoint{
			{
				PointID:           "point1",
				LocationHistoryID: historyID,
				Latitude:          -6.200000,
				Longitude:         106.816666,
				Timestamp:         time.Now(),
			},
			{
				PointID:           "point2",
				LocationHistoryID: historyID,
				Latitude:          -6.200001,
				Longitude:         106.816667,
				Timestamp:         time.Now().Add(5 * time.Minute),
			},
		},
	}

	responseData := res.LocationHistoryResponseDTO{
		LocationHistory: history,
	}

	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Location history retrieved successfully",
		Data:    responseData,
	})
}

func GetLocationHistoryPoints(c *fiber.Ctx) error {
	historyID := c.Params("location_history_id")

	points := []models.LocationHistoryPoint{
		{
			PointID:           "point1",
			LocationHistoryID: historyID,
			Latitude:          -6.200000,
			Longitude:         106.816666,
			Timestamp:         time.Now(),
		},
		{
			PointID:           "point2",
			LocationHistoryID: historyID,
			Latitude:          -6.200001,
			Longitude:         106.816667,
			Timestamp:         time.Now().Add(5 * time.Minute),
		},
	}

	responseData := res.LocationHistoryPointsResponseDTO{
		Points: points,
	}

	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Location history points retrieved successfully",
		Data:    responseData,
	})
}
