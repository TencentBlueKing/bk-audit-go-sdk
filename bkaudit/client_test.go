package bkaudit

import (
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
)

const (
	bkAppCode   = "bk-audit"
	bkAppSecret = "bk-audit"
	username    = "admin"
)

var action = AuditAction{ActionId: "view-host"}
var resourceType = AuditResource{ResourceTypeId: "host"}
var instance = AuditInstance{}
var context = AuditContext{Username: username}

func TestEventClient(t *testing.T) {
	// init client
	client, err := InitEventClient(bkAppCode, bkAppSecret, &Formatter{}, []BaseExporter{&Exporter{}}, 0, nil)
	if err != nil {
		t.Error(err)
	}
	// add event
	client.AddEvent(&action, &resourceType, &instance, &context, "", "", 0, 0, 0, "", map[string]any{})
	time.Sleep(1 * time.Second)
}

func TestValidateClient(t *testing.T) {
	// init client
	_, err := InitEventClient(bkAppCode, bkAppSecret, nil, []BaseExporter{}, 0, nil)
	if err == nil {
		t.Error("validate passed unexpected")
	}
}

func TestAddEventFailed(t *testing.T) {
	// init
	client, _ := InitEventClient(bkAppCode, bkAppSecret, &Formatter{}, []BaseExporter{&Exporter{}}, 0, nil)
	// add event
	// username invalid
	_context := AuditContext{}
	client.AddEvent(&action, &resourceType, &instance, &_context, "", "", 0, 0, 0, "", map[string]any{})
	// access type invalid
	_context2 := AuditContext{Username: username, AccessType: -2}
	client.AddEvent(&action, &resourceType, &instance, &_context2, "", "", 0, 0, 0, "", map[string]any{})
	// user identify type invalid
	_context3 := AuditContext{Username: username, UserIdentifyType: -2}
	client.AddEvent(&action, &resourceType, &instance, &_context3, "", "", 0, 0, 0, "", map[string]any{})
	// action id invalid
	_action := AuditAction{}
	client.AddEvent(&_action, &resourceType, &instance, &context, "", "", 0, 0, 0, "", map[string]any{})
	// json error
	_extendData := map[string]any{"func": func() {}}
	client.AddEvent(&action, &resourceType, &instance, &context, "", "", 0, 0, 0, "", _extendData)
	time.Sleep(1 * time.Second)
}

func TestCustomPreInit(t *testing.T) {
	var preInit = func() {
		log.Info("custom pre init")
	}
	_, _ = InitEventClient(bkAppCode, bkAppSecret, &Formatter{}, []BaseExporter{&Exporter{}}, 0, preInit)
}
