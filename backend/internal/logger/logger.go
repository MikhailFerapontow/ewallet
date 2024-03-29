package logger

import (
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func SetupLogger(level string) *slog.Logger {
	var log *slog.Logger

	switch level {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
