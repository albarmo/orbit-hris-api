package payroll_items

import (
	"github.com/Caknoooo/go-gin-clean-starter/modules/payroll_items/controller"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
    ctrl := do.MustInvoke[controller.PayrollItemController](injector)

    r := server.Group("/api/payroll-items")
    {
        r.POST("/", ctrl.Create)
        r.GET("/", ctrl.GetAll)
        r.GET("/:id", ctrl.GetByID)
        r.PUT("/:id", ctrl.Update)
        r.DELETE("/:id", ctrl.Delete)
    }
}
