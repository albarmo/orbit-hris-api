package dto

type LeaveTypeCreateRequest struct {
    Name           string `json:"name" binding:"required"`
    MaxDaysPerYear int    `json:"max_days_per_year" binding:"required"`
    IsPaid         bool   `json:"is_paid"`
}

type LeaveTypeUpdateRequest struct {
    Name           *string `json:"name"`
    MaxDaysPerYear *int    `json:"max_days_per_year"`
    IsPaid         *bool   `json:"is_paid"`
}

type LeaveTypeResponse struct {
    ID             string `json:"id"`
    Name           string `json:"name"`
    MaxDaysPerYear int    `json:"max_days_per_year"`
    IsPaid         bool   `json:"is_paid"`
}
