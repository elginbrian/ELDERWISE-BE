package services

import (
	"time"

	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"github.com/elginbrian/ELDERWISE-BE/internal/repository"
	"github.com/google/uuid"
)

type AgendaService struct {
	repo repository.AgendaRepository
}

func NewAgendaService(repo repository.AgendaRepository) *AgendaService {
	return &AgendaService{
		repo: repo,
	}
}

func (s *AgendaService) CreateAgenda(agenda *models.Agenda) (*models.Agenda, error) {
	agenda.AgendaID = uuid.New().String()
	agenda.CreatedAt = time.Now()
	agenda.UpdatedAt = time.Now()

	err := s.repo.Create(agenda)
	if err != nil {
		return nil, err
	}

	return agenda, nil
}

func (s *AgendaService) GetAgendaByID(agendaID string) (*models.Agenda, error) {
	return s.repo.FindByID(agendaID)
}

func (s *AgendaService) GetAgendasByElderID(elderID string) ([]models.Agenda, error) {
	return s.repo.FindByElderID(elderID)
}

func (s *AgendaService) UpdateAgenda(agenda *models.Agenda) (*models.Agenda, error) {
	existingAgenda, err := s.repo.FindByID(agenda.AgendaID)
	if err != nil {
		return nil, err
	}

	agenda.CreatedAt = existingAgenda.CreatedAt
	agenda.UpdatedAt = time.Now()

	err = s.repo.Update(agenda)
	if err != nil {
		return nil, err
	}

	return agenda, nil
}

func (s *AgendaService) DeleteAgenda(agendaID string) error {
	return s.repo.Delete(agendaID)
}


