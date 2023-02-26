package bkaudit

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type EventLogger interface {
	Info(args ...interface{})
}

type RuntimeLogger interface {
	Error(args ...interface{})
	Info(args ...interface{})
}

var (
	eLog EventLogger
	rLog RuntimeLogger
)

type PlainTextFormatter struct{}

func (f *PlainTextFormatter) Format(entry *log.Entry) (b []byte, err error) {
	return append([]byte(entry.Message), '\n'), nil
}

func init() {
	_eLog := log.New()
	_eLog.SetFormatter(&PlainTextFormatter{})
	SetEventLogger(_eLog)
	_rLog := log.New()
	_rLog.SetFormatter(&log.JSONFormatter{
		TimestampFormat:   time.RFC3339,
		DisableTimestamp:  false,
		DisableHTMLEscape: true,
		PrettyPrint:       false,
		FieldMap: log.FieldMap{
			log.FieldKeyMsg: "message",
		},
	})
	SetRuntimeLogger(_rLog)
}

func SetRuntimeLogger(l RuntimeLogger) {
	rLog = l
}

func SetEventLogger(l EventLogger) {
	eLog = l
}
