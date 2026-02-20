package dto

import (
	"time"

	"github.com/google/uuid"
)

const (
	MESSAGE_FAILED_GET_DATA_FROM_BODY = "failed get data from body"
	MESSAGE_SUCCESS_GET_DATA         = "success get data"
)

type (
	LeaveCreateRequest struct {
		EmployeeID uuid.UUID `json:"employee_id" binding:"required"`
		StartDate  time.Time `json:"start_date" binding:"required"`
		EndDate    time.Time `json:"end_date" binding:"required"`
		Reason     string    `json:"reason" binding:"required"`
	}

	LeaveUpdateRequest struct {
		Status string `json:"status" binding:"required,oneof=pending approved rejected"`
		Reason string `json:"reason"`
	}

	LeaveResponse struct {
		ID         uuid.UUID `json:"id"`
		EmployeeID uuid.UUID `json:"employee_id"`
		StartDate  time.Time `json:"start_date"`
		EndDate    time.Time `json:"end_date"`
		Reason     string    `json:"reason"`
		Status     string    `json:"status"`
	}
)
