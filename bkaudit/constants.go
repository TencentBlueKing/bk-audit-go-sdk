package bkaudit

const (
	AuditEventSignature   string = "bk_audit_event"
	AuditEventQueueLength int    = 50000
)

// AccessTypeEnum - Access Way Enum
type AccessTypeEnum int8

const (
	AccessTypeWeb     AccessTypeEnum = 0  // Access Through Web
	AccessTypeApi     AccessTypeEnum = 1  // Access Through ApiGateWay
	AccessTypeConsole AccessTypeEnum = 2  // Access Through Console
	AccessTypeOther   AccessTypeEnum = -1 // Access Through Another Way
)

// UserIdentifyTypeEnum - User Type Enum
type UserIdentifyTypeEnum int8

const (
	UserIdentifyTypePersonal UserIdentifyTypeEnum = 0  // Individual User
	UserIdentifyTypePlatform UserIdentifyTypeEnum = 1  // Platform User
	UserIdentifyTypeUnknown  UserIdentifyTypeEnum = -1 // Unknown User Type
)
