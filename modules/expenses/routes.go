package expenses

import (
	"github.com/Caknoooo/go-gin-clean-starter/modules/expenses/controller"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	ctrl := do.MustInvoke[controller.ExpenseController](injector)

	// categories
	cat := server.Group("/api/expense-categories")
	{
		cat.POST("/", ctrl.CreateCategory)
		cat.GET("/", ctrl.GetCategories)
		cat.GET("/:id", ctrl.GetCategory)
		cat.PUT("/:id", ctrl.UpdateCategory)
		cat.DELETE("/:id", ctrl.DeleteCategory)
	}

	// expenses
	exp := server.Group("/api/expenses")
	{
		exp.POST("/", ctrl.CreateExpense)
		exp.GET("/", ctrl.GetExpenses)
		exp.GET("/:id", ctrl.GetExpense)
		exp.PUT("/:id", ctrl.UpdateExpense)
		exp.DELETE("/:id", ctrl.DeleteExpense)
	}
}
