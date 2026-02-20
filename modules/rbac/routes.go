package rbac

import (
	"github.com/Caknoooo/go-gin-clean-starter/modules/rbac/controller"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	rbacController := do.MustInvoke[controller.RbacController](injector)

	rbacRoutes := server.Group("/api/rbac")
	{
		// TODO: add your endpoints here
	}
}
