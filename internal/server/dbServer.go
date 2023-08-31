package server

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Angstreminus/avito_intern_backend_2023/config"
)

func InitDatabase(cfg *config.Config) *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", cfg.PGhost, cfg.PGport, cfg.PGuser, cfg.PGname, cfg.PGuser)

	dbHandler, err := sql.Open(cfg.PGdriver, connStr)
	if err != nil {
		log.Fatalf("Error while initiazile db %v", err)
	}
	return dbHandler
}
