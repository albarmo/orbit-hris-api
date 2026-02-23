package payroll

import (
	"github.com/Caknoooo/go-gin-clean-starter/modules/payroll/controller"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	payrollController := do.MustInvoke[controller.PayrollController](injector)

	payrollRoutes := server.Group("/api/payroll")
	{
		payrollRoutes.POST("/", payrollController.Create)
		payrollRoutes.GET("/", payrollController.GetAll)
		payrollRoutes.GET("/:id", payrollController.GetByID)
		payrollRoutes.PUT("/:id", payrollController.Update)
		payrollRoutes.DELETE("/:id", payrollController.Delete)
	}
}
