package docs


type HealthCheck struct {
	Status  string `json:"status" example:"ok"`
	Version string `json:"version" example:"1.0.0"`
}

type Error struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Bad Request"`
}





