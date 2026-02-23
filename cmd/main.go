package main

import (
	"log"
	"os"

	"github.com/Caknoooo/go-gin-clean-starter/middlewares"
	"github.com/Caknoooo/go-gin-clean-starter/modules/attendance"
	"github.com/Caknoooo/go-gin-clean-starter/modules/auth"
	"github.com/Caknoooo/go-gin-clean-starter/modules/employee"
	"github.com/Caknoooo/go-gin-clean-starter/modules/expenses"
	"github.com/Caknoooo/go-gin-clean-starter/modules/leave_types"
	"github.com/Caknoooo/go-gin-clean-starter/modules/master"
	"github.com/Caknoooo/go-gin-clean-starter/modules/notifications"
	"github.com/Caknoooo/go-gin-clean-starter/modules/payroll"
	"github.com/Caknoooo/go-gin-clean-starter/modules/rbac"
	rbacRepository "github.com/Caknoooo/go-gin-clean-starter/modules/rbac/repository"
	rbacService "github.com/Caknoooo/go-gin-clean-starter/modules/rbac/service"
	"github.com/Caknoooo/go-gin-clean-starter/modules/user"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/constants"
	"github.com/Caknoooo/go-gin-clean-starter/providers"
	"github.com/Caknoooo/go-gin-clean-starter/script"
	"github.com/samber/do"
	"gorm.io/gorm"

	"github.com/common-nighthawk/go-figure"
	"github.com/gin-gonic/gin"
)

func args(injector *do.Injector) bool {
	if len(os.Args) > 1 {
		flag := script.Commands(injector)
		return flag
	}

	return true
}

func run(server *gin.Engine) {
	server.Static("/assets", "./assets")

	port := os.Getenv("GOLANG_PORT")
	if port == "" {
		port = "8888"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "0.0.0.0:" + port
	} else {
		serve = ":" + port
	}

	myFigure := figure.NewColorFigure("Caknoo", "", "green", true)
	myFigure.Print()

	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}

func main() {
	var (
		injector = do.New()
	)

	providers.RegisterDependencies(injector)

	do.Provide(injector, func(i *do.Injector) (rbacRepository.RbacRepository, error) {
		db := do.MustInvokeNamed[*gorm.DB](i, constants.DB)
		return rbacRepository.NewRbacRepository(db), nil
	})
	do.Provide(injector, func(i *do.Injector) (rbacService.RbacService, error) {
		repo := do.MustInvoke[rbacRepository.RbacRepository](i)
		db := do.MustInvokeNamed[*gorm.DB](i, constants.DB)
		return rbacService.NewRbacService(repo, db), nil
	})

	if !args(injector) {
		return
	}

	server := gin.Default()
	server.Use(middlewares.CORSMiddleware())

	// Register module routes
	user.RegisterRoutes(server, injector)
	auth.RegisterRoutes(server, injector)
	employee.RegisterRoutes(server, injector)
	master.RegisterRoutes(server, injector)
	leave_types.RegisterRoutes(server, injector)
	expenses.RegisterRoutes(server, injector)
	payroll.RegisterRoutes(server, injector)
	notifications.RegisterRoutes(server, injector)
	attendance.RegisterRoutes(server, injector)
	rbac.RegisterRoutes(server, injector)

	run(server)
}
