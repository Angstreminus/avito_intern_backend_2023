package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PGhost     string
	PGdriver   string
	PGuser     string
	PGpassword string
	PGname     string
	PGport     string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading config file")
	}
	return &Config{
		PGhost:     os.Getenv("POSTGRES_HOST"),
		PGdriver:   os.Getenv("POSTGRES_DRIVER"),
		PGuser:     os.Getenv("POSTGRES_USER"),
		PGpassword: os.Getenv("POSTGRES_PASSWORD"),
		PGname:     os.Getenv("POSTGRES_NAME"),
		PGport:     os.Getenv("POSTGRES_PORT"),
	}
}
