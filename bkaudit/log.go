package bkaudit

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type Logger interface {
	Error(args ...interface{})
	Info(args ...interface{})
}

var logger Logger

func init() {
	_log := &logrus.Logger{
		Out:          os.Stderr,
		Hooks:        make(logrus.LevelHooks),
		Level:        logrus.InfoLevel,
		ExitFunc:     os.Exit,
		ReportCaller: false,
		Formatter: &logrus.JSONFormatter{
			TimestampFormat:   time.RFC3339,
			DisableTimestamp:  false,
			DisableHTMLEscape: true,
			PrettyPrint:       false,
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyMsg: "message",
			},
		},
	}
	SetLogger(_log)
}

func SetLogger(l Logger) {
	logger = l
}
