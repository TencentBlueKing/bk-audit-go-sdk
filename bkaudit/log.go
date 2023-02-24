package bkaudit

import (
	log "github.com/sirupsen/logrus"
)

type PlainTextFormatter struct{}

func (f *PlainTextFormatter) Format(entry *log.Entry) (b []byte, err error) {
	return append([]byte(entry.Message), '\n'), nil
}

func initLog() {
	log.SetFormatter(&PlainTextFormatter{})
}
