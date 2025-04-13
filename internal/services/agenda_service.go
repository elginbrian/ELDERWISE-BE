package services

import (
	"time"

	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"github.com/elginbrian/ELDERWISE-BE/internal/repository"
	"github.com/google/uuid"
)

type AgendaService struct {
	AgendaRepository *repository.AgendaRepository
}

func NewAgendaService(agendaRepository *repository.AgendaRepository) *AgendaService {
	return &AgendaService{
		AgendaRepository: agendaRepository,
	}
}

func (s *AgendaService) CreateAgenda(agenda *models.Agenda) (*models.Agenda, error) {
	agenda.AgendaID = uuid.New().String()
	agenda.CreatedAt = time.Now()
	agenda.UpdatedAt = time.Now()

	err := s.AgendaRepository.Create(agenda)
	if err != nil {
		return nil, err
	}

	return agenda, nil
}

func (s *AgendaService) GetAgendaByID(agendaID string) (*models.Agenda, error) {
	return s.AgendaRepository.FindByID(agendaID)
}

func (s *AgendaService) GetAgendasByElderID(elderID string) ([]models.Agenda, error) {
	return s.AgendaRepository.FindByElderID(elderID)
}

func (s *AgendaService) UpdateAgenda(agenda *models.Agenda) (*models.Agenda, error) {
	existingAgenda, err := s.AgendaRepository.FindByID(agenda.AgendaID)
	if err != nil {
		return nil, err
	}

	agenda.CreatedAt = existingAgenda.CreatedAt
	agenda.UpdatedAt = time.Now()

	err = s.AgendaRepository.Update(agenda)
	if err != nil {
		return nil, err
	}

	return agenda, nil
}

func (s *AgendaService) DeleteAgenda(agendaID string) error {
	return s.AgendaRepository.Delete(agendaID)
}
