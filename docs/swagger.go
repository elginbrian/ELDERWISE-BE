package docs

import (
	"github.com/swaggo/swag"
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
func SwaggerInfo() {
	swag.Register(swag.Name, &swag.Spec{
		InfoInstanceName: "swagger",
		SwaggerTemplate:  DocTemplate,
	})
}

var DocTemplate string



