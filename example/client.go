package main

import (
	"github.com/TencentBlueKing/bk-audit-go-sdk/bkaudit"
	log "github.com/sirupsen/logrus"
)

// use single client
var client *bkaudit.EventClient

func init() {
	// init formatter
	var formatter = &bkaudit.Formatter{}
	// init exporter
	var exporters = []bkaudit.BaseExporter{&bkaudit.Exporter{}, &bkaudit.Exporter{}}
	// init client
	var err error
	client, err = bkaudit.InitEventClient("BkAppCode", "BkAppSecret", formatter, exporters, 0, nil)
	if err != nil {
		log.Info("init client failed")
		return
	}
}
