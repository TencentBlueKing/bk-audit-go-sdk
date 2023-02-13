package main

import (
	"github.com/TencentBlueKing/bk-audit-go-sdk/bkaudit"
	"os"
	"time"
)

var (
	logFileName  string
	totalRunTime int64
	sleepTime    time.Duration
	exportEach   int64
	file         *os.File
	client       *bkaudit.EventClient
	action       = bkaudit.AuditAction{ActionID: "view-file"}
	resourceType = bkaudit.AuditResource{ResourceTypeID: "host"}
	instance     = bkaudit.AuditInstance{InstanceID: "z0001"}
	context      = bkaudit.AuditContext{Username: "admin"}
)
