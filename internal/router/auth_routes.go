package router

import (
	"github.com/gin-gonic/gin"
	"github.com/misbahul-alam/notepad-api/internal/config"
	"github.com/misbahul-alam/notepad-api/internal/db/sqlc"
	"github.com/misbahul-alam/notepad-api/internal/handler"
	"github.com/misbahul-alam/notepad-api/internal/service"
)

func setupAuthRoutes(r *gin.Engine, q *sqlc.Queries, cfg *config.Config) {
	authService := service.NewAuthService(q, &cfg.JWT)
	authHandler := handler.NewAuthHandler(authService)

	auth := r.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/register", authHandler.Register)
	}

}
