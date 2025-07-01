package logging

import (
	"path/filepath"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

var (
	logger = log.With().Caller().Logger()
)

func init() {
	zerolog.TimestampFieldName = "timestamp"
	zerolog.DefaultContextLogger = &logger
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}
}

// return a zerolog logger
func GetLogger() zerolog.Logger {
	return logger
}
