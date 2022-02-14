package v5

import (
	"github.com/rs/zerolog"
	"io"
	"os"
	"socks/internal/kneesocks/config/v5"
	"socks/pkg/utils"
)

func BuildZerolog(config v5.LoggerConfig) (zerolog.Logger, error) {
	level, err := config.GetLevel()

	var loggers []io.Writer

	if err != nil {
		return utils.BuildDefaultZerolog(126, loggers)
	}

	if output, err := config.GetConsoleOutput(); err == nil {
		loggers = append(loggers, zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: output.TimeFormat,
		})
	} else {
		if err == v5.LoggerDisabledError {
			return utils.BuildDefaultZerolog(126, loggers)
		}
	}

	if output, err := config.GetFileOutput(); err == nil {
		file, err := os.OpenFile(output.Path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)

		if err != nil {
			return zerolog.Logger{}, err
		}

		loggers = append(loggers, file)
	} else {
		if err == v5.LoggerDisabledError {
			return utils.BuildDefaultZerolog(126, loggers)
		}
	}

	return utils.BuildDefaultZerolog(level, loggers)
}
