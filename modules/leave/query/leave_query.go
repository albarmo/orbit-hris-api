package query

import (
	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/Caknoooo/go-pagination"
	"gorm.io/gorm"
)

type Leave struct {
	entities.Leave
}

type LeaveFilter struct {
	pagination.BaseFilter
}

func (f *LeaveFilter) ApplyFilters(db *gorm.DB) *gorm.DB {
	return db
}

func (f *LeaveFilter) GetTableName() string {
	return "leaves"
}

func (f *LeaveFilter) GetSearchFields() []string {
	return []string{"reason"}
}

func (f *LeaveFilter) GetDefaultSort() string {
	return "created_at desc"
}

func (f *LeaveFilter) GetIncludes() []string {
	return f.Includes
}

func (f *LeaveFilter) GetPagination() pagination.PaginationRequest {
	return f.Pagination
}

func (f *LeaveFilter) Validate() {
	var validIncludes []string
	allowedIncludes := f.GetAllowedIncludes()
	for _, include := range f.Includes {
		if allowedIncludes[include] {
			validIncludes = append(validIncludes, include)
		}
	}
	f.Includes = validIncludes
}

func (f *LeaveFilter) GetAllowedIncludes() map[string]bool {
	return map[string]bool{"Employee": true}
}
