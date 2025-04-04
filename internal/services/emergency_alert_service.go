package services

import (
	"fmt"
	"time"

	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"github.com/elginbrian/ELDERWISE-BE/internal/repository"
	"github.com/google/uuid"
)

type EmergencyAlertService interface {
	CreateEmergencyAlert(alert *models.EmergencyAlert) error
	GetEmergencyAlertByID(alertID string) (*models.EmergencyAlert, error)
	UpdateEmergencyAlert(alertID string, alert *models.EmergencyAlert) error
	GetEmergencyAlertsByElderID(elderID string) ([]models.EmergencyAlert, error)
}

type emergencyAlertService struct {
	alertRepo      repository.EmergencyAlertRepository
	elderRepo      repository.ElderRepository
	caregiverRepo  repository.CaregiverRepository
	smsService     SMSService
}

func NewEmergencyAlertService(
	alertRepo repository.EmergencyAlertRepository,
	elderRepo repository.ElderRepository,
	caregiverRepo repository.CaregiverRepository,
	smsService SMSService,
) EmergencyAlertService {
	return &emergencyAlertService{
		alertRepo:      alertRepo,
		elderRepo:      elderRepo,
		caregiverRepo:  caregiverRepo,
		smsService:     smsService,
	}
}

func (s *emergencyAlertService) CreateEmergencyAlert(alert *models.EmergencyAlert) error {
	if alert.EmergencyAlertID == "" {
		alert.EmergencyAlertID = uuid.New().String()
	}
	
	if alert.Datetime.IsZero() {
		alert.Datetime = time.Now()
	}
	
	if err := s.alertRepo.Create(alert); err != nil {
		return err
	}
	
	return s.sendAlertNotification(alert)
}

func (s *emergencyAlertService) GetEmergencyAlertByID(alertID string) (*models.EmergencyAlert, error) {
	return s.alertRepo.FindByID(alertID)
}

func (s *emergencyAlertService) UpdateEmergencyAlert(alertID string, alert *models.EmergencyAlert) error {
	existingAlert, err := s.alertRepo.FindByID(alertID)
	if err != nil {
		return err
	}
	
	existingAlert.IsDismissed = alert.IsDismissed
	
	return s.alertRepo.Update(existingAlert)
}

func (s *emergencyAlertService) GetEmergencyAlertsByElderID(elderID string) ([]models.EmergencyAlert, error) {
	return s.alertRepo.FindByElderID(elderID)
}

func (s *emergencyAlertService) sendAlertNotification(alert *models.EmergencyAlert) error {
	elder, err := s.elderRepo.FindByID(alert.ElderID)
	if err != nil {
		return fmt.Errorf("failed to get elder info: %w", err)
	}
	
	caregiver, err := s.caregiverRepo.FindByID(alert.CaregiverID)
	if err != nil {
		return fmt.Errorf("failed to get caregiver info: %w", err)
	}
	
	message := fmt.Sprintf("⚠️ EMERGENCY ALERT ⚠️\n\nElder: %s has triggered an emergency alert at %s.\n\nLocation: https://maps.google.com/?q=%f,%f\n\nPlease contact them immediately or seek help if needed.",
		elder.Name,
		alert.Datetime.Format("Mon, 02 Jan 2006 15:04:05"),
		alert.ElderLat,
		alert.ElderLong,
	)
	
	if err := s.smsService.SendMessage(caregiver.PhoneNumber, message); err != nil {
		return fmt.Errorf("failed to send SMS notification: %w", err)
	}
	
	return nil
}