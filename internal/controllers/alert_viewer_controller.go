package controllers

import (
	"fmt"
	"html/template"
	"time"

	"github.com/elginbrian/ELDERWISE-BE/internal/repository"
	"github.com/gofiber/fiber/v2"
)

type AlertViewerController struct {
	alertRepo  repository.EmergencyAlertRepository
	elderRepo  repository.ElderRepository
}

func NewAlertViewerController(
	alertRepo repository.EmergencyAlertRepository,
	elderRepo repository.ElderRepository,
) *AlertViewerController {
	return &AlertViewerController{
		alertRepo: alertRepo,
		elderRepo: elderRepo,
	}
}

func (c *AlertViewerController) ViewAlerts(ctx *fiber.Ctx) error {
	cutoffTime := time.Now().Add(-24 * time.Hour)
	
	alerts, err := c.alertRepo.GetRecentAlerts(cutoffTime)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve alerts",
		})
	}
	
	var formattedAlerts []map[string]interface{}
	for _, alert := range alerts {
		elder, err := c.elderRepo.FindByID(alert.ElderID)
		
		elderName := "Unknown Elder"
		if err == nil && elder != nil {
			elderName = elder.Name
		}
		
		formattedAlerts = append(formattedAlerts, map[string]interface{}{
			"id":          alert.EmergencyAlertID,
			"elder_name":  elderName,
			"timestamp":   alert.Datetime.Format("2006-01-02 15:04:05"),
			"dismissed":   alert.IsDismissed,
			"location":    fmt.Sprintf("%.6f, %.6f", alert.ElderLat, alert.ElderLong),
			"map_url":     fmt.Sprintf("https://maps.google.com/?q=%f,%f", alert.ElderLat, alert.ElderLong),
		})
	}
	
	htmlContent := `
<!DOCTYPE html>
<html>
<head>
  <title>Elderwise Emergency Alerts</title>
  <style>
    body { font-family: Arial, sans-serif; margin: 20px; }
    h1 { color: #333; }
    .alert-card {
      border: 1px solid #ddd;
      border-radius: 8px;
      padding: 15px;
      margin-bottom: 15px;
      background-color: #f9f9f9;
    }
    .alert-card.dismissed { background-color: #e9e9e9; }
    .alert-header {
      display: flex;
      justify-content: space-between;
      margin-bottom: 10px;
    }
    .alert-elder { font-weight: bold; font-size: 18px; }
    .alert-time { color: #666; }
    .map-link {
      display: inline-block;
      margin-top: 10px;
      background-color: #0066CC;
      color: white;
      padding: 8px 15px;
      text-decoration: none;
      border-radius: 5px;
    }
    .no-alerts {
      padding: 20px;
      background-color: #f0f0f0;
      border-radius: 5px;
      text-align: center;
    }
  </style>
</head>
<body>
  <h1>Elderwise Emergency Alerts (Last 24 Hours)</h1>
  
  {{if .alerts}}
    {{range .alerts}}
      <div class="alert-card {{if .dismissed}}dismissed{{end}}">
        <div class="alert-header">
          <div class="alert-elder">{{.elder_name}}</div>
          <div class="alert-time">{{.timestamp}}</div>
        </div>
        <div>Location: {{.location}}</div>
        <div>Status: {{if .dismissed}}Dismissed{{else}}Active{{end}}</div>
        <a class="map-link" href="{{.map_url}}" target="_blank">View on Map</a>
      </div>
    {{end}}
  {{else}}
    <div class="no-alerts">
      No emergency alerts in the last 24 hours.
    </div>
  {{end}}
</body>
</html>
`
	
	tmpl, err := template.New("alerts").Parse(htmlContent)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Template error")
	}
	
	ctx.Set("Content-Type", "text/html")
	return tmpl.Execute(ctx.Response().BodyWriter(), fiber.Map{
		"alerts": formattedAlerts,
	})
}
