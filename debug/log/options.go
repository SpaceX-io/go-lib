package log

import "time"

type Option func(*Options)

type Options struct {
	Name   string
	Size   int
	Format FormatFunc
}

func Name(name string) Option {
	return func(o *Options) {
		o.Name = name
	}
}

func Size(size int) Option {
	return func(o *Options) {
		o.Size = size
	}
}

func Format(format FormatFunc) Option {
	return func(o *Options) {
		o.Format = format
	}
}

func DefaultOptions() Options {
	return Options{
		Size: DefaultSize,
	}
}

type ReadOptions struct {
	Since  time.Time
	Count  int
	Stream bool
}

type ReadOption func(*ReadOptions)

func Since(since time.Time) ReadOption {
	return func(ro *ReadOptions) {
		ro.Since = since
	}
}

func Count(count int) ReadOption {
	return func(ro *ReadOptions) {
		ro.Count = count
	}
}

func Stream(stream bool) ReadOption {
	return func(ro *ReadOptions) {
		ro.Stream = stream
	}
}
