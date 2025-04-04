package dto

type AreaRequestDTO struct {
	ElderID         string  `json:"elder_id" validate:"required"`
	CaregiverID     string  `json:"caregiver_id" validate:"required"`
	CenterLat       float64 `json:"center_lat" validate:"required"`
	CenterLong      float64 `json:"center_long" validate:"required"`
	FreeAreaRadius  int     `json:"free_area_radius" validate:"required"`
	WatchAreaRadius int     `json:"watch_area_radius" validate:"required"`
	IsActive        bool    `json:"is_active"`
}