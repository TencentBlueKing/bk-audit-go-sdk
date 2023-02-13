package main

import "fmt"

func main() {
	initRunParams()
	fmt.Printf(
		"Init Run Params Finished; TotalRuntime => %d; ExportEach => %d; SleepTime => %s\n",
		totalRunTime,
		exportEach,
		sleepTime,
	)
	initLogFile()
	fmt.Printf("Init Log File Finished; File Name => %s\n", logFileName)
	initClient()
	exportLog()
}
