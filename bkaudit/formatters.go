package bkaudit

// BaseFormatter - Interface for Formatter
type BaseFormatter interface {
	Format(
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
	eventID string,
	eventContent string,
	startTime int64,
	endTime int64,
	resultCode int64,
	resultContent string,
	extendData map[string]any,
) (auditEvent *AuditEvent, err error) {
	auditEvent = &AuditEvent{
		EventID:              eventID,
		EventContent:         eventContent,
		RequestID:            auditContext.RequestID,
		Username:             auditContext.Username,
		UserIdentifyType:     auditContext.UserIdentifyType,
		UserIdentifyTenantID: auditContext.UserIdentifyTenantId,
		StartTime:            startTime,
		EndTime:              endTime,
		AccessType:           auditContext.AccessType,
		AccessSourceIp:       auditContext.AccessSourceIp,
		AccessUserAgent:      auditContext.AccessUserAgent,
		ActionID:             action.ActionID,
		ResourceTypeID:       resourceType.ResourceTypeID,
		InstanceID:           instance.InstanceID,
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
