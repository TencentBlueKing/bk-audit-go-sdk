package bkaudit

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type PlainTextFormatter struct{}

var (
	EventLog   *log.Logger
	RuntimeLog *log.Logger
)

func (f *PlainTextFormatter) Format(entry *log.Entry) (b []byte, err error) {
	return append([]byte(entry.Message), '\n'), nil
}

func init() {
	EventLog = log.New()
	EventLog.SetFormatter(&PlainTextFormatter{})
	RuntimeLog = log.New()
	RuntimeLog.SetFormatter(&log.JSONFormatter{
		TimestampFormat:   time.RFC3339,
		DisableTimestamp:  false,
		DisableHTMLEscape: true,
		PrettyPrint:       false,
		FieldMap: log.FieldMap{
			log.FieldKeyMsg: "message",
		},
	})
}
