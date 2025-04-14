package services

import (
	"fmt"
	"log"
	"time"

	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"github.com/elginbrian/ELDERWISE-BE/internal/repository"
	"github.com/google/uuid"
)

type NotificationService struct {
	notificationRepo    repository.NotificationRepository
	locationHistoryRepo repository.LocationHistoryRepository
	areaRepo            repository.AreaRepository
	agendaRepo          repository.AgendaRepository
	emergencyAlertRepo  repository.EmergencyAlertRepository
	elderRepo           repository.ElderRepository
}

func NewNotificationService(
	notificationRepo repository.NotificationRepository,
	locationHistoryRepo repository.LocationHistoryRepository,
	areaRepo repository.AreaRepository,
	agendaRepo repository.AgendaRepository,
	emergencyAlertRepo repository.EmergencyAlertRepository,
	elderRepo repository.ElderRepository,
) *NotificationService {
	if notificationRepo == nil || locationHistoryRepo == nil || areaRepo == nil || 
	   agendaRepo == nil || emergencyAlertRepo == nil || elderRepo == nil {
		log.Println("WARNING: One or more repositories are nil in NotificationService")
	}
	
	return &NotificationService{
		notificationRepo:    notificationRepo,
		locationHistoryRepo: locationHistoryRepo,
		areaRepo:            areaRepo,
		agendaRepo:          agendaRepo,
		emergencyAlertRepo:  emergencyAlertRepo,
		elderRepo:           elderRepo,
	}
}

func (s *NotificationService) GetNotificationsByElderID(elderID string) ([]models.Notification, error) {
	return s.notificationRepo.FindByElderID(elderID)
}

func (s *NotificationService) MarkNotificationAsRead(notificationID string) error {
	return s.notificationRepo.MarkAsRead(notificationID)
}

func (s *NotificationService) CountUnreadNotifications(elderID string) (int64, error) {
	return s.notificationRepo.CountUnread(elderID)
}

func (s *NotificationService) CreateNotification(elderID, message string, notificationType models.NotificationType, relatedID string) error {
	notification := &models.Notification{
		NotificationID: uuid.New().String(),
		ElderID:        elderID,
		Type:           notificationType,
		Message:        message,
		Datetime:       time.Now(),
		IsRead:         false,
		RelatedID:      relatedID,
		CreatedAt:      time.Now(),
	}

	return s.notificationRepo.Create(notification)
}

func (s *NotificationService) CheckForNotifications(elderID string) ([]models.Notification, error) {
	var newNotifications []models.Notification

	if err := s.checkAreaBreaches(elderID, &newNotifications); err != nil {
		log.Printf("Error checking area breaches: %v", err)
	}

	if err := s.checkAgendas(elderID, &newNotifications); err != nil {
		log.Printf("Error checking agendas: %v", err)
	}

	if err := s.checkEmergencyAlerts(elderID, &newNotifications); err != nil {
		log.Printf("Error checking emergency alerts: %v", err)
	}

	return s.notificationRepo.FindByElderID(elderID)
}

func (s *NotificationService) checkAreaBreaches(elderID string, notifications *[]models.Notification) error {
	if s.locationHistoryRepo == nil || s.areaRepo == nil {
		return fmt.Errorf("locationHistoryRepo or areaRepo is nil")
	}

	locationHistory, err := s.locationHistoryRepo.GetElderLocationHistory(elderID)
	if err != nil {
		if err.Error() == "record not found" {
			return nil
		}
		return fmt.Errorf("failed to get elder location history: %w", err)
	}

	if len(locationHistory.Points) == 0 {
		return nil
	}

	latestPoint := locationHistory.Points[len(locationHistory.Points)-1]

	areas, err := s.areaRepo.FindByElderID(elderID)
	if err != nil {
		return fmt.Errorf("failed to get elder areas: %w", err)
	}

	for _, area := range areas {
		isInside := isPointInArea(latestPoint.Latitude, latestPoint.Longitude, area)
		if !isInside {
			notification := models.Notification{
				NotificationID: uuid.New().String(),
				ElderID:        elderID,
				Type:           models.NotificationTypeAreaBreach,
				Message:        fmt.Sprintf("Elder has left safe area: %d", area.WatchAreaRadius),
				Datetime:       time.Now(),
				IsRead:         false,
				RelatedID:      area.AreaID,
				CreatedAt:      time.Now(),
			}

			if err := s.notificationRepo.Create(&notification); err != nil {
				return fmt.Errorf("failed to create area breach notification: %w", err)
			}

			*notifications = append(*notifications, notification)
		}
	}

	return nil
}

