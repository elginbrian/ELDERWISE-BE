package routes

import (
	"github.com/elginbrian/ELDERWISE-BE/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Get("/", dummyHandler)

	api.Post("/auth/register", (&controllers.AuthController{}).RegisterHandler)
	api.Post("/auth/login", (&controllers.AuthController{}).LoginHandler)

	api.Get("/users/:user_id", controllers.GetUserByID)
	api.Get("/users/:user_id/caregivers", controllers.GetUserCaregivers)
	api.Get("/users/:user_id/elders", controllers.GetUserElders)

	api.Get("/caregivers/:caregiver_id", controllers.GetCaregiverByID)
	api.Post("/caregivers", controllers.CreateCaregiver)
	api.Put("/caregivers/:caregiver_id", controllers.UpdateCaregiver)

	api.Get("/elders/:elder_id", controllers.GetElderByID)
	api.Post("/elders", controllers.CreateElder)
	api.Put("/elders/:elder_id", controllers.UpdateElder)
	api.Get("/elders/:elder_id/areas", controllers.GetElderAreas)
	api.Get("/elders/:elder_id/location-history", controllers.GetElderLocationHistory)
	api.Get("/elders/:elder_id/agendas", controllers.GetElderAgendas)
	api.Get("/elders/:elder_id/emergency-alerts", controllers.GetElderEmergencyAlerts)

	api.Get("/areas/:area_id", controllers.GetAreaByID)
	api.Post("/areas", controllers.CreateArea)
	api.Put("/areas/:area_id", controllers.UpdateArea)
	api.Delete("/areas/:area_id", controllers.DeleteArea)
	api.Get("/caregivers/:caregiver_id/areas", controllers.GetAreasByCaregiver)

	api.Get("/location-history/:location_history_id", controllers.GetLocationHistoryByID)
	api.Get("/location-history/:location_history_id/points", controllers.GetLocationHistoryPoints)

	api.Get("/agendas/:agenda_id", controllers.GetAgendaByID)
	api.Post("/agendas", controllers.CreateAgenda)
	api.Put("/agendas/:agenda_id", controllers.UpdateAgenda)
	api.Delete("/agendas/:agenda_id", controllers.DeleteAgenda)

	api.Get("/emergency-alerts/:emergency_alert_id", dummyHandler)
	api.Post("/emergency-alerts", dummyHandler)
	api.Put("/emergency-alerts/:emergency_alert_id", dummyHandler)
}

func dummyHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "This is a dummy response"})
}
