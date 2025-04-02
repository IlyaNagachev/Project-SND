package logs

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

type Config struct {
	Level string `yaml:"level"`
}

func NewLogger(config *Config) (err error) {
	var level zerolog.Level
	switch config.Level {
	case "debug": level = zerolog.DebugLevel
	case "info": level = zerolog.InfoLevel
	case "warning": level = zerolog.WarnLevel
	case "error": level = zerolog.ErrorLevel
	}
	log.Logger = log.With().Caller().Logger().Level(level).Output(zerolog.ConsoleWriter{Out: os.Stderr})

	return
}