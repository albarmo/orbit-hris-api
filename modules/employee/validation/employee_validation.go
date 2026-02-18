package validation

import (
	"github.com/Caknoooo/go-gin-clean-starter/modules/employee/dto"
	"github.com/go-playground/validator/v10"
)

type EmployeeValidation struct {
	validate *validator.Validate
}

func NewEmployeeValidation() *EmployeeValidation {
	validate := validator.New()
	return &EmployeeValidation{
		validate: validate,
	}
}

func (v *EmployeeValidation) ValidateEmployeeCreateRequest(req dto.EmployeeCreateRequest) error {
	return v.validate.Struct(req)
}

func (v *EmployeeValidation) ValidateEmployeeUpdateRequest(req dto.EmployeeUpdateRequest) error {
	return v.validate.Struct(req)
}