package main

import (
	"github.com/gin-gonic/gin"
	"github.com/misbahul-alam/notepad-api/internal/config"
)

func main() {
	cfg := config.Load()
	dbURL := cfg.DB.DSN()

	println(dbURL)

	server := gin.Default()

	err := server.Run(":8080")
	if err != nil {
		return
	}
}
