package dto

type PayrollItemCreateRequest struct {
    PayrollID string  `json:"payroll_id" binding:"required,uuid"`
    Name      string  `json:"name" binding:"required"`
    Amount    float64 `json:"amount" binding:"required"`
}

type PayrollItemUpdateRequest struct {
    Name   *string  `json:"name"`
    Amount *float64 `json:"amount"`
}

type PayrollItemResponse struct {
    ID        string  `json:"id"`
    PayrollID string  `json:"payroll_id"`
    Name      string  `json:"name"`
    Amount    float64 `json:"amount"`
}
