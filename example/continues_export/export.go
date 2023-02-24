package main

import (
	"fmt"
	"github.com/TencentBlueKing/bk-audit-go-sdk/bkaudit"
	"github.com/google/uuid"
	"time"
)

func exportLog() {
	// init export time
	exportTime := time.Now().Add(*sleepTime)
	// init loop param
	var i int64 = 1
	// endless loop
	for {
		now := time.Now()
		// check run
		if now.Before(exportTime) {
			time.Sleep(*sleepTime / 10)
			continue
		}
		// init loop param
		if *totalRunTime == 0 || i <= *totalRunTime {
			// export log
			var j int64
			for j = 0; j < *exportEach; j++ {
				eventID := fmt.Sprintf("%d.%s", time.Now().UnixNano(), uuid.NewString())
				instance := bkaudit.AuditInstance{InstanceID: "z0001"}
				context := bkaudit.AuditContext{Username: "admin"}
				client.AddEvent(&action, &resourceType, &instance, &context, eventID, "", 0, 0, 0, "", map[string]any{})
			}
			// print log
			bkaudit.RuntimeLog.Infof(
				"StartTime: %s; CurrentTime: %s; ExportTotal: %d; CurrentRuntime: %d; TotalRunTime => %d",
				startTime,
				now.Format(time.RFC3339Nano),
				i*(*exportEach),
				i,
				*totalRunTime,
			)
			// update loop param
			i++
			exportTime = exportTime.Add(*sleepTime)
			checkFileAfterWrite()
			continue
		}
		break
	}
}
