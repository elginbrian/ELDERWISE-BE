package services

import (
	"fmt"
	"math"
	"time"

	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"github.com/elginbrian/ELDERWISE-BE/internal/repository"
)

type NotificationService struct {
	notificationRepo   *repository.NotificationRepository
	locationHistoryRepo *repository.LocationHistoryRepository
	areaRepo           repository.AreaRepository
	agendaRepo         *repository.AgendaRepository
	emergencyAlertRepo repository.EmergencyAlertRepository
	elderRepo          repository.ElderRepository
}

func NewNotificationService(
	notificationRepo *repository.NotificationRepository,
	locationHistoryRepo *repository.LocationHistoryRepository,
	areaRepo repository.AreaRepository,
	agendaRepo *repository.AgendaRepository,
	emergencyAlertRepo repository.EmergencyAlertRepository,
	elderRepo repository.ElderRepository,
) *NotificationService {
	return &NotificationService{
		notificationRepo:   notificationRepo,
		locationHistoryRepo: locationHistoryRepo,
		areaRepo:           areaRepo,
		agendaRepo:         agendaRepo,
		emergencyAlertRepo: emergencyAlertRepo,
		elderRepo:          elderRepo,
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

func (s *NotificationService) CheckForNotifications(elderID string) ([]models.Notification, error) {
	var notifications []models.Notification
	
	areaBreach, err := s.checkAreaBreach(elderID)
	if err != nil {
		return nil, fmt.Errorf("error checking area breach: %w", err)
	}
	if areaBreach != nil {
		notifications = append(notifications, *areaBreach)
	}
	
	overdueAgendas, err := s.checkOverdueAgendas(elderID)
	if err != nil {
		return nil, fmt.Errorf("error checking overdue agendas: %w", err)
	}
	notifications = append(notifications, overdueAgendas...)
	
	alerts, err := s.checkEmergencyAlerts(elderID)
	if err != nil {
		return nil, fmt.Errorf("error checking emergency alerts: %w", err)
	}
	notifications = append(notifications, alerts...)
	
	for i := range notifications {
		err := s.notificationRepo.Create(&notifications[i])
		if err != nil {
			return nil, fmt.Errorf("error creating notification: %w", err)
		}
	}
	
	return notifications, nil
}

func calculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const earthRadius = 6371.0
	
	lat1Rad := lat1 * math.Pi / 180
	lon1Rad := lon1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	lon2Rad := lon2 * math.Pi / 180
	
	diffLat := lat2Rad - lat1Rad
	diffLon := lon2Rad - lon1Rad
	
	a := math.Sin(diffLat/2)*math.Sin(diffLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(diffLon/2)*math.Sin(diffLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	
	return earthRadius * c * 1000
}

func (s *NotificationService) checkAreaBreach(elderID string) (*models.Notification, error) {
	history, err := s.locationHistoryRepo.GetElderLocationHistory(elderID)
	if err != nil {
		return nil, fmt.Errorf("failed to get elder location history: %w", err)
	}
	
	if len(history.Points) == 0 {
		return nil, nil
	}
	
	latestPoint := history.Points[len(history.Points)-1]
	
	areas, err := s.areaRepo.FindByElder(elderID)
	if err != nil {
		return nil, fmt.Errorf("failed to get elder areas: %w", err)
	}
	
	isOutsideAllAreas := true
	var breachedArea models.Area
	
	for _, area := range areas {
		distance := calculateDistance(
			latestPoint.Latitude, latestPoint.Longitude,
			area.CenterLat, area.CenterLong,
		)
		
		if distance <= float64(area.WatchAreaRadius) {
			isOutsideAllAreas = false
			break
		} else {
			breachedArea = area
		}
	}
	
	if isOutsideAllAreas && len(areas) > 0 {
		elder, err := s.elderRepo.FindByID(elderID)
		if err != nil {
			return nil, fmt.Errorf("failed to get elder information: %w", err)
		}
		
		return &models.Notification{
			ElderID:   elderID,
			Type:      models.NotificationTypeAreaBreach,
			Message:   fmt.Sprintf("%s has moved outside the permitted area '%d KM'", elder.Name, breachedArea.WatchAreaRadius),
			Datetime:  time.Now(),
			IsRead:    false,
			RelatedID: breachedArea.AreaID,
		}, nil
	}
	
	return nil, nil
}

func (s *NotificationService) checkOverdueAgendas(elderID string) ([]models.Notification, error) {
	var notifications []models.Notification
	
	agendas, err := s.agendaRepo.FindByElderID(elderID)
	if err != nil {
		return nil, fmt.Errorf("failed to get elder agendas: %w", err)
	}
	
	now := time.Now()
	
	for _, agenda := range agendas {
		if agenda.IsFinished {
			continue
		}
		
		if now.After(agenda.Datetime) {
			elder, err := s.elderRepo.FindByID(elderID)
			if err != nil {
				return nil, fmt.Errorf("failed to get elder information: %w", err)
			}
			
			notification := models.Notification{
				ElderID:   elderID,
				Type:      models.NotificationTypeAgendaOverdue,
				Message:   fmt.Sprintf("Reminder: %s has an overdue agenda - %s", elder.Name, agenda.Content1),
				Datetime:  time.Now(),
				IsRead:    false,
				RelatedID: agenda.AgendaID,
			}
			
			notifications = append(notifications, notification)
		}
	}
	
	return notifications, nil
}

func (s *NotificationService) checkEmergencyAlerts(elderID string) ([]models.Notification, error) {
	var notifications []models.Notification
	
	alerts, err := s.emergencyAlertRepo.FindByElderID(elderID)
	if err != nil {
		return nil, fmt.Errorf("failed to get elder emergency alerts: %w", err)
	}
	
	twentyFourHoursAgo := time.Now().Add(-24 * time.Hour)
	
	for _, alert := range alerts {
		if alert.IsDismissed || alert.Datetime.Before(twentyFourHoursAgo) {
			continue
		}
		
		elder, err := s.elderRepo.FindByID(elderID)
		if err != nil {
			return nil, fmt.Errorf("failed to get elder information: %w", err)
		}
		
		notification := models.Notification{
			ElderID:   elderID,
			Type:      models.NotificationTypeEmergencyAlert,
			Message:   fmt.Sprintf("URGENT: %s sent an emergency alert", elder.Name),
			Datetime:  time.Now(),
			IsRead:    false,
			RelatedID: alert.EmergencyAlertID,
		}
		
		notifications = append(notifications, notification)
	}
	
	return notifications, nil
}

