package main

import "github.com/TencentBlueKing/bk-audit-go-sdk/bkaudit"

var (
	viewHost = bkaudit.AuditAction{ActionID: "view-host"}
	host     = bkaudit.AuditResource{ResourceTypeID: "host"}
	context  = bkaudit.AuditContext{Username: "admin"}
	instance = bkaudit.AuditInstance{
		InstanceID:          "host_01",
		InstanceName:        "主机一号",
		InstanceSensitivity: 0,
		InstanceData:        map[string]any{"ip": "127.0.0.1"},
		InstanceOriginData:  map[string]any{},
	}
)
