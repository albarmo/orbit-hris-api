package providers

import (
	"github.com/Caknoooo/go-gin-clean-starter/config"
	attendanceController "github.com/Caknoooo/go-gin-clean-starter/modules/attendance/controller"
	attendanceRepository "github.com/Caknoooo/go-gin-clean-starter/modules/attendance/repository"
	attendanceService "github.com/Caknoooo/go-gin-clean-starter/modules/attendance/service"
	authController "github.com/Caknoooo/go-gin-clean-starter/modules/auth/controller"
	authRepo "github.com/Caknoooo/go-gin-clean-starter/modules/auth/repository"
	authService "github.com/Caknoooo/go-gin-clean-starter/modules/auth/service"
	employeeController "github.com/Caknoooo/go-gin-clean-starter/modules/employee/controller"
	employeeRepository "github.com/Caknoooo/go-gin-clean-starter/modules/employee/repository"
	employeeService "github.com/Caknoooo/go-gin-clean-starter/modules/employee/service"
	masterController "github.com/Caknoooo/go-gin-clean-starter/modules/master/controller"
	masterRepository "github.com/Caknoooo/go-gin-clean-starter/modules/master/repository"
	masterService "github.com/Caknoooo/go-gin-clean-starter/modules/master/service"
	rbacController "github.com/Caknoooo/go-gin-clean-starter/modules/rbac/controller"
	rbacRepositoryPkg "github.com/Caknoooo/go-gin-clean-starter/modules/rbac/repository"
	rbacService "github.com/Caknoooo/go-gin-clean-starter/modules/rbac/service"
	userController "github.com/Caknoooo/go-gin-clean-starter/modules/user/controller"
	"github.com/Caknoooo/go-gin-clean-starter/modules/user/repository"
	userService "github.com/Caknoooo/go-gin-clean-starter/modules/user/service"

	"github.com/Caknoooo/go-gin-clean-starter/pkg/constants"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func InitDatabase(injector *do.Injector) {
	do.ProvideNamed(injector, constants.DB, func(i *do.Injector) (*gorm.DB, error) {
		return config.SetUpDatabaseConnection(), nil
	})
}

func RegisterDependencies(injector *do.Injector) {
	InitDatabase(injector)

	do.ProvideNamed(injector, constants.JWTService, func(i *do.Injector) (authService.JWTService, error) {
		return authService.NewJWTService(), nil
	})

	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	jwtService := do.MustInvokeNamed[authService.JWTService](injector, constants.JWTService)

	userRepository := repository.NewUserRepository(db)
	refreshTokenRepository := authRepo.NewRefreshTokenRepository(db)
	employeeRepository := employeeRepository.NewEmployeeRepository(db)
	attendanceRepository := attendanceRepository.NewAttendanceRepository(db)
	masterRepository := masterRepository.NewMasterRepository(db)

	rbacRepository := rbacRepositoryPkg.NewRbacRepository(db)

	userService := userService.NewUserService(userRepository, db)
	authService := authService.NewAuthService(userRepository, refreshTokenRepository, jwtService, db)
	employeeService := employeeService.NewEmployeeService(employeeRepository, db)
	attendanceService := attendanceService.NewAttendanceService(attendanceRepository, db)
	masterService := masterService.NewMasterService(masterRepository, db)
	rbacService := rbacService.NewRbacService(rbacRepository, db)

	do.Provide(
		injector, func(i *do.Injector) (userController.UserController, error) {
			return userController.NewUserController(i, userService), nil
		},
	)

	do.Provide(
		injector, func(i *do.Injector) (authController.AuthController, error) {
			return authController.NewAuthController(i, authService), nil
		},
	)

	do.Provide(
		injector, func(i *do.Injector) (employeeController.EmployeeController, error) {
			return employeeController.NewEmployeeController(employeeService), nil
		},
	)

	do.Provide(
		injector, func(i *do.Injector) (attendanceController.AttendanceController, error) {
			return attendanceController.NewAttendanceController(i, attendanceService), nil
		},
	)

	do.Provide(
		injector, func(i *do.Injector) (masterController.MasterController, error) {
			return masterController.NewMasterController(masterService), nil
		},
	)

	do.Provide(
		injector, func(i *do.Injector) (rbacController.RbacController, error) {
			return rbacController.NewRbacController(i, rbacService), nil
		},
	)
}
