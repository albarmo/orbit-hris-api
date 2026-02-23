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
	shiftsController "github.com/Caknoooo/go-gin-clean-starter/modules/shifts/controller"
	shiftsRepository "github.com/Caknoooo/go-gin-clean-starter/modules/shifts/repository"
	shiftsService "github.com/Caknoooo/go-gin-clean-starter/modules/shifts/service"
	userController "github.com/Caknoooo/go-gin-clean-starter/modules/user/controller"
	"github.com/Caknoooo/go-gin-clean-starter/modules/user/repository"
	userService "github.com/Caknoooo/go-gin-clean-starter/modules/user/service"

	expensesController "github.com/Caknoooo/go-gin-clean-starter/modules/expenses/controller"
	expensesRepository "github.com/Caknoooo/go-gin-clean-starter/modules/expenses/repository"
	expensesService "github.com/Caknoooo/go-gin-clean-starter/modules/expenses/service"

	payrollItemsController "github.com/Caknoooo/go-gin-clean-starter/modules/payroll_items/controller"
	payrollItemsRepository "github.com/Caknoooo/go-gin-clean-starter/modules/payroll_items/repository"
	payrollItemsService "github.com/Caknoooo/go-gin-clean-starter/modules/payroll_items/service"

	notificationsController "github.com/Caknoooo/go-gin-clean-starter/modules/notifications/controller"
	notificationsRepository "github.com/Caknoooo/go-gin-clean-starter/modules/notifications/repository"
	notificationsService "github.com/Caknoooo/go-gin-clean-starter/modules/notifications/service"

	leaveTypesController "github.com/Caknoooo/go-gin-clean-starter/modules/leave_types/controller"
	leaveTypesRepository "github.com/Caknoooo/go-gin-clean-starter/modules/leave_types/repository"
	leaveTypesService "github.com/Caknoooo/go-gin-clean-starter/modules/leave_types/service"

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
	ProvideFirebase(injector)

	do.ProvideNamed(injector, constants.JWTService, func(i *do.Injector) (authService.JWTService, error) {
		return authService.NewJWTService(), nil
	})

	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	jwtService := do.MustInvokeNamed[authService.JWTService](injector, constants.JWTService)

	userRepo := repository.NewUserRepository(db)
	refreshTokenRepo := authRepo.NewRefreshTokenRepository(db)
	employeeRepo := employeeRepository.NewEmployeeRepository(db)
	attendanceRepo := attendanceRepository.NewAttendanceRepository(db)
	masterRepo := masterRepository.NewMasterRepository(db)
	shiftRepo := shiftsRepository.NewShiftRepository(db)

	emergencyContactRepo := employeeRepository.NewEmergencyContactRepository(db)

	leaveTypeRepo := leaveTypesRepository.NewLeaveTypeRepository(db)
	leaveTypeSvc := leaveTypesService.NewLeaveTypeService(leaveTypeRepo)

	expenseRepo := expensesRepository.NewExpenseRepository(db)
	expenseSvc := expensesService.NewExpenseService(expenseRepo)

	payrollItemRepo := payrollItemsRepository.NewPayrollItemRepository(db)
	payrollItemSvc := payrollItemsService.NewPayrollItemService(payrollItemRepo)

	notificationRepo := notificationsRepository.NewNotificationRepository(db)
	notificationSvc := notificationsService.NewNotificationService(notificationRepo, injector)

	rbacRepository := rbacRepositoryPkg.NewRbacRepository(db)

	userSvc := userService.NewUserService(userRepo, db)
	authSvc := authService.NewAuthService(userRepo, refreshTokenRepo, jwtService, db)

	passwordResetRepository := authRepo.NewPasswordResetRepository(db)
	passwordResetService := authService.NewPasswordResetService(passwordResetRepository)
	employeeSvc := employeeService.NewEmployeeService(employeeRepo, db)

	emergencyContactService := employeeService.NewEmergencyContactService(emergencyContactRepo)

	shiftSvc := shiftsService.NewShiftService(shiftRepo)
	attendanceSvc := attendanceService.NewAttendanceService(attendanceRepo, db)
	masterSvc := masterService.NewMasterService(masterRepo, db)
	rbacSvc := rbacService.NewRbacService(rbacRepository, db)

	do.Provide(
		injector, func(i *do.Injector) (userController.UserController, error) {
			return userController.NewUserController(i, userSvc), nil
		},
	)

	do.Provide(
		injector, func(i *do.Injector) (authController.AuthController, error) {
			return authController.NewAuthController(i, authSvc), nil
		},
	)

	do.Provide(
		injector, func(i *do.Injector) (authController.PasswordResetController, error) {
			return authController.NewPasswordResetController(passwordResetService), nil
		},
	)

	do.Provide(
		injector, func(i *do.Injector) (employeeController.EmployeeController, error) {
			return employeeController.NewEmployeeController(employeeSvc), nil
		},
	)

	do.Provide(
		injector, func(i *do.Injector) (employeeController.EmergencyContactController, error) {
			return employeeController.NewEmergencyContactController(emergencyContactService), nil
		},
	)

	do.Provide(
		injector, func(i *do.Injector) (shiftsController.ShiftController, error) {
			return shiftsController.NewShiftController(shiftSvc), nil
		},
	)

	do.Provide(
		injector, func(i *do.Injector) (expensesController.ExpenseController, error) {
			return expensesController.NewExpenseController(expenseSvc), nil
		},
	)

	do.Provide(
		injector, func(i *do.Injector) (payrollItemsController.PayrollItemController, error) {
			return payrollItemsController.NewPayrollItemController(payrollItemSvc), nil
		},
	)

	do.Provide(
		injector, func(i *do.Injector) (notificationsController.NotificationController, error) {
			return notificationsController.NewNotificationController(notificationSvc), nil
		},
	)

	do.Provide(
		injector, func(i *do.Injector) (leaveTypesController.LeaveTypeController, error) {
			return leaveTypesController.NewLeaveTypeController(leaveTypeSvc), nil
		},
	)

	do.Provide(
		injector, func(i *do.Injector) (attendanceController.AttendanceController, error) {
			return attendanceController.NewAttendanceController(i, attendanceSvc), nil
		},
	)

	do.Provide(
		injector, func(i *do.Injector) (masterController.MasterController, error) {
			return masterController.NewMasterController(masterSvc), nil
		},
	)

	do.Provide(
		injector, func(i *do.Injector) (rbacController.RbacController, error) {
			return rbacController.NewRbacController(i, rbacSvc), nil
		},
	)
}
