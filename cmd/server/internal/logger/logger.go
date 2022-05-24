package logger

import "go.uber.org/zap"

type Logger struct {
	log *zap.Logger
}

func New(log *zap.Logger) *Logger {
	return &Logger{log: log}
}

func (l *Logger) Info(msg string) {
	l.log.Info(msg)
}

func (l *Logger) Warn(msg string) {
	l.log.Warn(msg)
}

func (l *Logger) Fatal(msg string) {
	l.log.Fatal(msg)
}
