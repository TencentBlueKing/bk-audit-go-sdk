package bkaudit

const (
	AuditEventSignature   string = "bk_audit_event"
	AuditEventQueueLength int    = 50000
)

const (
	AccessTypeWeb     int8 = 0
	AccessTypeApi     int8 = 1
	AccessTypeConsole int8 = 2
	AccessTypeOther   int8 = -1
)

const (
	UserIdentifyTypePersonal int8 = 0
	UserIdentifyTypePlatform int8 = 1
	UserIdentifyTypeUnknown  int8 = -1
)
