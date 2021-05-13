package logger

import "context"

type loggerKey struct{}

func FromContext(ctx context.Context) (Logger, bool) {
	logger, ok := ctx.Value(loggerKey{}).(Logger)
	return logger, ok
}

func NewContext(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}
