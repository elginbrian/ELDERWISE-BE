package services

import (
	"fmt"
	"log"
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
	SetNotificationService(notificationService *NotificationService)
}

type emergencyAlertService struct {
	alertRepo           repository.EmergencyAlertRepository
	elderRepo           repository.ElderRepository
	caregiverRepo       repository.CaregiverRepository
	userRepo            repository.AuthRepository
	emailService        EmailService
	notificationService *NotificationService
}

func NewEmergencyAlertService(
	alertRepo repository.EmergencyAlertRepository,
	elderRepo repository.ElderRepository,
	caregiverRepo repository.CaregiverRepository,
	emailService EmailService,
) EmergencyAlertService {
	return &emergencyAlertService{
		alertRepo:      alertRepo,
		elderRepo:      elderRepo,
		caregiverRepo:  caregiverRepo,
		emailService:   emailService,
	}
}

func (s *emergencyAlertService) SetNotificationService(notificationService *NotificationService) {
	s.notificationService = notificationService
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
	
	if s.notificationService != nil {
		message := "Emergency alert triggered!"
		err := s.notificationService.CreateNotification(
			alert.ElderID, 
			message,
			models.NotificationTypeEmergencyAlert,
			alert.EmergencyAlertID,
		)
		if err != nil {
			log.Printf("Failed to create notification for emergency alert: %v", err)
		}
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
	
	user, err := s.userRepo.GetUserByID(caregiver.UserID)
	if err != nil {
		return fmt.Errorf("failed to get user info for caregiver: %w", err)
	}
	
	if user.Email == "" {
		return fmt.Errorf("user associated with caregiver has no email address")
	}
	
	subject := fmt.Sprintf("⚠️ EMERGENCY ALERT: %s needs help!", elder.Name)
	
	message := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
  <style>
	body { font-family: Arial, sans-serif; }
	.alert { background-color: #FFE0E0; padding: 15px; border-radius: 5px; }
	.alert-header { color: #D00000; font-size: 20px; font-weight: bold; }
	.map-link { margin-top: 15px; }
	.map-link a { background-color: #0066CC; color: white; padding: 10px 15px; text-decoration: none; border-radius: 5px; }
  </style>
</head>
<body>
  <div class="alert">
	<div class="alert-header">⚠️ EMERGENCY ALERT</div>
	<p><strong>%s needs immediate help!</strong></p>
	<p>Alert time: %s</p>
	<div class="map-link">
	  <a href="https://maps.google.com/?q=%f,%f" target="_blank">VIEW LOCATION ON MAP</a>
	</div>
  </div>
</body>
</html>`, elder.Name, alert.Datetime.Format("02/01 15:04"), alert.ElderLat, alert.ElderLong)

	s.emailService.SendMessageAsync(user.Email, subject, message)
	log.Printf("Emergency alert created for elder %s, email notification queued", elder.Name)
	return nil
}

