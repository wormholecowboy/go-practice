package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

type Logger struct {
	minLevel Level
	out      io.Writer
}

func New(min Level, out io.Writer) *Logger {
	if out == nil {
		out = os.Stderr
	}
	return &Logger{min, out}
}

func (l *Logger) log(level Level, args ...any) {
	if l.minLevel > level {
		return
	}
	ts := time.Now().UTC().Format(time.RFC3339)
	msg := fmt.Sprint(args...)
	fmt.Fprintf(l.out, "[%s] %s %s\n", level, ts, msg)
}

func (l *Logger) Debug(args ...any) {
	l.log(LevelDebug, args...)
}

func (l *Logger) Info(args ...any) {
	l.log(LevelInfo, args...)
}

func (l *Logger) Warn(args ...any) {
	l.log(LevelWarn, args...)
}

func (l *Logger) Error(args ...any) {
	l.log(LevelError, args...)
	os.Exit(1)
}

func (l *Logger) SetLevel(level Level) {
	l.minLevel = level
}

func main() {
  logger := Logger{LevelInfo, os.Stderr}

	logger.Debug("Should not see this")
	logger.Info("String:", "hello")
	logger.Warn("Warn int:", 42)
	logger.Error("Error struct:", struct{field string}{field:"thing"})

}
