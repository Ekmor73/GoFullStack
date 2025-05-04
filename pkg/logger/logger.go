package logger

import (
	"go-fiber/config"
	"os"

	"github.com/rs/zerolog"
)

func NewLogger(config *config.LogConfig) *zerolog.Logger {
	zerolog.SetGlobalLevel(zerolog.Level(config.Level))
	var logger zerolog.Logger
	if config.Format == "json" {
		logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	} else {
		conseleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
		logger = zerolog.New(conseleWriter).With().Timestamp().Logger()
	}
	return &logger
}
