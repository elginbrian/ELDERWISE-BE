package bootstrap

import (
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

	authService := services.NewAuthService(authRepo)
	caregiverService := services.NewCaregiverService(caregiverRepo)
	elderService := services.NewElderService(elderRepo)

	authController := controllers.NewAuthController(authService)
	caregiverController := controllers.NewCaregiverController(caregiverService)
	elderController := controllers.NewElderController(elderService)
	
	routeSetup := routes.NewRouteSetup(authController, caregiverController, elderController)

	app := fiber.New()
	routeSetup.Setup(app)

	return app
}
