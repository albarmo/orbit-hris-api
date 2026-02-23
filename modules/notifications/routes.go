package notifications

import (
	"github.com/Caknoooo/go-gin-clean-starter/modules/notifications/controller"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	ctrl := do.MustInvoke[controller.NotificationController](injector)

	r := server.Group("/api/notifications")
	{
		r.POST("/", ctrl.Create)
		r.GET("/user/:user_id", ctrl.GetByUser)
		r.GET("/:id", ctrl.GetByID)
		r.PUT("/:id", ctrl.Update)
		r.DELETE("/:id", ctrl.Delete)
	}
}
