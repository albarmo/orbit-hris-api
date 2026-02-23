package master

import (
	"github.com/Caknoooo/go-gin-clean-starter/middlewares"
	authservice "github.com/Caknoooo/go-gin-clean-starter/modules/auth/service"
	mastercontroller "github.com/Caknoooo/go-gin-clean-starter/modules/master/controller"
	shiftscontroller "github.com/Caknoooo/go-gin-clean-starter/modules/master/controller"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/constants"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	masterController := do.MustInvoke[mastercontroller.MasterController](injector)
	shiftController := do.MustInvoke[shiftscontroller.ShiftController](injector)

	jwtService := do.MustInvokeNamed[authservice.JWTService](injector, constants.JWTService)

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

		// Shifts (moved from modules/shifts)
		masterRoutes.POST("/shifts", shiftController.Create)
		masterRoutes.GET("/shifts", shiftController.GetAll)
		masterRoutes.GET("/shifts/:id", shiftController.GetByID)
		masterRoutes.PUT("/shifts/:id", shiftController.Update)
		masterRoutes.DELETE("/shifts/:id", shiftController.Delete)

		// Employee shifts
		masterRoutes.POST("/shifts/assign", shiftController.CreateEmployeeShift)
		masterRoutes.GET("/shifts/employee/:employee_id", shiftController.GetEmployeeShifts)
		masterRoutes.DELETE("/shifts/employee/:employee_id/:id", shiftController.DeleteEmployeeShift)
	}
}
