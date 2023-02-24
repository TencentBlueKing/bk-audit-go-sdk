package main

import "time"

var (
	startTime        string
	logFileName      *string
	maxFileSize      *float64
	maxBackupCount   *int64
	currentFileIndex = int64(1)
	totalRunTime     *int64
	sleepTime        *time.Duration
	exportEach       *int64
)
