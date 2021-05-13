package logger

var (
	DefaultLogger Logger
)

type Logger interface {
	Init(options ...Option) error
	Options() Options
	Fields(fields map[string]interface{}) Logger
	Log(level Level, v ...interface{})
	Logf(level Level, format string, v ...interface{})
	String() string
}

func Init(options ...Option) error {
	return DefaultLogger.Init(options...)
}

func Fields(fields map[string]interface{}) Logger {
	return DefaultLogger.Fields(fields)
}

func Log(level Level, format string, v ...interface{}) {
	DefaultLogger.Logf(level, format, v...)
}

func String() string {
	return DefaultLogger.String()
}
