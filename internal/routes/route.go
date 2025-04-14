package routes

import (
	"github.com/elginbrian/ELDERWISE-BE/internal/controllers"
	"github.com/elginbrian/ELDERWISE-BE/internal/middleware"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

type RouteSetup struct {
	AuthController           *controllers.AuthController
	UserController           *controllers.UserController
	CaregiverController      *controllers.CaregiverController
	ElderController          *controllers.ElderController
	AreaController           *controllers.AreaController
	StorageController        *controllers.StorageController
	EmergencyAlertController *controllers.EmergencyAlertController
	AgendaController         *controllers.AgendaController
	LocationHistoryController *controllers.LocationHistoryController
	AlertViewerController    *controllers.AlertViewerController
	NotificationController   *controllers.NotificationController
}

func NewRouteSetup(
	authController *controllers.AuthController,
	userController *controllers.UserController,
	caregiverController *controllers.CaregiverController,
	elderController *controllers.ElderController,
	areaController *controllers.AreaController,
	storageController *controllers.StorageController,
	emergencyAlertController *controllers.EmergencyAlertController,
	agendaController *controllers.AgendaController,
	locationHistoryController *controllers.LocationHistoryController,
	alertViewerController *controllers.AlertViewerController,
	notificationController *controllers.NotificationController,
	) *RouteSetup {
	return &RouteSetup{
		AuthController:           authController,
		UserController:           userController,
		CaregiverController:      caregiverController,
		ElderController:          elderController,
		AreaController:           areaController,
		StorageController:        storageController,
		EmergencyAlertController: emergencyAlertController,
		AgendaController:         agendaController,
		LocationHistoryController: locationHistoryController,
		AlertViewerController:    alertViewerController,
		NotificationController:   notificationController,
	}
}

func (rs *RouteSetup) Setup(app *fiber.App, jwtSecret string) {
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	
	api := app.Group("/api/v1")

	api.Get("/", dummyHandler)

	api.Post("/auth/register", rs.AuthController.RegisterHandler)
	api.Post("/auth/login", rs.AuthController.LoginHandler)

	authConfig := middleware.Config{
		JwtSecret: jwtSecret,
	}
	authMiddleware := middleware.AuthenticationMiddleware(authConfig)
	
	protected := api.Group("", authMiddleware)
	
	protected.Get("/auth/me", rs.AuthController.GetCurrentUser)

	protected.Get("/users/:user_id", rs.UserController.GetUserByID)
	protected.Get("/users/:user_id/caregivers", rs.UserController.GetUserCaregivers)
	protected.Get("/users/:user_id/elders", rs.UserController.GetUserElders)

	protected.Get("/caregivers/:caregiver_id", rs.CaregiverController.GetCaregiverByID)
	protected.Post("/caregivers", rs.CaregiverController.CreateCaregiver)
	protected.Put("/caregivers/:caregiver_id", rs.CaregiverController.UpdateCaregiver)

	protected.Get("/elders/:elder_id", rs.ElderController.GetElderByID)
	protected.Post("/elders", rs.ElderController.CreateElder)
	protected.Put("/elders/:elder_id", rs.ElderController.UpdateElder)
	protected.Get("/elders/:elder_id/areas", rs.ElderController.GetElderAreas)
	protected.Get("/elders/:elder_id/location-history", rs.LocationHistoryController.GetElderLocationHistory)
	protected.Get("/elders/:elder_id/agendas", rs.AgendaController.GetElderAgendas)
	protected.Get("/elders/:elder_id/emergency-alerts", controllers.GetElderEmergencyAlerts)

	protected.Get("/areas/:area_id", rs.AreaController.GetAreaByID)
	protected.Post("/areas", rs.AreaController.CreateArea)
	protected.Put("/areas/:area_id", rs.AreaController.UpdateArea)
	protected.Delete("/areas/:area_id", rs.AreaController.DeleteArea)
	protected.Get("/caregivers/:caregiver_id/areas", rs.AreaController.GetAreasByCaregiver)

	protected.Get("/location-history/:location_history_id", rs.LocationHistoryController.GetLocationHistoryByID)
	protected.Get("/location-history/:location_history_id/points", rs.LocationHistoryController.GetLocationHistoryPoints)
	protected.Post("/location-history", rs.LocationHistoryController.CreateLocationHistory)
	protected.Post("/location-history/:location_history_id/points", rs.LocationHistoryController.AddLocationPoint)

	protected.Get("/agendas/:agenda_id", rs.AgendaController.GetAgendaByID)
	protected.Post("/agendas", rs.AgendaController.CreateAgenda)
	protected.Put("/agendas/:agenda_id", rs.AgendaController.UpdateAgenda)
	protected.Delete("/agendas/:agenda_id", rs.AgendaController.DeleteAgenda)

	protected.Get("/emergency-alerts/:emergency_alert_id", rs.EmergencyAlertController.GetEmergencyAlertByID)
	protected.Post("/emergency-alerts", rs.EmergencyAlertController.CreateEmergencyAlert)
	protected.Put("/emergency-alerts/:emergency_alert_id", rs.EmergencyAlertController.UpdateEmergencyAlert)
	
	protected.Post("/storage/images", rs.StorageController.ProcessEntityImage) 
	
	protected.Get("/mock/emergency-alert", rs.EmergencyAlertController.MockEmergencyAlert)
	api.Get("/alerts-viewer", rs.AlertViewerController.ViewAlerts)

	protected.Get("/elders/:elder_id/notifications", rs.NotificationController.GetNotifications)
	protected.Get("/elders/:elder_id/notifications/check", rs.NotificationController.CheckNotifications)
	protected.Get("/elders/:elder_id/notifications/unread", rs.NotificationController.GetUnreadCount)
	protected.Put("/notifications/:notification_id/read", rs.NotificationController.MarkNotificationAsRead)
}

func dummyHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Welcome to Elderwise by Masukin Andre ke Raion"})
}




