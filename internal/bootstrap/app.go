package bootstrap

import (
	"log"
	"os"
	"time"

	"github.com/elginbrian/ELDERWISE-BE/config"
	"github.com/elginbrian/ELDERWISE-BE/internal/controllers"
	"github.com/elginbrian/ELDERWISE-BE/internal/repository"
	"github.com/elginbrian/ELDERWISE-BE/internal/routes"
	"github.com/elginbrian/ELDERWISE-BE/internal/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AppBootstrap(db *gorm.DB) *fiber.App {
	// Repositories
	authRepo := repository.NewAuthRepository(db)
	userRepo := repository.NewUserRepository(db)
	caregiverRepo := repository.NewCaregiverRepository(db)
	elderRepo := repository.NewElderRepository(db)
	areaRepo := repository.NewAreaRepository(db)
	storageRepo := repository.NewStorageRepository(db)
	emergencyAlertRepo := repository.NewEmergencyAlertRepository(db)
	locationHistoryRepo := repository.NewLocationHistoryRepository(db)
	agendaRepo := repository.NewAgendaRepository(db)

	supabaseConfig := config.NewSupabaseConfig()
	emailConfig := config.NewEmailConfig()
	
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "your-default-secret-key" 
	}
	
	if emailConfig.Provider == "mock" {
		log.Println("ERROR: Mock email provider not allowed")
		if emailConfig.Username != "" && emailConfig.Password != "" {
			emailConfig.Provider = "smtp"
			log.Println("Forced SMTP provider based on available credentials")
		} else if emailConfig.SendGridAPIKey != "" {
			emailConfig.Provider = "sendgrid"
			log.Println("Forced SendGrid provider based on available credentials")
		} else if emailConfig.MailgunAPIKey != "" && emailConfig.MailgunDomain != "" {
			emailConfig.Provider = "mailgun"
			log.Println("Forced Mailgun provider based on available credentials")
		} else {
			log.Println("ERROR: No valid email provider configuration found!")
		}
	}
	
	emailConfig.HealthCheckTimeout = 5 * time.Second

	var emailService services.EmailService
	realEmailService, err := services.NewEmailService(emailConfig)
	if err != nil {
		log.Printf("WARNING: Email service initialization failed: %v", err)
		log.Println("Using logging-only email service instead.")
		emailService = services.NewLoggingEmailService()
	} else if !realEmailService.HealthCheck() {
		log.Printf("WARNING: Email service health check failed. Email alerts may not be delivered!")
		log.Println("Using logging-only email service instead.")
		emailService = services.NewLoggingEmailService()
	} else {
		log.Println("Email service initialized successfully and health check passed")
		emailService = realEmailService
	}
	
	authService := services.NewAuthService(authRepo)
	authService.SetJWTSecret(jwtSecret)
	userService := services.NewUserService(userRepo)
	caregiverService := services.NewCaregiverService(caregiverRepo)
	elderService := services.NewElderService(elderRepo)
	areaService := services.NewAreaService(areaRepo)
	storageService := services.NewStorageService(storageRepo, elderRepo, caregiverRepo, supabaseConfig)
	emergencyAlertService := services.NewEmergencyAlertService(
		emergencyAlertRepo, 
		elderRepo, 
		caregiverRepo, 
		emailService,
	)
	locationHistoryService := services.NewLocationHistoryService(locationHistoryRepo)
	agendaService := services.NewAgendaService(agendaRepo)

	// Controllers
	authController := controllers.NewAuthController(authService)
	userController := controllers.NewUserController(userService)
	caregiverController := controllers.NewCaregiverController(caregiverService)
	elderController := controllers.NewElderController(elderService)
	areaController := controllers.NewAreaController(areaService)
	storageController := controllers.NewStorageController(storageService, supabaseConfig)
	emergencyAlertController := controllers.NewEmergencyAlertController(
		emergencyAlertService, 
		emailService,
		authService, 
	)
	locationHistoryController := controllers.NewLocationHistoryController(locationHistoryService)
	agendaController := controllers.NewAgendaController(agendaService)
	
	// Initialize the alert viewer controller
	alertViewerController := controllers.NewAlertViewerController(
		emergencyAlertRepo,
		elderRepo,
	)
	
	routeSetup := routes.NewRouteSetup(
		authController,
		userController,
		caregiverController,
		elderController, 
		areaController,
		storageController,
		emergencyAlertController,
		agendaController,
		locationHistoryController,
		alertViewerController,
	)

	app := fiber.New()
	
	routeSetup.Setup(app, jwtSecret)

	return app
}
