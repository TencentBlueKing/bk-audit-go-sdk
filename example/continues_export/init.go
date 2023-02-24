package main

import (
	"flag"
	"fmt"
	"github.com/TencentBlueKing/bk-audit-go-sdk/bkaudit"
	"os"
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

func initClient() {
	var err error
	client, err = bkaudit.InitEventClient(
		"",
		"",
		&bkaudit.Formatter{},
		[]bkaudit.BaseExporter{&bkaudit.Exporter{}},
		0,
		nil,
	)
	if err != nil {
		panic("client init failed")
	}
}

func initLogFile() {
	checkFileBeforeWrite()
	if file != nil {
		oldFile = file
		go func() {
			time.Sleep(10 * time.Second)
			_ = oldFile.Close()
		}()
	}
	_file, err := os.OpenFile(buildLogFileName(currentFileIndex), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic("open log file error")
	}
	file = _file
	bkaudit.EventLog.SetOutput(file)
}

func getNextFileIndex() {
	if currentFileIndex >= *maxBackupCount {
		currentFileIndex = 1
		return
	}
	currentFileIndex++
}

func checkFileBeforeWrite() {
	nextFile, _ := os.Open(buildLogFileName(currentFileIndex))
	defer func() { _ = nextFile.Close() }()
	// get file stat info
	stat, err := nextFile.Stat()
	// if error, open new file
	if err != nil || float64(stat.Size()) > *maxFileSize*1024*1024 {
		removeOldLogFile()
	}
}

func checkFileAfterWrite() {
	// get file stat info
	stat, err := file.Stat()
	// if error, open new file
	if err != nil || float64(stat.Size()) > *maxFileSize*1024*1024 {
		getNextFileIndex()
		initLogFile()
	}
}

func buildLogFileName(index int64) string {
	return fmt.Sprintf("%s.%d.log", *logFileName, index)
}

func removeOldLogFile() {
	fileName := buildLogFileName(currentFileIndex)
	_ = os.Remove(fileName)
}
