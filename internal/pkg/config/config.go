package config

import (
	"log/slog"
	"os"
)

var (
	Addr   int
	Dir    string
	Logger *slog.Logger
)

func NewLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, nil))
}
