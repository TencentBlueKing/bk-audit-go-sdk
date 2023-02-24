package main

import (
	"github.com/TencentBlueKing/bk-audit-go-sdk/bkaudit"
	"time"
)

func main() {
	// add event
	context := bkaudit.AuditContext{Username: "admin"}
	instance := bkaudit.AuditInstance{
		InstanceID:          "host_01",
		InstanceName:        "主机一号",
		InstanceSensitivity: 0,
		InstanceData:        map[string]any{"ip": "127.0.0.1"},
		InstanceOriginData:  map[string]any{},
	}
	client.AddEvent(&viewHost, &host, &instance, &context, "", "", 0, 0, 0, "", map[string]any{})
	// wait channel
	time.Sleep(3 * time.Second)
}
