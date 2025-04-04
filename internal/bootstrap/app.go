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
	
	authRepo := repository.NewAuthRepository(db)
	caregiverRepo := repository.NewCaregiverRepository(db)
	elderRepo := repository.NewElderRepository(db)
	areaRepo := repository.NewAreaRepository(db)
	storageRepo := repository.NewStorageRepository(db)

	supabaseConfig := config.NewSupabaseConfig()
	
	authService := services.NewAuthService(authRepo)
	caregiverService := services.NewCaregiverService(caregiverRepo)
	elderService := services.NewElderService(elderRepo)
	areaService := services.NewAreaService(areaRepo)
	storageService := services.NewStorageService(storageRepo, elderRepo, caregiverRepo, supabaseConfig)

	authController := controllers.NewAuthController(authService)
	caregiverController := controllers.NewCaregiverController(caregiverService)
	elderController := controllers.NewElderController(elderService)
	areaController := controllers.NewAreaController(areaService)
	storageController := controllers.NewStorageController(storageService, supabaseConfig)
	
	routeSetup := routes.NewRouteSetup(
		authController, 
		caregiverController, 
		elderController, 
		areaController,
		storageController,
	)

	app := fiber.New()
	routeSetup.Setup(app)

	return app
}
