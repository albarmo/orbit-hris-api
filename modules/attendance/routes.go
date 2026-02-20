package attendance

import (
	"github.com/Caknoooo/go-gin-clean-starter/middlewares"
	"github.com/Caknoooo/go-gin-clean-starter/modules/attendance/controller"
	"github.com/Caknoooo/go-gin-clean-starter/modules/auth/service"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/constants"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	attendanceController := do.MustInvoke[controller.AttendanceController](injector)
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)

	attendanceRoutes := server.Group("/api/attendances")
	attendanceRoutes.Use(middlewares.Authenticate(jwtService))
	{
		attendanceRoutes.GET("", attendanceController.GetAll)
		attendanceRoutes.GET("/:id", attendanceController.GetByID)
		attendanceRoutes.POST("/check-in", attendanceController.CheckIn)
		attendanceRoutes.PUT("/check-out", attendanceController.CheckOut)
		attendanceRoutes.PUT("/:id", attendanceController.Update)
		attendanceRoutes.PUT("/:id/approve", attendanceController.Approve)
		attendanceRoutes.DELETE("/:id", attendanceController.Delete)
	}
}
