package routes

import (
	"github.com/This-is-Ayush-Sharma/API-Uptime-Monitor/internal/handlers"
	"github.com/This-is-Ayush-Sharma/API-Uptime-Monitor/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) {
	// initialize service
	authService := services.NewAuthService(db)

	// initialize handler
	authHandler := handlers.NewAuthHandler(authService)

	// base api group
	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}
	}
}
