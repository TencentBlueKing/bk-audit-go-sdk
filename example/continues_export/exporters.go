package main

import (
	"github.com/TencentBlueKing/bk-audit-go-sdk/bkaudit"
	log "github.com/sirupsen/logrus"
	"time"
)

type fileExporter struct{}

func (e *fileExporter) Export(queue bkaudit.BaseQueue) {
	for event := range queue {
		// get string data
		data, err := event.String()
		if err != nil {
			log.Error("export event failed: ", err)
			return
		}
		// Directly Export to Log
		_, err = file.Write([]byte(data + "\n"))
		if err != nil {
			log.Error("export event failed: ", err)
		}
	}
}

func exportLog() {
	var i, j int64
	for i = 1; totalRunTime == 0 || i <= totalRunTime; i++ {
		go func() {
			for j = 0; j < exportEach; j++ {
				client.AddEvent(&action, &resourceType, &instance, &context, "", "", 0, 0, 0, "", map[string]any{})
			}
			log.Printf("CurrentRuntime: %d; TotalRunTime => %d", i, totalRunTime)
		}()
		time.Sleep(sleepTime)
	}
}
