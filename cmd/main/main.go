package main

import (
	"fmt"
	cfg "github.com/AlexeyChudov/todoApp/internal/config"
	sl "github.com/AlexeyChudov/todoApp/internal/lib/logger/slog"
	repo "github.com/AlexeyChudov/todoApp/internal/repository"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	var config cfg.Config = cfg.MustLoad()
	var logger *slog.Logger = setupLogger(config.Env)

	logger.Info("Starting todolist web app", slog.String("env", config.Env))
	logger.Debug("Debug level enabled")

	//TODO: init db
	// postgres connection string format: postgresql://[user[:password]@][netloc][:port][/dbname][?param1=value1&...]
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.User, config.Password,
		config.Host, config.Port, config.Database)
	storage, err := repo.NewStorage(connStr)
	if err != nil {
		logger.Error("Failed to create storage", sl.Err(err))
		os.Exit(1)
	}
	_ = storage
	logger.Info("Connected to storage", slog.String("DATABASE", config.Database))

	id, err := storage.InsertTask("have breakfast", "I/m strarving", "05.12.2025", "95.12.2025",
		"high")
	if err != nil {
		logger.Error("Failed to insert task", sl.Err(err))
	}
	fmt.Printf("Inserted ID: %s\n", id)
	//TODO: init router

	//TODO: init server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	}
	return log

}
