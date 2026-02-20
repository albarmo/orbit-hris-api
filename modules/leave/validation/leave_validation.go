package validation

import (
	"github.com/go-playground/validator/v10"
)

type LeaveValidation struct {
	validate *validator.Validate
}

func NewLeaveValidation() *LeaveValidation {
	validate := validator.New()
	return &LeaveValidation{
		validate: validate,
	}
}

func (v *LeaveValidation) ValidateCreate(req interface{}) error {
	return v.validate.Struct(req)
}

func (v *LeaveValidation) ValidateUpdate(req interface{}) error {
	return v.validate.Struct(req)
}
