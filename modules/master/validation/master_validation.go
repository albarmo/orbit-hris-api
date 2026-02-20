package validation

import (
	"github.com/go-playground/validator/v10"
)

type MasterValidation struct {
	validate *validator.Validate
}

func NewMasterValidation() *MasterValidation {
	validate := validator.New()
	return &MasterValidation{
		validate: validate,
	}
}

func (v *MasterValidation) ValidateDepartmentCreateRequest(req interface{}) error {
	return v.validate.Struct(req)
}

func (v *MasterValidation) ValidateDepartmentUpdateRequest(req interface{}) error {
	return v.validate.Struct(req)
}

func (v *MasterValidation) ValidateLocationCreateRequest(req interface{}) error {
	return v.validate.Struct(req)
}

func (v *MasterValidation) ValidateLocationUpdateRequest(req interface{}) error {
	return v.validate.Struct(req)
}

func (v *MasterValidation) ValidatePositionCreateRequest(req interface{}) error {
	return v.validate.Struct(req)
}

func (v *MasterValidation) ValidatePositionUpdateRequest(req interface{}) error {
	return v.validate.Struct(req)
}
