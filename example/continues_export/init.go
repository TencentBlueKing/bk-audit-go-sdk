package main

import (
	"flag"
	"github.com/TencentBlueKing/bk-audit-go-sdk/bkaudit"
	"os"
)

func initRunParams() {
	_logFileName := flag.String("name", "report.log", "")
	_totalRunTime := flag.Int64("total", 0, "")
	_sleepTime := flag.Duration("sleep", 1000*1000*1000, "")
	flag.Parse()
	logFileName, totalRunTime, sleepTime = *_logFileName, *_totalRunTime, *_sleepTime
}

func initClient() {
	var err error
	client, err = bkaudit.InitEventClient("", "", &bkaudit.Formatter{}, []bkaudit.BaseExporter{&fileExporter{}}, 0, nil)
	if err != nil {
		panic("client init failed")
	}
}

func initLogFile() {
	var err error
	file, err = os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic("open log file error")
	}
}