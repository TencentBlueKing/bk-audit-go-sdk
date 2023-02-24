package main

import (
	"flag"
	"github.com/TencentBlueKing/bk-audit-go-sdk/bkaudit"
	"time"
)

func initRunParams() {
	logFileName = flag.String("name", "report", "")
	maxFileSize = flag.Float64("size", 1024, "") // 1024M
	maxBackupCount = flag.Int64("backup", 5, "")
	totalRunTime = flag.Int64("total", 0, "")
	sleepTime = flag.Duration("sleep", 1000*1000*1000, "")
	exportEach = flag.Int64("each", 1, "")
	flag.Parse()
	startTime = time.Now().Format(time.RFC3339Nano)
}

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
	exportLog()
	defer func() { _ = file.Close() }()
}
