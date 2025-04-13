package repository

import (
	"time"

	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NotificationRepository struct {
	DB *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{DB: db}
}

func (r *NotificationRepository) Create(notification *models.Notification) error {
	if notification.NotificationID == "" {
		notification.NotificationID = uuid.New().String()
	}
	
	if notification.Datetime.IsZero() {
		notification.Datetime = time.Now()
	}
	
	if notification.CreatedAt.IsZero() {
		notification.CreatedAt = time.Now()
	}
	
	return r.DB.Create(notification).Error
}

func (r *NotificationRepository) FindByID(notificationID string) (*models.Notification, error) {
	var notification models.Notification
	err := r.DB.Where("notification_id = ?", notificationID).First(&notification).Error
	if err != nil {
		return nil, err
	}
	return &notification, nil
}

func (r *NotificationRepository) FindByElderID(elderID string) ([]models.Notification, error) {
	var notifications []models.Notification
	err := r.DB.Where("elder_id = ?", elderID).Order("datetime DESC").Find(&notifications).Error
	return notifications, err
}

func (r *NotificationRepository) MarkAsRead(notificationID string) error {
	return r.DB.Model(&models.Notification{}).Where("notification_id = ?", notificationID).
		Update("is_read", true).Error
}

func (r *NotificationRepository) CountUnread(elderID string) (int64, error) {
	var count int64
	err := r.DB.Model(&models.Notification{}).Where("elder_id = ? AND is_read = ?", elderID, false).Count(&count).Error
	return count, err
}

