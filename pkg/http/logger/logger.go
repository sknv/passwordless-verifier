package logger

import (
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"

	pkglog "github.com/sknv/passwordless-verifier/pkg/log"
)

type Logger struct {
	*logrus.Logger
}

// WithLogger replaces logger for an echo server.
func WithLogger(e *echo.Echo) {
	e.Logger = &Logger{Logger: pkglog.L()}
}

func (l *Logger) Level() log.Lvl {
	switch l.Logger.Level {
	case logrus.DebugLevel:
		return log.DEBUG
	case logrus.WarnLevel:
		return log.WARN
	case logrus.ErrorLevel:
		return log.ERROR
	case logrus.InfoLevel:
		return log.INFO
	}
	return log.OFF
}

func (l *Logger) SetLevel(lvl log.Lvl) {
	switch lvl {
	case log.DEBUG:
		l.Logger.SetLevel(logrus.DebugLevel)
	case log.WARN:
		l.Logger.SetLevel(logrus.WarnLevel)
	case log.ERROR:
		l.Logger.SetLevel(logrus.ErrorLevel)
	case log.INFO:
		l.Logger.SetLevel(logrus.InfoLevel)
	default:
		l.Errorf("invalid level %d", lvl)
	}
}

func (l *Logger) SetPrefix(string) {}

func (l *Logger) Prefix() string { return "" }

func (l *Logger) SetHeader(string) {}

func (l *Logger) Output() io.Writer {
	return l.Logger.Out
}

func (l *Logger) SetOutput(w io.Writer) {
	l.Logger.SetOutput(w)
}

func (l *Logger) Printj(j log.JSON) {
	l.Logger.WithFields(logrus.Fields(j)).Print()
}

func (l Logger) Debugj(j log.JSON) {
	l.Logger.WithFields(logrus.Fields(j)).Debug()
}

func (l Logger) Infoj(j log.JSON) {
	l.Logger.WithFields(logrus.Fields(j)).Info()
}

func (l Logger) Warnj(j log.JSON) {
	l.Logger.WithFields(logrus.Fields(j)).Warn()
}

func (l Logger) Errorj(j log.JSON) {
	l.Logger.WithFields(logrus.Fields(j)).Error()
}

func (l Logger) Fatalj(j log.JSON) {
	l.Logger.WithFields(logrus.Fields(j)).Fatal()
}

func (l Logger) Panicj(j log.JSON) {
	l.Logger.WithFields(logrus.Fields(j)).Panic()
}
