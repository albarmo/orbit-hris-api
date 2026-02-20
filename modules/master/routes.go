package master

import (
	"github.com/Caknoooo/go-gin-clean-starter/middlewares"
	"github.com/Caknoooo/go-gin-clean-starter/modules/auth/service"
	"github.com/Caknoooo/go-gin-clean-starter/modules/master/controller"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/constants"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	masterController := do.MustInvoke[controller.MasterController](injector)

	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)

	masterRoutes := server.Group("/api/master")
	masterRoutes.Use(middlewares.Authenticate(jwtService))
	{
		// Departments
		masterRoutes.GET("/departments", masterController.GetDepartments)
		masterRoutes.GET("/departments/:id", masterController.GetDepartmentByID)
		masterRoutes.POST("/departments", masterController.CreateDepartment)
		masterRoutes.PUT("/departments/:id", masterController.UpdateDepartment)
		masterRoutes.DELETE("/departments/:id", masterController.DeleteDepartment)

		// Locations
		masterRoutes.GET("/locations", masterController.GetLocations)
		masterRoutes.GET("/locations/:id", masterController.GetLocationByID)
		masterRoutes.POST("/locations", masterController.CreateLocation)
		masterRoutes.PUT("/locations/:id", masterController.UpdateLocation)
		masterRoutes.DELETE("/locations/:id", masterController.DeleteLocation)

		// Positions
		masterRoutes.GET("/positions", masterController.GetPositions)
		masterRoutes.GET("/positions/:id", masterController.GetPositionByID)
		masterRoutes.POST("/positions", masterController.CreatePosition)
		masterRoutes.PUT("/positions/:id", masterController.UpdatePosition)
		masterRoutes.DELETE("/positions/:id", masterController.DeletePosition)
	}
}
