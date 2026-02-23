package dto

type ExpenseCategoryCreateRequest struct {
    Name string `json:"name" binding:"required"`
    LimitAmount *float64 `json:"limit_amount,omitempty"`
}

type ExpenseCategoryResponse struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

type ExpenseCreateRequest struct {
    EmployeeID string  `json:"employee_id" binding:"required,uuid"`
    CategoryID string  `json:"category_id" binding:"required,uuid"`
    Amount     float64 `json:"amount" binding:"required"`
    Description string `json:"description"`
    ReceiptURL string  `json:"receipt_url"`
}

type ExpenseUpdateRequest struct {
    CategoryID  *string  `json:"category_id" binding:"omitempty,uuid"`
    Amount      *float64 `json:"amount"`
    Description *string  `json:"description"`
    ReceiptURL  *string  `json:"receipt_url"`
    Status      *string  `json:"status"`
}

type ExpenseResponse struct {
    ID          string  `json:"id"`
    EmployeeID  string  `json:"employee_id"`
    CategoryID  string  `json:"category_id"`
    Amount      float64 `json:"amount"`
    Description string  `json:"description"`
    ReceiptURL  string  `json:"receipt_url"`
    Status      string  `json:"status"`
}
