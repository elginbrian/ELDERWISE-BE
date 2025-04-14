package services

import (
	"errors"
	"strings"
	"time"

	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"github.com/elginbrian/ELDERWISE-BE/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(user *models.User) (*models.User, error)
	Login(email, password string) (string, error)
	GetUserByID(userID string) (*models.User, error)
	GetUserFromToken(tokenString string) (*models.User, error)
	SetJWTSecret(secret string)
}

type authService struct {
	repo      repository.AuthRepository
	jwtSecret []byte
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) SetJWTSecret(secret string) {
	s.jwtSecret = []byte(secret)
}

func (s *authService) Register(user *models.User) (*models.User, error) {
	if existingUser, _ := s.repo.GetUserByEmail(user.Email); existingUser != nil {
		return nil, errors.New("email already exists")
	}

	user.UserID = uuid.New().String()
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()
	
	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *authService) Login(email, password string) (string, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.UserID,
		"email":   user.Email,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *authService) GetUserByID(userID string) (*models.User, error) {
	return s.repo.GetUserByID(userID)
}

func (s *authService) GetUserFromToken(tokenString string) (*models.User, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		return s.jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return nil, errors.New("user ID not found in token")
	}

	return s.GetUserByID(userID)
}




