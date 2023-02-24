package main

import (
	"github.com/TencentBlueKing/bk-audit-go-sdk/bkaudit"
)

var (
	action       = bkaudit.AuditAction{ActionID: "view-file"}
	resourceType = bkaudit.AuditResource{ResourceTypeID: "host"}
)
