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
	// init export time
	exportTime := time.Now().Add(sleepTime)
	// init loop param
	var i int64 = 1
	// endless loop
	for {
		now := time.Now()
		// check run
		if now.Before(exportTime) {
			time.Sleep(sleepTime / 10)
			continue
		}
		// init loop param
		if totalRunTime == 0 || i <= totalRunTime {
			// export log
			var j int64
			for j = 0; j < exportEach; j++ {
				client.AddEvent(&action, &resourceType, &instance, &context, "", "", 0, 0, 0, "", map[string]any{})
			}
			// print log
			fmt.Printf(
				"StartTime: %s; CurrentTime: %s; ExportTotal: %d; CurrentRuntime: %d; TotalRunTime => %d\n",
				startTime,
				now.Format(time.RFC3339Nano),
				i*exportEach,
				i,
				totalRunTime,
			)
			// update loop param
			i++
			exportTime = exportTime.Add(sleepTime)
			continue
		}
		break
	}
}
