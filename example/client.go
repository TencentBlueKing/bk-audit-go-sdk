package main

import (
	"fmt"
	"github.com/TencentBlueKing/bk-audit-go-sdk/bkaudit"
)

var (
	err    error
	client *bkaudit.EventClient
)

var formatter = &bkaudit.Formatter{}
var exporters = []bkaudit.BaseExporter{&bkaudit.Exporter{}, &bkaudit.Exporter{}}

func initClient() {
	// init client
	client, err = bkaudit.InitEventClient("BkAppCode", "BkAppSecret", formatter, exporters, 0, nil)
	if err != nil {
		fmt.Println("init client failed")
		return
	}
}
