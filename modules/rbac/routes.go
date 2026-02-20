package rbac

import (
	"github.com/Caknoooo/go-gin-clean-starter/middlewares"
	"github.com/Caknoooo/go-gin-clean-starter/modules/auth/service"
	"github.com/Caknoooo/go-gin-clean-starter/modules/rbac/controller"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/constants"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	rbacController := do.MustInvoke[controller.RbacController](injector)

	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)

	rbacRoutes := server.Group("/api/rbac")
	rbacRoutes.Use(middlewares.Authenticate(jwtService))
	{
		// Roles
		rbacRoutes.GET("/roles", rbacController.GetRoles)
		rbacRoutes.GET("/roles/:id", rbacController.GetRoleByID)
		rbacRoutes.POST("/roles", rbacController.CreateRole)
		rbacRoutes.PUT("/roles/:id", rbacController.UpdateRole)
		rbacRoutes.DELETE("/roles/:id", rbacController.DeleteRole)

		// Permissions
		rbacRoutes.GET("/permissions", rbacController.GetPermissions)
		rbacRoutes.GET("/permissions/:id", rbacController.GetPermissionByID)
		rbacRoutes.POST("/permissions", rbacController.CreatePermission)
		rbacRoutes.PUT("/permissions/:id", rbacController.UpdatePermission)
		rbacRoutes.DELETE("/permissions/:id", rbacController.DeletePermission)

		// User roles
		rbacRoutes.GET("/users/:user_id/roles", rbacController.GetRolesByUser)
		rbacRoutes.POST("/users/:user_id/roles", rbacController.AssignRole)
		rbacRoutes.DELETE("/users/:user_id/roles/:role_id", rbacController.RemoveRole)
	}
}
