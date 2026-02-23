package dto

type ShiftCreateRequest struct {
    Name                 string    `json:"name" binding:"required"`
    StartTime            string    `json:"start_time" binding:"required"`
    EndTime              string    `json:"end_time" binding:"required"`
    LateToleranceMinutes int       `json:"late_tolerance_minutes" binding:"required"`
}

type ShiftUpdateRequest struct {
    Name                 *string `json:"name"`
    StartTime            *string `json:"start_time"`
    EndTime              *string `json:"end_time"`
    LateToleranceMinutes *int    `json:"late_tolerance_minutes"`
}

type ShiftResponse struct {
    ID                   string `json:"id"`
    Name                 string `json:"name"`
    StartTime            string `json:"start_time"`
    EndTime              string `json:"end_time"`
    LateToleranceMinutes int    `json:"late_tolerance_minutes"`
}

type EmployeeShiftCreateRequest struct {
    EmployeeID    string    `json:"employee_id" binding:"required,uuid"`
    ShiftID       string    `json:"shift_id" binding:"required,uuid"`
    EffectiveDate string    `json:"effective_date" binding:"required"`
}

type EmployeeShiftResponse struct {
    ID           string `json:"id"`
    EmployeeID   string `json:"employee_id"`
    ShiftID      string `json:"shift_id"`
    EffectiveDate string `json:"effective_date"`
}
