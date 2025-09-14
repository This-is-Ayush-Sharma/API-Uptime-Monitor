package handlers

import (
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
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.AuthService.Register(req.Email, req.Password)

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
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.AuthService.FindByEmail(req.Email)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found."})
		return
	}

	if !h.AuthService.CheckPassword(user, req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials."})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user logged in",
		"user":    user,
		"token":   "Token_HERE",
	})
}
