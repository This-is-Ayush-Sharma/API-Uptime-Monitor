package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/This-is-Ayush-Sharma/API-Uptime-Monitor/internal/models"
)

// Auth Service handles user auth operations
type AuthService struct {
	DB *gorm.DB
}

// NewAuthService create a new AuthService instance
func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{DB: db}
}

// Register create a new user with hashed password
func (s *AuthService) Register(email, password string) (*models.User, error) {
	var existing models.User
	err := s.DB.Where("email = ?", email).First(&existing).Error

	if err == nil {
		return nil, errors.New("Email already exists!")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	// save the user in DB
	if err := s.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// checkPassword compare the given password with stored hash
func (s *AuthService) CheckPassword(user *models.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

// FindByEmail fetch a user by email
func (s *AuthService) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := s.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
