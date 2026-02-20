package query

import (
	"github.com/Caknoooo/go-gin-clean-starter/pkg/pagination"
	"github.com/gin-gonic/gin"
)

// DepartmentFilter used to filter departments list
type DepartmentFilter struct {
    pagination.Filter
    Name string `form:"name"`
}

func (f *DepartmentFilter) Bind(ctx *gin.Context) {
    f.Filter.Bind(ctx)
    f.Name = ctx.Query("name")
}

// LocationFilter used to filter locations list
type LocationFilter struct {
    pagination.Filter
    Name string `form:"name"`
}

func (f *LocationFilter) Bind(ctx *gin.Context) {
    f.Filter.Bind(ctx)
    f.Name = ctx.Query("name")
}

// PositionFilter used to filter positions list
type PositionFilter struct {
    pagination.Filter
    Name         string `form:"name"`
    DepartmentID string `form:"department_id"`
}

func (f *PositionFilter) Bind(ctx *gin.Context) {
    f.Filter.Bind(ctx)
    f.Name = ctx.Query("name")
    f.DepartmentID = ctx.Query("department_id")
}
