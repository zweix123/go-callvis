package log

import (
	"os"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()

func init() {
	// zerolog default level is zerolog.TraceLevel
	logger = logger.Level(zerolog.Disabled) // go-callvis default level is zerolog.Disabled, don't take any logs
}

func SetValid() {
	logger = logger.Level(zerolog.DebugLevel)
}

func Log(format string, args ...any) {
	logger.Debug().Msgf(format, args...)
}
