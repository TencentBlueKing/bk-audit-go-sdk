package bkaudit

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

// AuditEvent - Audit Event Contains All Audit Fields
type AuditEvent struct {
	EventID              string               `json:"event_id" validate:"event_id"`
	EventContent         string               `json:"event_content"`
	RequestID            string               `json:"request_id"`
	Username             string               `json:"username" validate:"required"`
	UserIdentifyType     UserIdentifyTypeEnum `json:"user_identify_type" validate:"user_identify_type"`
	UserIdentifyTenantID string               `json:"user_identify_tenant_id"`
	StartTime            int64                `json:"start_time" validate:"milli_timestamp"`
	EndTime              int64                `json:"end_time" validate:"milli_timestamp"`
	BkAppCode            string               `json:"bk_app_code"`
	AccessType           AccessTypeEnum       `json:"access_type" validate:"access_type"`
	AccessSourceIp       string               `json:"access_source_ip"`
	AccessUserAgent      string               `json:"access_user_agent"`
	ActionID             string               `json:"action_id" validate:"required"`
	ResourceTypeID       string               `json:"resource_type_id"`
	InstanceID           string               `json:"instance_id"`
	InstanceName         string               `json:"instance_name"`
	InstanceSensitivity  int64                `json:"instance_sensitivity"`
	InstanceData         map[string]any       `json:"instance_data"`
	InstanceOriginData   map[string]any       `json:"instance_origin_data"`
	ResultCode           int64                `json:"result_code"`
	ResultContent        string               `json:"result_content"`
	ExtendData           map[string]any       `json:"extend_data"`
	AuditEventSignature  string               `json:"audit_event_signature"`
}

// String - Trans Audit Event to String
func (auditEvent *AuditEvent) String() (string, error) {
	// trans struct to string
	data, err := json.Marshal(auditEvent)
	return string(data), err
}

func (auditEvent *AuditEvent) Validate() error {
	return validate.Struct(auditEvent)
}

var validate *validator.Validate

func initValidator() {
	validate = validator.New()
	_ = validate.RegisterValidation("event_id", validateEventID)
	_ = validate.RegisterValidation("milli_timestamp", validateMilliTimestamp)
	_ = validate.RegisterValidation("user_identify_type", validateUserIdentifyType)
	_ = validate.RegisterValidation("access_type", validateAccessType)
}

func validateEventID(field validator.FieldLevel) bool {
	eventID, ok := field.Field().Interface().(string)
	if !ok {
		return false
	}
	if eventID == "" {
		field.Field().SetString(uuid.NewString())
	}
	return true
}

func validateMilliTimestamp(field validator.FieldLevel) bool {
	timestamp, ok := field.Field().Interface().(int64)
	if !ok {
		return false
	}
	if timestamp == 0 {
		field.Field().SetInt(time.Now().UnixMilli())
	}
	return true
}

func validateUserIdentifyType(field validator.FieldLevel) bool {
	userIdentifyType, ok := field.Field().Interface().(UserIdentifyTypeEnum)
	if !ok {
		return false
	}
	allTypes := []any{UserIdentifyTypePersonal, UserIdentifyTypePlatform, UserIdentifyTypeUnknown}
	return checkValueInSlice(userIdentifyType, allTypes)
}

func validateAccessType(field validator.FieldLevel) bool {
	accessType, ok := field.Field().Interface().(AccessTypeEnum)
	if !ok {
		return false
	}
	allTypes := []any{AccessTypeWeb, AccessTypeApi, AccessTypeConsole, AccessTypeOther}
	return checkValueInSlice(accessType, allTypes)
}

func checkValueInSlice(value any, allValue []any) bool {
	for index := range allValue {
		if allValue[index] == value {
			return true
		}
	}
	return false
}

// AuditAction - IAM Action
type AuditAction struct {
	ActionID string `json:"action_id"`
}

// AuditResource - IAM Resource Type
type AuditResource struct {
	ResourceTypeID string `json:"resource_type_id"`
}

// AuditInstance - Instance for Audit Event
type AuditInstance struct {
	InstanceID          string         `json:"instance_id"`
	InstanceName        string         `json:"instance_name"`
	InstanceSensitivity int64          `json:"instance_sensitivity"`
	InstanceData        map[string]any `json:"instance_data"`
	InstanceOriginData  map[string]any `json:"instance_origin_data"`
}

// AuditContext - Context for Audit Event
type AuditContext struct {
	Username             string               `json:"username"`
	RequestID            string               `json:"requestID"`
	AccessType           AccessTypeEnum       `json:"accessType"`
	AccessSourceIp       string               `json:"accessSourceIp"`
	AccessUserAgent      string               `json:"accessUserAgent"`
	UserIdentifyType     UserIdentifyTypeEnum `json:"userIdentifyType"`
	UserIdentifyTenantId string               `json:"userIdentifyTenantId"`
	ExtraData            map[string]any       `json:"extra_data"`
}
