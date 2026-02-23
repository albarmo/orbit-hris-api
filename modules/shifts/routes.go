package shifts

import (
	"github.com/Caknoooo/go-gin-clean-starter/modules/shifts/controller"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	ctrl := do.MustInvoke[controller.ShiftController](injector)

	r := server.Group("/api/shifts")
	{
		r.POST("/", ctrl.Create)
		r.GET("/", ctrl.GetAll)
		r.GET("/:id", ctrl.GetByID)
		r.PUT("/:id", ctrl.Update)
		r.DELETE("/:id", ctrl.Delete)

		// employee shifts
		r.POST("/assign", ctrl.CreateEmployeeShift)
		r.GET("/employee/:employee_id", ctrl.GetEmployeeShifts)
		r.DELETE("/employee/:employee_id/:id", ctrl.DeleteEmployeeShift)
	}
}
