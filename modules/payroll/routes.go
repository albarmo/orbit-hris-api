package payroll

import (
	payrollItemsController "github.com/Caknoooo/go-gin-clean-starter/modules/payroll/controller"
	payrollcontroller "github.com/Caknoooo/go-gin-clean-starter/modules/payroll/controller"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	payrollController := do.MustInvoke[payrollcontroller.PayrollController](injector)
	payrollItemCtrl := do.MustInvoke[payrollItemsController.PayrollItemController](injector)

	payrollRoutes := server.Group("/api/payroll")
	{
		payrollRoutes.POST("/", payrollController.Create)
		payrollRoutes.GET("/", payrollController.GetAll)
		payrollRoutes.GET("/:id", payrollController.GetByID)
		payrollRoutes.PUT("/:id", payrollController.Update)
		payrollRoutes.DELETE("/:id", payrollController.Delete)

		// Payroll items merged into payroll module
		payrollRoutes.POST("/items", payrollItemCtrl.Create)
		payrollRoutes.GET("/items", payrollItemCtrl.GetAll)
		payrollRoutes.GET("/items/:id", payrollItemCtrl.GetByID)
		payrollRoutes.PUT("/items/:id", payrollItemCtrl.Update)
		payrollRoutes.DELETE("/items/:id", payrollItemCtrl.Delete)
	}

	// Compatibility: keep legacy /api/payroll-items path working
	legacy := server.Group("/api/payroll-items")
	{
		legacy.POST("/", payrollItemCtrl.Create)
		legacy.GET("/", payrollItemCtrl.GetAll)
		legacy.GET("/:id", payrollItemCtrl.GetByID)
		legacy.PUT("/:id", payrollItemCtrl.Update)
		legacy.DELETE("/:id", payrollItemCtrl.Delete)
	}
}
