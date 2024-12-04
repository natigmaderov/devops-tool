package main

import (
	"github.com/natigmaderov/devops-tool/cmd/api"
	"github.com/natigmaderov/devops-tool/config"
	"github.com/natigmaderov/devops-tool/db"
	"log"
)

func main() {
	cfg := db.PostgresConfig{
		Host:     config.Envs.Host,
		Port:     5432,
		User:     config.Envs.User,
		Password: config.Envs.Password,
		DBName:   config.Envs.DBName,
		SSLMode:  config.Envs.SSLMode,
	}
	postgresStorage, err := db.NewPostgresStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServe(":8080", postgresStorage.Pool)
	defer postgresStorage.Pool.Close()
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
