package query

import (
	"github.com/Caknoooo/go-gin-clean-starter/pkg/pagination"
	"github.com/gin-gonic/gin"
)

// RoleFilter provides pagination + search by name for roles
type RoleFilter struct {
	pagination.Filter
	Name string `form:"name"`
}

func (f *RoleFilter) Bind(ctx *gin.Context) {
	f.Filter.Bind(ctx)
	f.Name = ctx.Query("name")
}

// PermissionFilter provides pagination + search by name for permissions
type PermissionFilter struct {
	pagination.Filter
	Name string `form:"name"`
}

func (f *PermissionFilter) Bind(ctx *gin.Context) {
	f.Filter.Bind(ctx)
	f.Name = ctx.Query("name")
}

