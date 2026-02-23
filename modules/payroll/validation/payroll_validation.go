package validation

import (
	"github.com/go-playground/validator/v10"
)

type PayrollValidation struct {
	validate *validator.Validate
}

func NewPayrollValidation() *PayrollValidation {
	validate := validator.New()
	return &PayrollValidation{
		validate: validate,
	}
}
