package dto

const (
	MESSAGE_FAILED_GET_DATA_FROM_BODY = "failed get data from body"
	MESSAGE_SUCCESS_GET_DATA         = "success get data"
)

type (
	PayrollCreateRequest struct {
		EmployeeID      string  `json:"employee_id" binding:"required,uuid"`
		PayrollPeriodID string  `json:"payroll_period_id" binding:"required,uuid"`
		BasicSalary     float64 `json:"basic_salary" binding:"required"`
		TotalAllowance  float64 `json:"total_allowance"`
		TotalDeduction  float64 `json:"total_deduction"`
	}

	PayrollUpdateRequest struct {
		BasicSalary     *float64 `json:"basic_salary"`
		TotalAllowance  *float64 `json:"total_allowance"`
		TotalDeduction  *float64 `json:"total_deduction"`
		PayrollPeriodID *string  `json:"payroll_period_id" binding:"omitempty,uuid"`
	}

	PayrollResponse struct {
		ID              string  `json:"id"`
		EmployeeID      string  `json:"employee_id"`
		PayrollPeriodID string  `json:"payroll_period_id"`
		BasicSalary     float64 `json:"basic_salary"`
		TotalAllowance  float64 `json:"total_allowance"`
		TotalDeduction  float64 `json:"total_deduction"`
		NetSalary       float64 `json:"net_salary"`
		GeneratedAt     string  `json:"generated_at"`
	}
)
