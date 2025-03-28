package main

import (
	"log"
	"os"

	"github.com/elginbrian/ELDERWISE-BE/config"
	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"github.com/elginbrian/ELDERWISE-BE/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	
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
	)
	if err != nil {
		log.Fatal("Auto migration gagal: ", err)
	}

	app := fiber.New()

	routes.Setup(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server berjalan pada port %s", port)
	log.Fatal(app.Listen(":" + port))
}
