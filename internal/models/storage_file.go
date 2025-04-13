package models

import "time"

type StorageFile struct {
	FileID     string    `json:"file_id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	BucketName string    `json:"bucket_name"`
	Path       string    `json:"path"`
	Size       int64     `json:"size"`
	MimeType   string    `json:"mime_type"`
	URL        string    `json:"url"`
	UploadedAt time.Time `json:"uploaded_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	MetaData   string    `json:"meta_data"`
}

type StorageWebhookPayload struct {
	Type      string `json:"type"`
	Table     string `json:"table"`
	Schema    string `json:"schema"`
	Record    Record `json:"record"`
	OldRecord Record `json:"old_record,omitempty"`
}

type Record struct {
	ID         string                 `json:"id"`
	Name       string                 `json:"name"`
	BucketID   string                 `json:"bucket_id"`
	Owner      string                 `json:"owner"`
	Path       string                 `json:"path"`
	CreatedAt  string                 `json:"created_at"`
	UpdatedAt  string                 `json:"updated_at"`
	Metadata   map[string]interface{} `json:"metadata"`
}


