package router

import (
	"github.com/gin-gonic/gin"
	"github.com/misbahul-alam/notepad-api/internal/config"
	"github.com/misbahul-alam/notepad-api/internal/db/sqlc"
)

func SetupRouter(q *sqlc.Queries, cfg *config.Config) *gin.Engine {
	r := gin.Default()

	setupAuthRoutes(r, q, cfg)

	return r
}
