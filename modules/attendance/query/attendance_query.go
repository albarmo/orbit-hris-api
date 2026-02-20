package query

import (
	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/Caknoooo/go-pagination"
	"gorm.io/gorm"
)

type Attendance struct {
	entities.Attendance
}

type AttendanceFilter struct {
	pagination.BaseFilter
}

// ApplyFilters applies custom filters to the query
func (f *AttendanceFilter) ApplyFilters(query *gorm.DB) *gorm.DB {
	return query
}

func (f *AttendanceFilter) GetTableName() string {
	return "attendances"
}

func (f *AttendanceFilter) GetSearchFields() []string {
	return []string{}
}

func (f *AttendanceFilter) GetDefaultSort() string {
	return "id asc"
}

func (f *AttendanceFilter) GetIncludes() []string {
	return f.Includes
}

func (f *AttendanceFilter) GetPagination() pagination.PaginationRequest {
	return f.Pagination
}

func (f *AttendanceFilter) Validate() {
	var validIncludes []string
	allowedIncludes := f.GetAllowedIncludes()
	for _, include := range f.Includes {
		if allowedIncludes[include] {
			validIncludes = append(validIncludes, include)
		}
	}
	f.Includes = validIncludes
}

func (f *AttendanceFilter) GetAllowedIncludes() map[string]bool {
	return map[string]bool{
		"Employee": true,
		"Location": true,
	}
}
