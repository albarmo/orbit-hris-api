package validation

import (
	"github.com/go-playground/validator/v10"
)

type RbacValidation struct {
	validate *validator.Validate
}

func NewRbacValidation() *RbacValidation {
	validate := validator.New()
	return &RbacValidation{
		validate: validate,
	}
}

func (v *RbacValidation) ValidateRoleCreate(req interface{}) error {
	return v.validate.Struct(req)
}

func (v *RbacValidation) ValidateRoleUpdate(req interface{}) error {
	return v.validate.Struct(req)
}

func (v *RbacValidation) ValidatePermissionCreate(req interface{}) error {
	return v.validate.Struct(req)
}

func (v *RbacValidation) ValidatePermissionUpdate(req interface{}) error {
	return v.validate.Struct(req)
}

func (v *RbacValidation) ValidateAssignRole(req interface{}) error {
	return v.validate.Struct(req)
}
