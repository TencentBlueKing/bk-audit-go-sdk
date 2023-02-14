package main

import (
	"fmt"
	"github.com/TencentBlueKing/bk-audit-go-sdk/bkaudit"
	"time"
)

type fileExporter struct{}

func (e *fileExporter) Export(queue bkaudit.BaseQueue) {
	for event := range queue {
		// get string data
		data, err := event.String()
		if err != nil {
			fmt.Printf("export event failed: %s\n", err)
			return
		}
		// Directly Export to Log
		_, err = file.Write([]byte(data + "\n"))
		if err != nil {
			fmt.Printf("export event failed: %s\n", err)
		}
	}
}

func exportLog() {
	var i, j int64
	for i = 1; totalRunTime == 0 || i <= totalRunTime; i++ {
		for j = 0; j < exportEach; j++ {
			client.AddEvent(&action, &resourceType, &instance, &context, "", "", 0, 0, 0, "", map[string]any{})
		}
		fmt.Printf(
			"StartTime: %s; CurrentTime: %s; ExportTotal: %d; CurrentRuntime: %d; TotalRunTime => %d\n",
			startTime,
			time.Now().Format(time.RFC3339),
			i*exportEach,
			i,
			totalRunTime,
		)
		time.Sleep(sleepTime)
	}
}
