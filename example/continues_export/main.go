package main

import "fmt"

func main() {
	initRunParams()
	fmt.Printf("Init Run Params Finished; TotalRuntime => %d; SleepTime => %s\n", totalRunTime, sleepTime)
	initLogFile()
	fmt.Printf("Init Log File Finished; File Name => %s\n", logFileName)
	initClient()
	exportLog()
}
