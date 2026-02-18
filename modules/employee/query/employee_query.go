package query

import (
	"github.com/Caknoooo/go-pagination"
	"gorm.io/gorm"
)

type Employee struct {
	ID               string `json:"id"`
	UserID           string `json:"user_id"`
	EmployeeCode     string `json:"employee_code"`
	SupervisorID     string `json:"supervisor_id"`
	DepartmentID     string `json:"department_id"`
	PositionID       string `json:"position_id"`
	JoinDate         string `json:"join_date"`
	EndDate          string `json:"end_date"`
	EmploymentType   string `json:"employment_type"`
	EmploymentStatus string `json:"employment_status"`
	ProbationEndDate string `json:"probation_end_date"`
}

type EmployeeFilter struct {
	pagination.BaseFilter
}

func (f *EmployeeFilter) ApplyFilters(query *gorm.DB) *gorm.DB {
	// Apply your filters here
	return query
}

func (f *EmployeeFilter) GetTableName() string {
	return "employees"
}

func (f *EmployeeFilter) GetSearchFields() []string {
	return []string{"employee_code"}
}

func (f *EmployeeFilter) GetDefaultSort() string {
	return "created_at desc"
}

func (f *EmployeeFilter) GetIncludes() []string {
	return f.Includes
}

func (f *EmployeeFilter) GetPagination() pagination.PaginationRequest {
	return f.Pagination
}

func (f *EmployeeFilter) Validate() {
	var validIncludes []string
	allowedIncludes := f.GetAllowedIncludes()
	for _, include := range f.Includes {
		if allowedIncludes[include] {
			validIncludes = append(validIncludes, include)
		}
	}
	f.Includes = validIncludes
}

func (f *EmployeeFilter) GetAllowedIncludes() map[string]bool {
	return map[string]bool{
		"User":       true,
		"Department": true,
		"Position":   true,
		"Supervisor": true,
	}
}
