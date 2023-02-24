package main

import "github.com/TencentBlueKing/bk-audit-go-sdk/bkaudit"

var client *bkaudit.EventClient

func init() {
	var err error
	client, err = bkaudit.InitEventClient(
		"",
		"",
		&bkaudit.Formatter{},
		[]bkaudit.BaseExporter{&bkaudit.Exporter{}},
		0,
		nil,
	)
	if err != nil {
		panic("client init failed")
	}
}
