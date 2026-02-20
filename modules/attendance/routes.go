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
		attendanceRoutes.GET("", middlewares.Authenticate(jwtService), attendanceController.GetAll)
		attendanceRoutes.GET("/:id", middlewares.Authenticate(jwtService), attendanceController.GetByID)
		attendanceRoutes.POST("/check-in", middlewares.Authenticate(jwtService), attendanceController.CheckIn)
		attendanceRoutes.PUT("/check-out", middlewares.Authenticate(jwtService), attendanceController.CheckOut)
		attendanceRoutes.PUT("/:id", middlewares.Authenticate(jwtService), attendanceController.Update)
		attendanceRoutes.PUT("/:id/approve", middlewares.Authenticate(jwtService), attendanceController.Approve)
		attendanceRoutes.DELETE("/:id", middlewares.Authenticate(jwtService), attendanceController.Delete)
		attendanceRoutes.GET("/employee/:employee_id", middlewares.Authenticate(jwtService), attendanceController.GetByEmployeeID)
	}
}