func (s *NotificationService) checkAgendas(elderID string, notifications *[]models.Notification) error {
	if s.agendaRepo == nil {
		return fmt.Errorf("agendaRepo is nil")
	}

	agendas, err := s.agendaRepo.FindByElderID(elderID)
	if err != nil {
		if err.Error() == "record not found" {
			return nil
		}
		return fmt.Errorf("failed to get elder agendas: %w", err)
	}

	existingNotifications, err := s.notificationRepo.FindByElderID(elderID)
	if err != nil {
		return fmt.Errorf("failed to check existing notifications: %w", err)
	}
	
	now := time.Now()

	for _, agenda := range agendas {
		alreadyNotified := false
		for _, notification := range existingNotifications {
			if notification.RelatedID == agenda.AgendaID {
				alreadyNotified = true
				break
			}
		}
		
		if alreadyNotified {
			continue
		}
		
		if agenda.IsFinished && agenda.Datetime.After(now) {
			notification := models.Notification{
				NotificationID: uuid.New().String(),
				ElderID:        elderID,
				Type:           models.NotificationTypeAgendaCompleted,
				Message:        fmt.Sprintf("Agenda '%s' completed ahead of schedule", agenda.Content1),
				Datetime:       time.Now(),
				IsRead:         false,
				RelatedID:      agenda.AgendaID,
				CreatedAt:      time.Now(),
			}

			if err := s.notificationRepo.Create(&notification); err != nil {
				return fmt.Errorf("failed to create agenda completion notification: %w", err)
			}

			*notifications = append(*notifications, notification)
			continue
		}
		
		if !agenda.IsFinished && !agenda.Datetime.After(now) {
			overdueDuration := now.Sub(agenda.Datetime)
			if overdueDuration >= 30*time.Minute {
				notification := models.Notification{
					NotificationID: uuid.New().String(),
					ElderID:        elderID,
					Type:           models.NotificationTypeAgendaOverdue,
					Message:        fmt.Sprintf("Agenda '%s' is overdue", agenda.Content1),
					Datetime:       time.Now(),
					IsRead:         false,
					RelatedID:      agenda.AgendaID,
					CreatedAt:      time.Now(),
				}

				if err := s.notificationRepo.Create(&notification); err != nil {
					return fmt.Errorf("failed to create agenda overdue notification: %w", err)
				}

				*notifications = append(*notifications, notification)
			}
		}
	}

	return nil
}

func (s *NotificationService) checkEmergencyAlerts(elderID string, notifications *[]models.Notification) error {
	if s.emergencyAlertRepo == nil {
		return fmt.Errorf("emergencyAlertRepo is nil")
	}

	alerts, err := s.emergencyAlertRepo.FindByElderID(elderID)
	if err != nil {
		if err.Error() == "record not found" {
			return nil
		}
		return fmt.Errorf("failed to get elder emergency alerts: %w", err)
	}

	existingNotifications, err := s.notificationRepo.FindByElderID(elderID)
	if err != nil {
		return fmt.Errorf("failed to check existing notifications: %w", err)
	}

	oneHourAgo := time.Now().Add(-1 * time.Hour)

	for _, alert := range alerts {
		if alert.Datetime.Before(oneHourAgo) {
			continue
		}

		alreadyNotified := false
		for _, notification := range existingNotifications {
			if notification.Type == models.NotificationTypeEmergencyAlert && notification.RelatedID == alert.EmergencyAlertID {
				alreadyNotified = true
				break
			}
		}

		if !alreadyNotified {
			notification := models.Notification{
				NotificationID: uuid.New().String(),
				ElderID:        elderID,
				Type:           models.NotificationTypeEmergencyAlert,
				Message:        "Emergency alert triggered!",
				Datetime:       time.Now(),
				IsRead:         false,
				RelatedID:      alert.EmergencyAlertID,
				CreatedAt:      time.Now(),
			}

			if err := s.notificationRepo.Create(&notification); err != nil {
				return fmt.Errorf("failed to create emergency alert notification: %w", err)
			}

			*notifications = append(*notifications, notification)
		}
	}

	return nil
}

func isPointInArea(lat, lng float64, area models.Area) bool {
	centerLat := area.CenterLat
	centerLng := area.CenterLong
	radius := area.WatchAreaRadius

	if radius <= 0 {
		log.Printf("Warning: Area %s has invalid radius %v", area.AreaID, radius)
		return false
	}

	latDiff := lat - centerLat
	lngDiff := lng - centerLng
	distance := latDiff*latDiff + lngDiff*lngDiff
	
	return distance <= float64(radius)*float64(radius)
}




