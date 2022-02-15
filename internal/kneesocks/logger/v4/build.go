package v4

import (
	"github.com/rs/zerolog"
	"io"
	"os"
	"path"
	"socks/internal/kneesocks/config/v4"
	"socks/pkg/utils"
)

func BuildZerolog(config v4.LoggerConfig) (zerolog.Logger, error) {
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
		if err == v4.LoggerDisabledError {
			return utils.BuildDefaultZerolog(126, loggers)
		}
	}

	if output, err := config.GetFileOutput(); err == nil {
		dir := path.Dir(output.Path)

		dirErr := os.MkdirAll(dir, 0700)

		if dirErr != nil {
			return zerolog.Logger{}, dirErr
		}

		file, err := os.OpenFile(output.Path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)

		if err != nil {
			return zerolog.Logger{}, err
		}

		loggers = append(loggers, file)
	} else {
		if err == v4.LoggerDisabledError {
			return utils.BuildDefaultZerolog(126, loggers)
		}
	}

	return utils.BuildDefaultZerolog(level, loggers)
}
