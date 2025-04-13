package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/elginbrian/ELDERWISE-BE/config"
	"github.com/elginbrian/ELDERWISE-BE/docs"
	"github.com/elginbrian/ELDERWISE-BE/internal/bootstrap"
	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/swaggo/fiber-swagger"
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
  
	docs.SwaggerInfo()

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
		&models.Notification{},
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

	checkServices()

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

func checkServices() {
	emailProvider := os.Getenv("EMAIL_PROVIDER")
	
	switch emailProvider {
	case "smtp":
		host := os.Getenv("EMAIL_HOST")
		if host == "" {
			host = "smtp.gmail.com"
		}
		port := os.Getenv("EMAIL_PORT")
		if port == "" {
			port = "465"
		}
		
		addr := fmt.Sprintf("%s:%s", host, port)
		log.Printf("Testing connection to SMTP server %s (this may timeout in restricted networks)...", addr)
		
		conn, err := net.DialTimeout("tcp", addr, 3*time.Second)
		if err != nil {
			log.Printf("WARNING: Cannot connect to SMTP server %s: %v", addr, err)
			log.Println("This is likely due to network restrictions or firewall rules.")
			log.Println("The application will continue, but email alerts won't be delivered.")
		
		} else {
			conn.Close()
			log.Printf("Successfully connected to SMTP server %s", addr)
		}
	
	case "sendgrid":
		if os.Getenv("SENDGRID_API_KEY") == "" {
			log.Fatalf("FATAL: SendGrid API key not provided")
		}
		log.Println("SendGrid provider configured (API connectivity will be tested on first use)")
		
	case "mailgun":
		if os.Getenv("MAILGUN_API_KEY") == "" || os.Getenv("MAILGUN_DOMAIN") == "" {
			log.Fatalf("FATAL: Mailgun API key or domain not provided")
		}
		log.Println("Mailgun provider configured (API connectivity will be tested on first use)")
		
	default:
		log.Printf("WARNING: Unknown email provider: %s", emailProvider)
	}
}


