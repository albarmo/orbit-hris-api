package employee

import (
	"github.com/Caknoooo/go-gin-clean-starter/middlewares"
	"github.com/Caknoooo/go-gin-clean-starter/modules/auth/service"
	"github.com/Caknoooo/go-gin-clean-starter/modules/employee/controller"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/constants"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	employeeController := do.MustInvoke[controller.EmployeeController](injector)
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)

	employeeRoutes := server.Group("/api/employee")
	{
		employeeRoutes.POST("", middlewares.Authenticate(jwtService), employeeController.Create)
		employeeRoutes.GET("", middlewares.Authenticate(jwtService), employeeController.GetAllEmployee)
		employeeRoutes.GET("/:id", middlewares.Authenticate(jwtService), employeeController.GetEmployeeByID)
		employeeRoutes.PUT("/:id", middlewares.Authenticate(jwtService), employeeController.Update)
		employeeRoutes.DELETE("/:id", middlewares.Authenticate(jwtService), employeeController.Delete)

		// Child tables
		employeeRoutes.POST("/:id/personal_info", middlewares.Authenticate(jwtService), employeeController.CreatePersonalInfo)
		employeeRoutes.GET("/:id/personal_info", middlewares.Authenticate(jwtService), employeeController.GetPersonalInfo)
		employeeRoutes.PUT("/:id/personal_info", middlewares.Authenticate(jwtService), employeeController.UpdatePersonalInfo)
		employeeRoutes.DELETE("/:id/personal_info", middlewares.Authenticate(jwtService), employeeController.DeletePersonalInfo)

		employeeRoutes.POST("/:id/addresses", middlewares.Authenticate(jwtService), employeeController.CreateAddress)
		employeeRoutes.GET("/:id/addresses", middlewares.Authenticate(jwtService), employeeController.GetAddresses)
		employeeRoutes.GET("/:id/addresses/:address_id", middlewares.Authenticate(jwtService), employeeController.GetAddressByID)
		employeeRoutes.PUT("/:id/addresses/:address_id", middlewares.Authenticate(jwtService), employeeController.UpdateAddress)
		employeeRoutes.DELETE("/:id/addresses/:address_id", middlewares.Authenticate(jwtService), employeeController.DeleteAddress)

		employeeRoutes.POST("/:id/legal_info", middlewares.Authenticate(jwtService), employeeController.CreateLegalInfo)
		employeeRoutes.GET("/:id/legal_info", middlewares.Authenticate(jwtService), employeeController.GetLegalInfo)
		employeeRoutes.PUT("/:id/legal_info", middlewares.Authenticate(jwtService), employeeController.UpdateLegalInfo)
		employeeRoutes.DELETE("/:id/legal_info", middlewares.Authenticate(jwtService), employeeController.DeleteLegalInfo)

		employeeRoutes.POST("/:id/payroll_profile", middlewares.Authenticate(jwtService), employeeController.CreatePayrollProfile)
		employeeRoutes.GET("/:id/payroll_profile", middlewares.Authenticate(jwtService), employeeController.GetPayrollProfile)
		employeeRoutes.PUT("/:id/payroll_profile", middlewares.Authenticate(jwtService), employeeController.UpdatePayrollProfile)
		employeeRoutes.DELETE("/:id/payroll_profile", middlewares.Authenticate(jwtService), employeeController.DeletePayrollProfile)
	}
}
