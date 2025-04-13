package repository

import (
	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"gorm.io/gorm"
)

type AgendaRepository interface {
	Create(agenda *models.Agenda) error
	FindByID(agendaID string) (*models.Agenda, error)
	FindByElderID(elderID string) ([]models.Agenda, error)
	Update(agenda *models.Agenda) error
	Delete(agendaID string) error
}

type agendaRepositoryImpl struct {
	DB *gorm.DB
}

func NewAgendaRepository(db *gorm.DB) AgendaRepository {
	return &agendaRepositoryImpl{DB: db}
}

func (r *agendaRepositoryImpl) Create(agenda *models.Agenda) error {
	return r.DB.Create(agenda).Error
}

func (r *agendaRepositoryImpl) FindByID(agendaID string) (*models.Agenda, error) {
	var agenda models.Agenda
	err := r.DB.Where("agenda_id = ?", agendaID).First(&agenda).Error
	if err != nil {
		return nil, err
	}
	return &agenda, nil
}

func (r *agendaRepositoryImpl) FindByElderID(elderID string) ([]models.Agenda, error) {
	var agendas []models.Agenda
	err := r.DB.Where("elder_id = ?", elderID).Find(&agendas).Error
	if err != nil {
		return nil, err
	}
	return agendas, nil
}

func (r *agendaRepositoryImpl) Update(agenda *models.Agenda) error {
	return r.DB.Save(agenda).Error
}

func (r *agendaRepositoryImpl) Delete(agendaID string) error {
	return r.DB.Where("agenda_id = ?", agendaID).Delete(&models.Agenda{}).Error
}

