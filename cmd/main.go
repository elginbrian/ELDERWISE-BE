package main

import (
	"log"
	"os"
	"strconv"

	"github.com/elginbrian/ELDERWISE-BE/config"
	"github.com/elginbrian/ELDERWISE-BE/internal/bootstrap"
	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	logFile, err := setupLogging()
	if err != nil {
		log.Printf("Warning: Could not set up log file: %v", err)
	} else if logFile != nil {
		defer logFile.Close()
	}
	
	log.Println("Starting Elderwise Backend Service")
	
	runTests, _ := strconv.ParseBool(os.Getenv("NETWORK_TEST_ON_STARTUP"))
	if runTests {
		log.Println("Network tests will be executed during container startup")
	}
	
	db := config.ConnectDB()

	err = db.AutoMigrate(
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

func setupLogging() (*os.File, error) {
	err := os.MkdirAll("logs", 0755)
	if err != nil {
		return nil, err
	}
	
	logFile, err := os.OpenFile("logs/elderwise.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	
	return logFile, nil
}
