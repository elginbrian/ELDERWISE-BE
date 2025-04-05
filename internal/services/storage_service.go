package services

import (
	"fmt"
	"time"

	"github.com/elginbrian/ELDERWISE-BE/config"
	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"github.com/elginbrian/ELDERWISE-BE/internal/repository"
	"github.com/google/uuid"
)

type StorageService interface {
	GetFileURL(bucketName, path string) string
	SaveFile(file *models.StorageFile) error
	ProcessImageUpload(upload *models.StorageUpload) error
}

type storageService struct {
	repo           repository.StorageRepository
	elderRepo      repository.ElderRepository
	caregiverRepo  repository.CaregiverRepository
	supabaseConfig *config.SupabaseConfig
}

func NewStorageService(
	repo repository.StorageRepository,
	elderRepo repository.ElderRepository,
	caregiverRepo repository.CaregiverRepository,
	supabaseConfig *config.SupabaseConfig,
) StorageService {
	return &storageService{
		repo:           repo,
		elderRepo:      elderRepo,
		caregiverRepo:  caregiverRepo,
		supabaseConfig: supabaseConfig,
	}
}

func (s *storageService) GetFileURL(bucketName, path string) string {
	return fmt.Sprintf("%s/storage/v1/object/public/%s/%s", s.supabaseConfig.URL, bucketName, path)
}

func (s *storageService) SaveFile(file *models.StorageFile) error {
	if file.FileID == "" {
		file.FileID = uuid.New().String()
	}
	
	if file.CreatedAt.IsZero() {
		file.CreatedAt = time.Now()
	}
	if file.UpdatedAt.IsZero() {
		file.UpdatedAt = time.Now()
	}
	if file.UploadedAt.IsZero() {
		file.UploadedAt = time.Now()
	}
	
	return s.repo.SaveFile(file)
}

func (s *storageService) ProcessImageUpload(upload *models.StorageUpload) error {
	now := time.Now()
	uploadedAt := now
	if upload.CreatedAt != nil {
		uploadedAt = *upload.CreatedAt
	}
	
	file := &models.StorageFile{
		FileID:     upload.ID,
		Name:       getFileNameFromPath(upload.Path),
		BucketName: getBucketFromURL(upload.URL),
		Path:       upload.Path,
		URL:        upload.URL,
		UploadedAt: uploadedAt,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
	
	if err := s.SaveFile(file); err != nil {
		return fmt.Errorf("failed to save file record: %w", err)
	}
	
	if upload.EntityType == nil || upload.EntityID == nil {
		return nil
	}

	entityType := models.EntityType(*upload.EntityType)
	entityID := *upload.EntityID
	
	switch entityType {
	case models.EntityTypeElder:
		return s.updateElderImage(entityID, upload.URL)
	case models.EntityTypeCaregiver:
		return s.updateCaregiverImage(entityID, upload.URL)
	default:
		return nil 
	}
}

func (s *storageService) updateElderImage(elderID, imageURL string) error {
	elder, err := s.elderRepo.FindByID(elderID)
	if err != nil {
		return fmt.Errorf("elder not found: %w", err)
	}
	
	elder.PhotoURL = imageURL
	elder.UpdatedAt = time.Now()
	
	return s.elderRepo.Update(elder)
}

func (s *storageService) updateCaregiverImage(caregiverID, imageURL string) error {
	caregiver, err := s.caregiverRepo.FindByID(caregiverID)
	if err != nil {
		return fmt.Errorf("caregiver not found: %w", err)
	}
	
	caregiver.ProfileURL = imageURL
	caregiver.UpdatedAt = time.Now()
	
	return s.caregiverRepo.Update(caregiver)
}

func getFileNameFromPath(path string) string {
	if path == "" {
		return ""
	}
	parts := []rune(path)
	for i := len(parts) - 1; i >= 0; i-- {
		if parts[i] == '/' {
			return string(parts[i+1:])
		}
	}
	return path
}

func getBucketFromURL(url string) string {
	return "elderwise-images" 
}
