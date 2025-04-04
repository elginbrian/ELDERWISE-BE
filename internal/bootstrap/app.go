package bootstrap

import (
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
	caregiverRepo := repository.NewCaregiverRepository(db)
	elderRepo := repository.NewElderRepository(db)
	areaRepo := repository.NewAreaRepository(db)
	storageRepo := repository.NewStorageRepository(db)
	emergencyAlertRepo := repository.NewEmergencyAlertRepository(db)

	// Configs
	supabaseConfig := config.NewSupabaseConfig()
	smsConfig := config.NewSMSConfig()
	
	// Services
	smsService := services.NewSMSService(smsConfig)
	authService := services.NewAuthService(authRepo)
	caregiverService := services.NewCaregiverService(caregiverRepo)
	elderService := services.NewElderService(elderRepo)
	areaService := services.NewAreaService(areaRepo)
	storageService := services.NewStorageService(storageRepo, elderRepo, caregiverRepo, supabaseConfig)
	emergencyAlertService := services.NewEmergencyAlertService(
		emergencyAlertRepo, 
		elderRepo, 
		caregiverRepo, 
		smsService,
	)

	// Controllers
	authController := controllers.NewAuthController(authService)
	caregiverController := controllers.NewCaregiverController(caregiverService)
	elderController := controllers.NewElderController(elderService)
	areaController := controllers.NewAreaController(areaService)
	storageController := controllers.NewStorageController(storageService, supabaseConfig)
	emergencyAlertController := controllers.NewEmergencyAlertController(
		emergencyAlertService, 
		smsService, 
	)
	
	routeSetup := routes.NewRouteSetup(
		authController, 
		caregiverController, 
		elderController, 
		areaController,
		storageController,
		emergencyAlertController,
	)

	app := fiber.New()
	routeSetup.Setup(app)

	return app
}
