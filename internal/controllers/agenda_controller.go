package controllers

import (
	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"github.com/elginbrian/ELDERWISE-BE/internal/services"
	req "github.com/elginbrian/ELDERWISE-BE/pkg/dto/request"
	res "github.com/elginbrian/ELDERWISE-BE/pkg/dto/response"
	"github.com/gofiber/fiber/v2"
)

type AgendaController struct {
	AgendaService *services.AgendaService
}

func NewAgendaController(agendaService *services.AgendaService) *AgendaController {
	return &AgendaController{
		AgendaService: agendaService,
	}
}

// GetAgendaByID godoc
// @Summary Get an agenda by ID
// @Description Get detailed information about an agenda by its ID
// @Tags agendas
// @Accept json
// @Produce json
// @Param agenda_id path string true "Agenda ID"
// @Success 200 {object} res.ResponseWrapper{data=res.AgendaResponseDTO} "Agenda retrieved successfully"
// @Failure 404 {object} res.ResponseWrapper "Agenda not found"
// @Failure 500 {object} res.ResponseWrapper "Failed to retrieve agenda"
// @Router /agendas/{agenda_id} [get]
// @Security Bearer
func (c *AgendaController) GetAgendaByID(ctx *fiber.Ctx) error {
	agendaID := ctx.Params("agenda_id")
	
	agenda, err := c.AgendaService.GetAgendaByID(agendaID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Agenda not found",
			Error:   err.Error(),
		})
	}

	responseData := res.AgendaResponseDTO{Agenda: *agenda}
	return ctx.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Agenda retrieved successfully",
		Data:    responseData,
	})
}

// CreateAgenda godoc
// @Summary Create a new agenda
// @Description Create a new agenda with the provided information
// @Tags agendas
// @Accept json
// @Produce json
// @Param agenda body req.AgendaRequestDTO true "Agenda information"
// @Success 201 {object} res.ResponseWrapper{data=res.AgendaResponseDTO} "Agenda created successfully"
// @Failure 400 {object} res.ResponseWrapper "Invalid request payload"
// @Failure 500 {object} res.ResponseWrapper "Failed to create agenda"
// @Router /agendas [post]
// @Security Bearer
func (c *AgendaController) CreateAgenda(ctx *fiber.Ctx) error {
	var agendaReq req.AgendaRequestDTO
	if err := ctx.BodyParser(&agendaReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
	}

	agenda := models.Agenda{
		ElderID:     agendaReq.ElderID,
		CaregiverID: agendaReq.CaregiverID,
		Category:    agendaReq.Category,
		Content1:    agendaReq.Content1,
		Content2:    agendaReq.Content2,
		Datetime:    agendaReq.Datetime,
		IsFinished:  agendaReq.IsFinished,
	}

	createdAgenda, err := c.AgendaService.CreateAgenda(&agenda)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to create agenda",
			Error:   err.Error(),
		})
	}

	responseData := res.AgendaResponseDTO{Agenda: *createdAgenda}
	return ctx.Status(fiber.StatusCreated).JSON(res.ResponseWrapper{
		Success: true,
		Message: "Agenda created successfully",
		Data:    responseData,
	})
}

// UpdateAgenda godoc
// @Summary Update an agenda
// @Description Update an agenda with the provided information
// @Tags agendas
// @Accept json
// @Produce json
// @Param agenda_id path string true "Agenda ID"
// @Param agenda body req.AgendaRequestDTO true "Agenda information"
// @Success 200 {object} res.ResponseWrapper{data=res.AgendaResponseDTO} "Agenda updated successfully"
// @Failure 400 {object} res.ResponseWrapper "Invalid request payload"
// @Failure 404 {object} res.ResponseWrapper "Agenda not found"
// @Failure 500 {object} res.ResponseWrapper "Failed to update agenda"
// @Router /agendas/{agenda_id} [put]
// @Security Bearer
func (c *AgendaController) UpdateAgenda(ctx *fiber.Ctx) error {
	agendaID := ctx.Params("agenda_id")
	
	var agendaReq req.AgendaRequestDTO
	if err := ctx.BodyParser(&agendaReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
	}

	agenda := models.Agenda{
		AgendaID:    agendaID,
		ElderID:     agendaReq.ElderID,
		CaregiverID: agendaReq.CaregiverID,
		Category:    agendaReq.Category,
		Content1:    agendaReq.Content1,
		Content2:    agendaReq.Content2,
		Datetime:    agendaReq.Datetime,
		IsFinished:  agendaReq.IsFinished,
	}

	updatedAgenda, err := c.AgendaService.UpdateAgenda(&agenda)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to update agenda",
			Error:   err.Error(),
		})
	}

	responseData := res.AgendaResponseDTO{Agenda: *updatedAgenda}
	return ctx.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Agenda updated successfully",
		Data:    responseData,
	})
}

// DeleteAgenda godoc
// @Summary Delete an agenda
// @Description Delete an agenda by its ID
// @Tags agendas
// @Accept json
// @Produce json
// @Param agenda_id path string true "Agenda ID"
// @Success 200 {object} res.ResponseWrapper "Agenda deleted successfully"
// @Failure 404 {object} res.ResponseWrapper "Agenda not found"
// @Failure 500 {object} res.ResponseWrapper "Failed to delete agenda"
// @Router /agendas/{agenda_id} [delete]
// @Security Bearer
func (c *AgendaController) DeleteAgenda(ctx *fiber.Ctx) error {
	agendaID := ctx.Params("agenda_id")
	
	err := c.AgendaService.DeleteAgenda(agendaID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to delete agenda",
			Error:   err.Error(),
		})
	}

	return ctx.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Agenda deleted successfully",
		Data:    agendaID,
	})
}

// GetElderAgendas godoc
// @Summary Get agendas for a specific elder
// @Description Get all agendas associated with a specific elder
// @Tags elders
// @Accept json
// @Produce json
// @Param elder_id path string true "Elder ID"
// @Success 200 {object} res.ResponseWrapper{data=res.AgendasResponseDTO} "Agendas retrieved successfully"
// @Failure 404 {object} res.ResponseWrapper "Elder not found"
// @Failure 500 {object} res.ResponseWrapper "Failed to retrieve agendas"
// @Router /elders/{elder_id}/agendas [get]
// @Security Bearer
func (c *AgendaController) GetElderAgendas(ctx *fiber.Ctx) error {
	elderID := ctx.Params("elder_id")
	
	agendas, err := c.AgendaService.GetAgendasByElderID(elderID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(res.ResponseWrapper{
			Success: false,
			Message: "Failed to retrieve agendas",
			Error:   err.Error(),
		})
	}

	responseData := res.AgendasResponseDTO{Agendas: agendas}
	return ctx.JSON(res.ResponseWrapper{
		Success: true,
		Message: "Agendas retrieved successfully",
		Data:    responseData,
	})
}



