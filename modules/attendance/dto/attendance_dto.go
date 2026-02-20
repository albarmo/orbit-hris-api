package dto

import "github.com/google/uuid"

type CheckInDTO struct {
	EmployeeID uuid.UUID `json:"employee_id" binding:"required"`
	LocationID uuid.UUID `json:"location_id" binding:"required"`
}

type CheckOutDTO struct {
	EmployeeID uuid.UUID `json:"employee_id" binding:"required"`
}

type UpdateAttendanceDTO struct {
	Status string `json:"status" binding:"required"`
}

type ApproveAttendanceDTO struct {
	Status string `json:"status" binding:"required,oneof=approved rejected"`
}
