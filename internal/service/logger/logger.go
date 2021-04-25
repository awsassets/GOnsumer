package logger

import "github.com/rs/zerolog"

type (
	LoggerService struct {
		Logger zerolog.Logger
	}
)

func (l *LoggerService) Error(err error) {
	l.Logger.Err(err)
}

func (l *LoggerService) Info(msg string) {
	l.Logger.Info().Msg(msg)
}
