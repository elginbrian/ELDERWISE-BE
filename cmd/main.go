package main

import (
	"log"
	"os"

	"github.com/elginbrian/ELDERWISE-BE/config"
	"github.com/elginbrian/ELDERWISE-BE/docs"
	"github.com/elginbrian/ELDERWISE-BE/internal/bootstrap"
	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/swaggo/fiber-swagger" // swagger middleware
)

// @title ELDERWISE API
// @version 1.0
// @description ELDERWISE backend service API documentation
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@elderwise.app

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /api/v1
// @schemes http https

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	// Initialize Swagger docs
	docs.SwaggerInfo()
	
	db := config.ConnectDB()

	err := db.AutoMigrate(
		&models.User{},
		&models.Caregiver{},
		&models.Elder{},
		&models.Area{},
		&models.Agenda{},
		&models.EmergencyAlert{},
		&models.LocationHistory{},
		&models.LocationHistoryPoint{},
		&models.StorageFile{}, 
	)
	if err != nil {
		log.Fatal("Auto migration gagal: ", err)
	}

	app := bootstrap.AppBootstrap(db)
	
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",                              
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",     
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: false, 
	}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server berjalan pada port %s", port)
	log.Fatal(app.Listen(":" + port))
}
