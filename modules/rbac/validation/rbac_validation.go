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
