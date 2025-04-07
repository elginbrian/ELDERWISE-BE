package bootstrap

import (
	"os"

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

	// Configs
	supabaseConfig := config.NewSupabaseConfig()
	emailConfig := config.NewEmailConfig()
	
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "your-default-secret-key" 
	}
	
	// Services
	emailService := services.NewEmailService(emailConfig)
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
	
	routeSetup := routes.NewRouteSetup(
		authController,
		userController,
		caregiverController,
		elderController, 
		areaController,
		storageController,
		emergencyAlertController,
		agendaController,
		locationHistoryController ,
	)

	app := fiber.New()
	
	routeSetup.Setup(app, jwtSecret)

	return app
}
