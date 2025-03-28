package controllers

import (
	"time"

	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	res "github.com/elginbrian/ELDERWISE-BE/pkg/dto/response"
	"github.com/gofiber/fiber/v2"
)

func GetAgendaByID(c *fiber.Ctx) error {
	agendaID := c.Params("agenda_id")
	agenda := models.Agenda{
		AgendaID:    agendaID,
		ElderID:     "dummy-elder-id",
		CaregiverID: "dummy-caregiver-id",
		Category:    "Medical",
		Content1:    "Doctor appointment",
		Content2:    "Follow-up consultation",
		Datetime:    time.Now().Add(24 * time.Hour),
		IsFinished:  false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	responseData := res.AgendaResponseDTO{Agenda: agenda}
	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Agenda retrieved successfully",
		Data:    responseData,
	})
}

func CreateAgenda(c *fiber.Ctx) error {
	var agenda models.Agenda
	if err := c.BodyParser(&agenda); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
	}

	agenda.AgendaID = "dummy-agenda-id"
	agenda.CreatedAt = time.Now()
	agenda.UpdatedAt = time.Now()

	responseData := res.AgendaResponseDTO{Agenda: agenda}
	return c.Status(fiber.StatusCreated).JSON(res.ResponseWrapper{
		Success: true,
		Message: "Agenda created successfully",
		Data:    responseData,
	})
}

func UpdateAgenda(c *fiber.Ctx) error {
	agendaID := c.Params("agenda_id")
	var agenda models.Agenda
	if err := c.BodyParser(&agenda); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
	}

	agenda.AgendaID = agendaID
	agenda.UpdatedAt = time.Now()

	responseData := res.AgendaResponseDTO{Agenda: agenda}
	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Agenda updated successfully",
		Data:    responseData,
	})
}

func DeleteAgenda(c *fiber.Ctx) error {
	agendaID := c.Params("agenda_id")
	return c.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Agenda deleted successfully",
		Data:    agendaID,
	})
}
