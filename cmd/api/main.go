package main

import (
	"github.com/gin-gonic/gin"
	"github.com/misbahul-alam/notepad-api/internal/config"
	"github.com/misbahul-alam/notepad-api/internal/db"
)

func main() {
	cfg := config.Load()
	dbURL := cfg.DB.DSN()
	conn := db.NewPostgres(dbURL)
	defer conn.Close()

	server := gin.Default()

	err := server.Run(":8080")
	if err != nil {
		return
	}
}
