package app

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"log/slog"
	"os"
	"url-shortener/internal/config/cli"
	httpC "url-shortener/internal/http-controller"

	"url-shortener/internal/repository/postgres"
	uc "url-shortener/internal/usecase"
)

func Run() {
	cl := cli.MustLoad()
	log := setupLogger(cl.EnvConfig.Env)

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cl.PostgresConfig.DSN)))
	pgConnection := bun.NewDB(sqldb, pgdialect.New())

	pgRepo := postgres.New(pgConnection)

	usecase := uc.New(pgRepo)
	e := echo.New()

	httpController := httpC.New(e, usecase, log)
	httpController.Start(cl.HttpConfig)
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	const (
		envLocal = "local"
		envDev   = "dev"
		envProd  = "prod"
	)

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default: // If env config is invalid, set prod settings by default due to security
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
