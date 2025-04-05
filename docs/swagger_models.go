package docs

// This file contains type definitions for Swagger documentation
// These definitions help Swagger understand the structure of our request/response objects

// HealthCheck represents a simple health check response
type HealthCheck struct {
	Status  string `json:"status" example:"ok"`
	Version string `json:"version" example:"1.0.0"`
}

// Error represents an error response
type Error struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Bad Request"`
}

// These are just model definitions for Swagger documentation.
// The actual implementations are in the respective packages.
// Swagger uses this file to generate the API documentation.
