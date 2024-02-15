package main

import (
	"log/slog"
	"os"
	"server/internal/config"
	"server/internal/logger"
	"server/pkg/handler"
	"server/pkg/repository"
	"server/pkg/service"
	"server/server"

	"github.com/spf13/viper"
)

func main() {
	config.MustLoad()

	log := logger.SetupLogger(viper.GetString("logger.level"))
	log.Info("starting http server", slog.String("level", viper.GetString("logger.level")))
	log.Debug("Debug messages are enabled")

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Error("Failed to initialize db", logger.Err(err))
		os.Exit(1)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services, log)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Error("Error occured while running http server", logger.Err(err))
		os.Exit(1)
	}
}
