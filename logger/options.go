package logger

import (
	"context"
	"io"
)

type Option func(*Options)
type Options struct {
	Level           Level
	Fields          map[string]interface{}
	Out             io.Writer
	CallerSkipCount int
	Context         context.Context
}

func WithFields(fileds map[string]interface{}) Option {
	return func(o *Options) {
		o.Fields = fileds
	}
}

func WithLevel(level Level) Option {
	return func(o *Options) {
		o.Level = level
	}
}

func WithOutPut(out io.Writer) Option {
	return func(o *Options) {
		o.Out = out
	}
}

func WithCallerSkipCount(c int) Option {
	return func(o *Options) {
		o.CallerSkipCount = c
	}
}

func SetOption(k, v interface{}) Option {
	return func(o *Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, k, v)
	}
}
