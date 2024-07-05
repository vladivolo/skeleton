package logger

import (
	"os"
	"strings"

	"log/slog"

	"github.com/vladivolo/skeleton/shared/configs"
)

type Logger struct {
	logger *slog.Logger
}

func New(conf configs.Log) *Logger {
	level := parseLogLevel(conf.Level)

	var (
		handler slog.Handler
	)

	opts := &slog.HandlerOptions{
		AddSource: conf.AddSource,
		Level:     level,
	}

	if conf.Format == "json" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	return &Logger{
		logger: slog.New(handler),
	}
}

func (log *Logger) WithField(key string, value interface{}) *Logger {
	return &Logger{
		logger: log.logger.With(key, value),
	}
}

func (log *Logger) Debug(text string, args ...any) {
	log.logger.Debug(text, args...)
}

func (log *Logger) Info(text string, args ...any) {
	log.logger.Info(text, args...)
}

func (log *Logger) Warn(text string, args ...any) {
	log.logger.Warn(text, args...)
}

func (log *Logger) Error(text string, args ...any) {
	log.logger.Error(text, args...)
}

func parseLogLevel(level string) slog.Level {
	l := slog.LevelError

	switch strings.ToLower(level) {
	case "debug":
		l = slog.LevelDebug
	case "info":
		l = slog.LevelInfo
	case "warning", "warn":
		l = slog.LevelWarn
	}

	return l
}
