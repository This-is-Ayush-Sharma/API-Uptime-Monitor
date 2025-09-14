package services

import (
	"errors"
	"github.com/This-is-Ayush-Sharma/API-Uptime-Monitor/internal/dto"
	"github.com/This-is-Ayush-Sharma/API-Uptime-Monitor/internal/utils"
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
func (s *AuthService) Register(req dto.RegisterDTO) (*models.User, error) {
	var existing models.User
	err := s.DB.Where("email = ?", req.Email).First(&existing).Error

	if err == nil {
		return nil, errors.New("Email already exists!")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:    req.Email,
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

func (s *AuthService) Login(req dto.LoginDTO) (string, error) {
	var user models.User
	err := s.DB.Where("email = ?", req.Email).First(&user).Error

	if err != nil {
		return "", errors.New("Email does not exist!")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("Password incorrect!")
	}

	token, err := utils.GenerateToken(user.UUID.String())

	if err != nil {
		return "", errors.New("Error generating token!")
	}
	return token, nil
}
