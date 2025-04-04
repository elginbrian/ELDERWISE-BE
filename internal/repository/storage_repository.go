package repository

import (
	"time"

	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StorageRepository interface {
	SaveFile(file *models.StorageFile) error
	GetFileByPath(path string) (*models.StorageFile, error)
	GetFileByID(fileID string) (*models.StorageFile, error)
}

type storageRepository struct {
	db *gorm.DB
}

func NewStorageRepository(db *gorm.DB) StorageRepository {
	return &storageRepository{db: db}
}

func (r *storageRepository) SaveFile(file *models.StorageFile) error {
	if file.FileID == "" {
		file.FileID = uuid.New().String()
	}
	
	if file.CreatedAt.IsZero() {
		file.CreatedAt = time.Now()
	}
	
	file.UpdatedAt = time.Now()
	
	return r.db.Save(file).Error
}

func (r *storageRepository) GetFileByPath(path string) (*models.StorageFile, error) {
	var file models.StorageFile
	if err := r.db.Where("path = ?", path).First(&file).Error; err != nil {
		return nil, err
	}
	return &file, nil
}

func (r *storageRepository) GetFileByID(fileID string) (*models.StorageFile, error) {
	var file models.StorageFile
	if err := r.db.Where("file_id = ?", fileID).First(&file).Error; err != nil {
		return nil, err
	}
	return &file, nil
}
