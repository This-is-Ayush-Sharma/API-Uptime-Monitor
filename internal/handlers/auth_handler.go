package handlers

import (
	"github.com/This-is-Ayush-Sharma/API-Uptime-Monitor/internal/dto"
	"github.com/This-is-Ayush-Sharma/API-Uptime-Monitor/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthHandler manages auth endpoints
type AuthHandler struct {
	AuthService *services.AuthService
}

// new AuthHandler
func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

// Register godoc
// @Summary Register a new user
// @Description Create a new account with email and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body map[string]string true "User credentials"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.AuthService.Register(req)

	if err != nil {
		// Check if the error is specifically due to an existing email
		if err.Error() == "Email already exists!" {
			c.JSON(http.StatusConflict, gin.H{"error": "This email is already registered. Please use a different one."})
			return
		}

		// Handle all other errors as an internal server error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user."})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user created",
		"user":    user,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.AuthService.Login(req)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Header("Authorization", "Bearer "+token)
	c.JSON(http.StatusOK, gin.H{
		"message": "user logged in",
	})
}
