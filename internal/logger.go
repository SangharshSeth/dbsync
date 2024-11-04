package internal

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

func InitLogger() *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	return logger
}
