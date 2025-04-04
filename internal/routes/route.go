package routes

import (
	"github.com/elginbrian/ELDERWISE-BE/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

type RouteSetup struct {
	AuthController      *controllers.AuthController
	CaregiverController *controllers.CaregiverController
	ElderController     *controllers.ElderController
	AreaController      *controllers.AreaController
	StorageController   *controllers.StorageController
}

func NewRouteSetup(
	authController *controllers.AuthController,
	caregiverController *controllers.CaregiverController,
	elderController *controllers.ElderController,
	areaController *controllers.AreaController,
	storageController *controllers.StorageController,
	) *RouteSetup {
	return &RouteSetup{
		AuthController:      authController,
		CaregiverController: caregiverController,
		ElderController:     elderController,
		AreaController:      areaController,
		StorageController:   storageController,
	}
}

func (rs *RouteSetup) Setup(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Get("/", dummyHandler)

	api.Post("/auth/register", rs.AuthController.RegisterHandler)
	api.Post("/auth/login", rs.AuthController.LoginHandler)

	api.Get("/users/:user_id", controllers.GetUserByID)
	api.Get("/users/:user_id/caregivers", controllers.GetUserCaregivers)
	api.Get("/users/:user_id/elders", controllers.GetUserElders)

	api.Get("/caregivers/:caregiver_id", rs.CaregiverController.GetCaregiverByID)
	api.Post("/caregivers", rs.CaregiverController.CreateCaregiver)
	api.Put("/caregivers/:caregiver_id", rs.CaregiverController.UpdateCaregiver)

	api.Get("/elders/:elder_id", rs.ElderController.GetElderByID)
	api.Post("/elders", rs.ElderController.CreateElder)
	api.Put("/elders/:elder_id", rs.ElderController.UpdateElder)
	api.Get("/elders/:elder_id/areas", controllers.GetElderAreas)
	api.Get("/elders/:elder_id/location-history", controllers.GetElderLocationHistory)
	api.Get("/elders/:elder_id/agendas", controllers.GetElderAgendas)
	api.Get("/elders/:elder_id/emergency-alerts", controllers.GetElderEmergencyAlerts)

	api.Get("/areas/:area_id", rs.AreaController.GetAreaByID)
	api.Post("/areas", rs.AreaController.CreateArea)
	api.Put("/areas/:area_id", rs.AreaController.UpdateArea)
	api.Delete("/areas/:area_id", rs.AreaController.DeleteArea)
	api.Get("/caregivers/:caregiver_id/areas", rs.AreaController.GetAreasByCaregiver)

	api.Get("/location-history/:location_history_id", controllers.GetLocationHistoryByID)
	api.Get("/location-history/:location_history_id/points", controllers.GetLocationHistoryPoints)

	api.Get("/agendas/:agenda_id", controllers.GetAgendaByID)
	api.Post("/agendas", controllers.CreateAgenda)
	api.Put("/agendas/:agenda_id", controllers.UpdateAgenda)
	api.Delete("/agendas/:agenda_id", controllers.DeleteAgenda)

	api.Get("/emergency-alerts/:emergency_alert_id", dummyHandler)
	api.Post("/emergency-alerts", dummyHandler)
	api.Put("/emergency-alerts/:emergency_alert_id", dummyHandler)
	
	api.Post("/storage/images", rs.StorageController.ProcessEntityImage) 
}

func dummyHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Welcome to Elderwise by Masukin Andre ke Raion"})
}
