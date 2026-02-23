package leave_types

import (
	"github.com/Caknoooo/go-gin-clean-starter/modules/leave_types/controller"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	ctrl := do.MustInvoke[controller.LeaveTypeController](injector)

	r := server.Group("/api/leave-types")
	{
		r.POST("/", ctrl.Create)
		r.GET("/", ctrl.GetAll)
		r.GET("/:id", ctrl.GetByID)
		r.PUT("/:id", ctrl.Update)
		r.DELETE("/:id", ctrl.Delete)
	}
}
