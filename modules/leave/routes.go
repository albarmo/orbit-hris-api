package leave

import (
	"github.com/Caknoooo/go-gin-clean-starter/middlewares"
	"github.com/Caknoooo/go-gin-clean-starter/modules/auth/service"
	"github.com/Caknoooo/go-gin-clean-starter/modules/leave/controller"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/constants"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	leaveController := do.MustInvoke[controller.LeaveController](injector)

	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)

	leaveRoutes := server.Group("/api/leaves")
	leaveRoutes.Use(middlewares.Authenticate(jwtService))
	{
		leaveRoutes.GET("", leaveController.GetAll)
		leaveRoutes.GET(":id", leaveController.GetByID)
		leaveRoutes.POST("", leaveController.Create)
		leaveRoutes.PUT(":id", leaveController.Update)
		leaveRoutes.DELETE(":id", leaveController.Delete)
	}
}
