package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

// DefaultLevel is the default logger level.
const DefaultLevel = logrus.InfoLevel

// Config is a logger config.
type Config struct {
	Level string
}

// Option configures *logrus.Logger.
type Option func(*logrus.Logger)

// Init the global logger instance from the provided config once.
func Init(cfg Config) {
	lvl := ParseLevel(cfg.Level)
	Build(lvl)
}

// ParseLevel parses the provided string level and returns a corresponding known level or DefaultLevel.
func ParseLevel(level string) logrus.Level {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl = DefaultLevel
	}

	return lvl
}

// Build the default logger with the provided log level.
func Build(level logrus.Level, options ...Option) {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(level)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Apply options
	for _, opt := range options {
		opt(logrus.StandardLogger())
	}
}
