package main

import "github.com/TencentBlueKing/bk-audit-go-sdk/bkaudit"

var (
	viewHost = bkaudit.AuditAction{ActionID: "view-host"}
	host     = bkaudit.AuditResource{ResourceTypeID: "host"}
)
