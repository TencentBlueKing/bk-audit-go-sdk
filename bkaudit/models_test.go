package bkaudit

import (
	"github.com/go-playground/validator/v10"
	"testing"
)

type User struct {
	Username  string `validate:"timestamp"`
	Age       int8   `validate:"age"`
	FirstName string `validate:"at"`
	LastName  string `validate:"uit"`
}

func TestValidateFieldError(t *testing.T) {
	validate := validator.New()
	_ = validate.RegisterValidation("age", validateEventId)
	_ = validate.RegisterValidation("timestamp", validateMilliTimestamp)
	_ = validate.RegisterValidation("at", validateAccessType)
	_ = validate.RegisterValidation("uit", validateUserIdentifyType)
	err := validate.Struct(&User{})
	if err == nil {
		t.Error("validate passed unexpected")
	}
}
