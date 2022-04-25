package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var (
	infoLogger  = zerolog.New(os.Stdout).With().Timestamp().Logger() // информирующий логгер, пишет в os.Stdout
	errorLogger = zerolog.New(os.Stderr).With().Timestamp().Logger() // логгер ошибок, пишет в os.Stderr
)

// Инициализация логгера
func Init() {
	zerolog.MessageFieldName = "msg"
	zerolog.LevelFieldName = "lvl"
	zerolog.TimestampFieldName = "ts"
	zerolog.TimeFieldFormat = "06 Jan 06 16:04:01"
}

// Информирующее логгирование в os.Stdout
func Info(subject string) *zerolog.Event {
	return infoLogger.Info().
		Str("subject", subject)
}

// Логгирование ошибок в os.Stderr
func Error(subject string) *zerolog.Event {
	return errorLogger.Error().
		Str("subject", subject)
}

// Логгирование предупреждений в os.Stderr
func Warning(subject string) *zerolog.Event {
	return errorLogger.Warn().
		Str("subject", subject)
}

// Отладочное логгирование в os.Stderr
func Debug(subject string) *zerolog.Event {
	return errorLogger.Debug().
		Str("subject", subject)
}
