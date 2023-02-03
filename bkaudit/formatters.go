package bkaudit

// BaseFormatter - Interface for Formatter
type BaseFormatter interface {
	Format(
		action *AuditAction,
		resourceType *AuditResource,
		instance *AuditInstance,
		auditContext *AuditContext,
		eventId string,
		eventContent string,
		startTime int64,
		endTime int64,
		resultCode int64,
		resultContent string,
		extendData map[string]any,
	) (auditEvent *AuditEvent, err error)
}

// Formatter - Build in Formatter
type Formatter struct{}

// Format - Generate Audit Event
func (f *Formatter) Format(
	action *AuditAction,
	resourceType *AuditResource,
	instance *AuditInstance,
	auditContext *AuditContext,
	eventId string,
	eventContent string,
	startTime int64,
	endTime int64,
	resultCode int64,
	resultContent string,
	extendData map[string]any,
) (auditEvent *AuditEvent, err error) {
	auditEvent = &AuditEvent{
		EventId:              eventId,
		EventContent:         eventContent,
		RequestId:            auditContext.RequestId,
		Username:             auditContext.Username,
		UserIdentifyType:     auditContext.UserIdentifyType,
		UserIdentifyTenantId: auditContext.UserIdentifyTenantId,
		StartTime:            startTime,
		EndTime:              endTime,
		AccessType:           auditContext.AccessType,
		AccessSourceIp:       auditContext.AccessSourceIp,
		AccessUserAgent:      auditContext.AccessUserAgent,
		ActionId:             action.ActionId,
		ResourceTypeId:       resourceType.ResourceTypeId,
		InstanceId:           instance.InstanceId,
		InstanceName:         instance.InstanceName,
		InstanceSensitivity:  instance.InstanceSensitivity,
		InstanceData:         instance.InstanceData,
		InstanceOriginData:   instance.InstanceOriginData,
		ResultCode:           resultCode,
		ResultContent:        resultContent,
		ExtendData:           extendData,
		AuditEventSignature:  AuditEventSignature,
	}
	// Check Event Validate
	if err := auditEvent.Validate(); err == nil {
		return auditEvent, nil
	} else {
		return nil, err
	}
}
