package log

import (
	"context"
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

type (
	ctxMarker struct{}

	ctxProvider struct {
		logger logrus.FieldLogger
		fields logrus.Fields
	}
)

var (
	_ctxMarkerKey = ctxMarker{}

	_nullLogger = &logrus.Logger{
		Out:       ioutil.Discard,
		Formatter: &logrus.TextFormatter{},
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.PanicLevel,
	}
)

// AddFields adds logger fields to the context for later extraction.
func AddFields(ctx context.Context, fields logrus.Fields) {
	log, ok := ctx.Value(_ctxMarkerKey).(*ctxProvider)
	if !ok || log == nil {
		return
	}

	for k, v := range fields {
		log.fields[k] = v
	}
}

// Extract takes the call-scoped logrus.FieldLogger from the context with provided fields.
func Extract(ctx context.Context) logrus.FieldLogger {
	log, ok := ctx.Value(_ctxMarkerKey).(*ctxProvider)
	if !ok || log == nil {
		return _nullLogger
	}

	// Add log fields added until now if needed
	if len(log.fields) < 1 {
		return log.logger
	}

	fields := make(logrus.Fields, len(log.fields))
	for k, v := range log.fields {
		fields[k] = v
	}

	return log.logger.WithFields(fields)
}

// ToContext adds the logrus.FieldLogger to the context for extraction later
// returning the new context that has been created.
func ToContext(ctx context.Context, logger logrus.FieldLogger) context.Context {
	if logger == nil {
		return ctx
	}

	log := ctxProvider{
		logger: logger,
		fields: make(logrus.Fields),
	}

	return context.WithValue(ctx, _ctxMarkerKey, &log)
}
