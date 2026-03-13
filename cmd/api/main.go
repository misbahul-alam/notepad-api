package main

import (
	"github.com/misbahul-alam/notepad-api/internal/config"
	"github.com/misbahul-alam/notepad-api/internal/db"
	"github.com/misbahul-alam/notepad-api/internal/db/sqlc"
	"github.com/misbahul-alam/notepad-api/internal/router"
)

func main() {
	cfg := config.Load()
	dbURL := cfg.DB.DSN()
	conn := db.NewPostgres(dbURL)
	defer conn.Close()

	queries := sqlc.New(conn)

	r := router.SetupRouter(queries, cfg)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
