package main

import (
	"github.com/TencentBlueKing/bk-audit-go-sdk/bkaudit"
)

func main() {
	initRunParams()
	bkaudit.RuntimeLog.Infof("StartAt %s", startTime)
	bkaudit.RuntimeLog.Infof(
		"Init Run Params Finished; TotalRuntime => %d; ExportEach => %d; SleepTime => %s",
		*totalRunTime,
		*exportEach,
		*sleepTime,
	)
	initLogFile()
	bkaudit.RuntimeLog.Infof("Init Log File Finished; File Name => %s", *logFileName)
	initClient()
	exportLog()
	defer func() { _ = file.Close() }()
}
