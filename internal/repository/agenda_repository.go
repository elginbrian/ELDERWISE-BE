package repository

import (
	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"gorm.io/gorm"
)

type AgendaRepository struct {
	DB *gorm.DB
}

func NewAgendaRepository(db *gorm.DB) *AgendaRepository {
	return &AgendaRepository{DB: db}
}

func (r *AgendaRepository) Create(agenda *models.Agenda) error {
	return r.DB.Create(agenda).Error
}

func (r *AgendaRepository) FindByID(agendaID string) (*models.Agenda, error) {
	var agenda models.Agenda
	err := r.DB.Where("agenda_id = ?", agendaID).First(&agenda).Error
	if err != nil {
		return nil, err
	}
	return &agenda, nil
}

func (r *AgendaRepository) FindByElderID(elderID string) ([]models.Agenda, error) {
	var agendas []models.Agenda
	err := r.DB.Where("elder_id = ?", elderID).Find(&agendas).Error
	if err != nil {
		return nil, err
	}
	return agendas, nil
}

func (r *AgendaRepository) Update(agenda *models.Agenda) error {
	return r.DB.Save(agenda).Error
}

func (r *AgendaRepository) Delete(agendaID string) error {
	return r.DB.Where("agenda_id = ?", agendaID).Delete(&models.Agenda{}).Error
}