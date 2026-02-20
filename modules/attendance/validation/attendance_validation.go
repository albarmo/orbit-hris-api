package validation

import (
	"github.com/Caknoooo/go-gin-clean-starter/modules/attendance/dto"
	"github.com/go-playground/validator/v10"
)

type AttendanceValidation struct {
	validate *validator.Validate
}

func NewAttendanceValidation() *AttendanceValidation {
	validate := validator.New()
	return &AttendanceValidation{
		validate: validate,
	}
}

func (v *AttendanceValidation) CheckIn(req dto.CheckInDTO) error {
	return v.validate.Struct(req)
}

func (v *AttendanceValidation) CheckOut(req dto.CheckOutDTO) error {
	return v.validate.Struct(req)
}

func (v *AttendanceValidation) UpdateAttendance(req dto.UpdateAttendanceDTO) error {
	return v.validate.Struct(req)
}

func (v *AttendanceValidation) ApproveAttendance(req dto.ApproveAttendanceDTO) error {
	return v.validate.Struct(req)
}
