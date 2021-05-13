package logger

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	dlog "github.com/SpaceX-io/go-lib/debug/log"
)

type defaultLogger struct {
	sync.RWMutex
	opts Options
}

func init() {
	lvl, err := GetLevel(os.Getenv("METRIX_LOG_LEVEL"))
	if err != nil {
		lvl = InfoLevel
	}
	DefaultLogger = NewHelper(NewLogger(WithLevel(lvl)))
}

func (l *defaultLogger) Init(opts ...Option) error {
	for _, o := range opts {
		o(&l.opts)
	}

	return nil
}

func (l *defaultLogger) String() string {
	return "default"
}

func (l *defaultLogger) Fields(fields map[string]interface{}) Logger {
	l.Lock()
	l.opts.Fields = copyFields(fields)
	l.Unlock()
	return l
}

func (l *defaultLogger) Log(level Level, v ...interface{}) {
	if !l.opts.Level.Enable(level) {
		return
	}
	l.RLock()
	fields := copyFields(l.opts.Fields)
	l.RUnlock()

	fields["level"] = level.String()

	if _, file, line, ok := runtime.Caller(l.opts.CallerSkipCount); ok {
		fields["file"] = fmt.Sprintf("%s:%d", logCallerFilePath(file), line)
	}

	rec := dlog.Record{
		Timestamp: time.Now(),
		Message:   fmt.Sprint(v...),
		Metadata:  make(map[string]string, len(fields)),
	}

	keys := make([]string, 0, len(fields))
	for k, v := range fields {
		keys = append(keys, k)
		rec.Metadata[k] = fmt.Sprintf("%v", v)
	}

	sort.Strings(keys)
	metadata := ""

	for _, k := range keys {
		metadata += fmt.Sprintf(" %s=%v", k, fields[k])
	}

	t := rec.Timestamp.Format("2006-01-02 15:04:05")
	_, err := l.opts.Out.Write([]byte(fmt.Sprintf("%s %s %v\n", t, metadata, rec.Message)))
	if err != nil {
		log.Printf("log [Log] write error: %s \n", err.Error())
	}
}

func (l *defaultLogger) Logf(level Level, format string, v ...interface{}) {
	if level < l.opts.Level {
		return
	}
	l.RLock()
	fields := copyFields(l.opts.Fields)
	l.Unlock()

	fields["level"] = level.String()

	if _, file, line, ok := runtime.Caller(l.opts.CallerSkipCount); ok {
		fields["file"] = fmt.Sprintf("%s:%d", logCallerFilePath(file), line)
	}

	rec := dlog.Record{
		Timestamp: time.Now(),
		Message:   fmt.Sprintf(format, v...),
		Metadata:  make(map[string]string, len(fields)),
	}

	keys := make([]string, 0, len(fields))
	for k, v := range fields {
		keys = append(keys, k)
		rec.Metadata[k] = fmt.Sprintf("%v", v)
	}

	sort.Strings(keys)
	metadata := ""

	for _, k := range keys {
		metadata += fmt.Sprintf(" %s=%v", k, fields[k])
	}

	t := rec.Timestamp.Format("2006-01-02 15:04:05")
	_, err := l.opts.Out.Write([]byte(fmt.Sprintf("%s %s %v\n", t, metadata, rec.Message)))
	if err != nil {
		log.Printf("log [Logf] write error: %s \n", err.Error())
	}
}

func copyFields(src map[string]interface{}) map[string]interface{} {
	dst := make(map[string]interface{}, len(src))
	for k, v := range src {
		dst[k] = v
	}

	return dst
}

func NewLogger(opts ...Option) Logger {
	options := Options{
		Level:           InfoLevel,
		Fields:          make(map[string]interface{}),
		Out:             os.Stderr,
		CallerSkipCount: 2,
		Context:         context.Background(),
	}

	l := &defaultLogger{opts: options}
	if err := l.Init(opts...); err != nil {
		l.Log(FatalLevel, err)
	}
	return l
}

func logCallerFilePath(logingFilePath string) string {
	idx := strings.LastIndexByte(logingFilePath, '/')
	if idx == -1 {
		return logingFilePath
	}
	idx = strings.LastIndexByte(logingFilePath[:idx], '/')
	if idx == -1 {
		return logingFilePath
	}
	return logingFilePath[idx+1:]
}

func (l *defaultLogger) Options() Options {
	l.RLock()
	opts := l.opts
	opts.Fields = copyFields(l.opts.Fields)
	l.RUnlock()
	return opts
}
