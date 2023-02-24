// Package bkaudit - Generate Audit Event and Export

package bkaudit

import (
	"errors"
)

// EventClient - Client to Generate Event
type EventClient struct {
	BkAppCode   string
	BkAppSecret string
	formatter   BaseFormatter
	exporters   []BaseExporter
	queues      []BaseQueue
}

func (client *EventClient) check() (err error) {
	// Formatter and Exporter should be initialized before use
	if client.formatter == nil || len(client.exporters) == 0 {
		return errors.New("formatter or exporter not set")
	}
	return nil
}

// AddEvent - Generate Audit Event and Export to Stdout, Log ...
func (client *EventClient) AddEvent(
	action *AuditAction,
	resourceType *AuditResource,
	instance *AuditInstance,
	auditContext *AuditContext,
	eventID string,
	eventContent string,
	startTime int64,
	endTime int64,
	resultCode int64,
	resultContent string,
	extendData map[string]any,
) {
	// Build Audit Event
	auditEvent, err := (client.formatter).Format(
		action,
		resourceType,
		instance,
		auditContext,
		eventID,
		eventContent,
		startTime,
		endTime,
		resultCode,
		resultContent,
		extendData,
	)
	if err != nil {
		RuntimeLog.Error("format event failed: ", err)
		return
	}
	// Add BkAppCode
	auditEvent.BkAppCode = client.BkAppCode
	// Export Audit Event
	for _, q := range client.queues {
		q <- auditEvent
	}
}

// InitEventClient - Init an Event Client
func InitEventClient(
	bkAppCode string,
	bkAppSecret string,
	formatter BaseFormatter,
	exporters []BaseExporter,
	queueLength int,
	preInit func(),
) (client *EventClient, err error) {
	// pre init
	if preInit != nil {
		preInit()
	}
	// Init Validator
	initValidator()
	// Init Queue
	if queueLength == 0 {
		queueLength = AuditEventQueueLength
	}
	// Start Exporter
	queues := make([]BaseQueue, len(exporters))
	for i := 0; i < len(exporters); i++ {
		q := make(BaseQueue, queueLength)
		queues[i] = q
		go exporters[i].Export(q)
	}
	// Init Client
	client = &EventClient{
		BkAppCode:   bkAppCode,
		BkAppSecret: bkAppSecret,
		formatter:   formatter,
		exporters:   exporters,
		queues:      queues,
	}
	// Check Client Valid
	if err = client.check(); err != nil {
		return nil, err
	}
	return client, nil
}
