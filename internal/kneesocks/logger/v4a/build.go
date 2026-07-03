package v4a

import (
	"io"
	"os"
	"path"
	"socks/internal/kneesocks/config/v4a"
	"socks/pkg/utils"

	"github.com/rs/zerolog"
)

func BuildZerolog(config v4a.LoggerConfig) (zerolog.Logger, error) {
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
		if err == v4a.LoggerDisabledError {
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
		if err == v4a.LoggerDisabledError {
			return utils.BuildDefaultZerolog(126, loggers)
		}
	}

	return utils.BuildDefaultZerolog(level, loggers)
}
