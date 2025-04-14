package middleware

import (
	"strings"

	res "github.com/elginbrian/ELDERWISE-BE/pkg/dto/response"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Config struct {
	JwtSecret string
}

func AuthenticationMiddleware(config Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(res.ResponseWrapper{
				Success: false,
				Message: "Unauthorized",
				Error:   "Missing or invalid authentication token",
			})
		}
		
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "invalid token signing method")
			}
			
			return []byte(config.JwtSecret), nil
		})
		
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(res.ResponseWrapper{
				Success: false,
				Message: "Unauthorized",
				Error:   "Invalid token: " + err.Error(),
			})
		}
		
		if !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(res.ResponseWrapper{
				Success: false,
				Message: "Unauthorized",
				Error:   "Invalid token",
			})
		}
		
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Locals("user", claims)
		}
		
		return c.Next()
	}
}




