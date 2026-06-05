package tests

import (
	"testing"

	"github.com/Caknoooo/go-gin-clean-starter/modules/attendance/dto"
	"github.com/Caknoooo/go-gin-clean-starter/modules/attendance/validation"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAttendanceValidation_NewAttendanceValidation(t *testing.T) {
	validator := validation.NewAttendanceValidation()
	require.NotNil(t, validator)
}

func TestAttendanceValidation_CheckIn_NoError(t *testing.T) {
	validator := validation.NewAttendanceValidation()

	err := validator.CheckIn(dto.CheckInDTO{
		EmployeeID: uuid.New(),
		LocationID: uuid.New(),
	})

	assert.NoError(t, err)
}

func TestAttendanceValidation_CheckOut_NoError(t *testing.T) {
	validator := validation.NewAttendanceValidation()

	err := validator.CheckOut(dto.CheckOutDTO{
		EmployeeID: uuid.New(),
	})

	assert.NoError(t, err)
}

func TestAttendanceValidation_UpdateAttendance_NoError(t *testing.T) {
	validator := validation.NewAttendanceValidation()

	err := validator.UpdateAttendance(dto.UpdateAttendanceDTO{
		Status: "present",
	})

	assert.NoError(t, err)
}

func TestAttendanceValidation_ApproveAttendance_NoError(t *testing.T) {
	validator := validation.NewAttendanceValidation()

	err := validator.ApproveAttendance(dto.ApproveAttendanceDTO{
		Status: "approved",
	})

	assert.NoError(t, err)
}
