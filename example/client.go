package main

import (
	"github.com/TencentBlueKing/bk-audit-go-sdk/bkaudit"
	log "github.com/sirupsen/logrus"
	"os"
)

// use single client
var client *bkaudit.EventClient

// logger formatter
type plainTextLogFormatter struct{}

func (f *plainTextLogFormatter) Format(entry *log.Entry) ([]byte, error) {
	return append([]byte(entry.Message), '\n'), nil
}

func init() {
	// init formatter
	var formatter = &bkaudit.EventFormatter{}
	// init exporter
	var logger = &log.Logger{
		Out:          os.Stderr,
		Formatter:    &plainTextLogFormatter{},
		Hooks:        make(log.LevelHooks),
		Level:        log.InfoLevel,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}
	var exporters = []bkaudit.Exporter{&bkaudit.LoggerExporter{Logger: logger}}
	// init client
	var err error
	client, err = bkaudit.InitEventClient("BkAppCode", "BkAppSecret", formatter, exporters, 0, nil)
	if err != nil {
		log.Info("init client failed")
		return
	}
}
