package utils

import (
	"github.com/rs/zerolog"
	"io"
)

func BuildDefaultZerolog(level int, loggers []io.Writer) (zerolog.Logger, error) {
	return zerolog.New(zerolog.MultiLevelWriter(loggers...)).
		With().
		Timestamp().
		Logger().
		Level(zerolog.Level(level)), nil
}
