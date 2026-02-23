package query

import (
	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/Caknoooo/go-pagination"
	"gorm.io/gorm"
)

type Payroll struct {
	entities.Payroll
}

type PayrollFilter struct {
	pagination.BaseFilter
}

func (f *PayrollFilter) ApplyFilters(db *gorm.DB) *gorm.DB {
	// implement filter rules here if needed (e.g., by employee, period, date range)
	return db
}

func (f *PayrollFilter) GetTableName() string {
	return "payrolls"
}

func (f *PayrollFilter) GetSearchFields() []string {
	return []string{"employee_id"}
}

func (f *PayrollFilter) GetDefaultSort() string {
	return "generated_at desc"
}

func (f *PayrollFilter) GetIncludes() []string {
	return f.Includes
}

func (f *PayrollFilter) GetPagination() pagination.PaginationRequest {
	return f.Pagination
}

func (f *PayrollFilter) Validate() {
	var validIncludes []string
	allowedIncludes := f.GetAllowedIncludes()
	for _, include := range f.Includes {
		if allowedIncludes[include] {
			validIncludes = append(validIncludes, include)
		}
	}
	f.Includes = validIncludes
}

func (f *PayrollFilter) GetAllowedIncludes() map[string]bool {
	return map[string]bool{"Employee": true, "PayrollPeriod": true}
}
