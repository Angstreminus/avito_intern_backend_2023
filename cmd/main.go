package main

import (
	"github.com/Angstreminus/avito_intern_backend_2023/config"
	"github.com/Angstreminus/avito_intern_backend_2023/internal/server"
)

// @title Avito Test App
// @version 1.0
// @description Users Segments Managment App

// @host localhost:8080
// @BasePath /

func main() {
	config := config.NewConfig()
	dbHandler := server.NewDatabaseHandler(config)
	Server := server.NewHttpServer(config, dbHandler)
	Server.Start()
}
