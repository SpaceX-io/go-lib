package log

import (
	"encoding/json"
	"fmt"
	"time"
)

var (
	DefaultSize   = 256
	DefaultFormat = TextFormat
)

type Record struct {
	Timestamp time.Time         `json:"timestamp"`
	Metadata  map[string]string `json:"metadata"`
	Message   interface{}       `json:"message"`
}

type Log interface {
	Read(...ReadOption) ([]Record, error)
	Write(Record) error
	Stream() (Steam, error)
}

func TextFormat(r Record) string {
	t := r.Timestamp.Format("2006-01-02 15:04:05")
	return fmt.Sprintf("%s %v", t, r.Message)
}

func JsonFormat(r Record) string {
	b, _ := json.Marshal(r)
	return string(b) + ""
}

type FormatFunc func(Record) string

type Steam interface {
	Chan() <-chan Record
	Stop() error
}
