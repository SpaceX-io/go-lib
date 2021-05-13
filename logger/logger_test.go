package logger

import "testing"

func TestLogger(t *testing.T) {
	logger := NewLogger(WithLevel(TraceLevel))
	h1 := NewHelper(logger).WithFields(map[string]interface{}{"key1": "val1"})
	h1.Trace("trace msg")
	h1.Warn("warn msg")

	logger.Fields(map[string]interface{}{"key2": "val2"}).Log(InfoLevel, "test msg")
}
