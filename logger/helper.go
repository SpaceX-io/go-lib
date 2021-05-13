package logger

type Helper struct {
	Logger
	fields map[string]interface{}
}

func NewHelper(logger Logger) *Helper {
	return &Helper{Logger: logger}
}

func (h *Helper) Info(args ...interface{}) {
	if !h.Logger.Options().Level.Enable(InfoLevel) {
		return
	}
	h.Logger.Fields(h.fields).Log(InfoLevel, args...)
}

func (h *Helper) Infof(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enable(InfoLevel) {
		return
	}
	h.Logger.Fields(h.fields).Logf(InfoLevel, template, args...)
}

func (h *Helper) Trace(args ...interface{}) {
	if !h.Logger.Options().Level.Enable(TraceLevel) {
		return
	}
	h.Logger.Fields(h.fields).Log(TraceLevel, args...)
}

func (h *Helper) Tracef(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enable(TraceLevel) {
		return
	}
	h.Logger.Fields(h.fields).Logf(TraceLevel, template, args...)
}

func (h *Helper) Debug(args ...interface{}) {
	if !h.Logger.Options().Level.Enable(DebugLevel) {
		return
	}
	h.Logger.Fields(h.fields).Log(DebugLevel, args...)
}

func (h *Helper) Debugf(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enable(DebugLevel) {
		return
	}
	h.Logger.Fields(h.fields).Logf(DebugLevel, template, args...)
}

func (h *Helper) Warn(args ...interface{}) {
	if !h.Logger.Options().Level.Enable(WarnLevel) {
		return
	}
	h.Logger.Fields(h.fields).Log(WarnLevel, args...)
}

func (h *Helper) Warnf(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enable(WarnLevel) {
		return
	}
	h.Logger.Fields(h.fields).Logf(WarnLevel, template, args...)
}

func (h *Helper) Error(args ...interface{}) {
	if !h.Logger.Options().Level.Enable(ErrorLevel) {
		return
	}
	h.Logger.Fields(h.fields).Log(ErrorLevel, args...)
}

func (h *Helper) Errorf(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enable(WarnLevel) {
		return
	}
	h.Logger.Fields(h.fields).Logf(WarnLevel, template, args...)
}

func (h *Helper) Fatal(args ...interface{}) {
	if !h.Logger.Options().Level.Enable(FatalLevel) {
		return
	}
	h.Logger.Fields(h.fields).Log(FatalLevel, args...)
}

func (h *Helper) Fatalf(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enable(FatalLevel) {
		return
	}
	h.Logger.Fields(h.fields).Logf(FatalLevel, template, args...)
}

func (h *Helper) WithError(err error) *Helper {
	fields := copyFields(h.fields)
	fields["error"] = err
	return &Helper{Logger: h.Logger, fields: fields}
}

func (h *Helper) WithFields(fields map[string]interface{}) *Helper {
	nfields := copyFields(fields)
	for k, v := range h.fields {
		nfields[k] = v
	}

	return &Helper{Logger: h.Logger, fields: nfields}
}
