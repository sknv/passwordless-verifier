package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Formatter string

const FormatterText Formatter = "text"

// Config is a logger config.
type Config struct {
	Level     string
	Formatter Formatter
}

// Option configures *logrus.Logger.
type Option func(*logrus.Logger)

// Init the global logger instance from the provided config once.
func Init(cfg Config) {
	lvl, formatter := ParseLevel(cfg.Level), GetFormatter(cfg.Formatter)
	Build(lvl, formatter)
}

// ParseLevel parses the provided string level and returns a corresponding known level or DefaultLevel.
func ParseLevel(level string) logrus.Level {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl = logrus.InfoLevel // default level
	}

	return lvl
}

func GetFormatter(formatter Formatter) logrus.Formatter {
	if formatter == FormatterText {
		return &logrus.TextFormatter{}
	}

	return &logrus.JSONFormatter{} // default formatter
}

// Build the default logger with the provided log level.
func Build(level logrus.Level, formatter logrus.Formatter, options ...Option) {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(level)
	logrus.SetFormatter(formatter)

	// Apply options
	for _, opt := range options {
		opt(logrus.StandardLogger())
	}
}
