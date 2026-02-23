package migrations

import (
	"github.com/Caknoooo/go-gin-clean-starter/database"
	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"gorm.io/gorm"
)

func init() {
    database.RegisterMigration("20260223092500_create_expense_categories_and_expenses_tables", UpCreateExpenseCategoriesAndExpensesTables, DownCreateExpenseCategoriesAndExpensesTables)
}

func UpCreateExpenseCategoriesAndExpensesTables(db *gorm.DB) error {
    if err := db.AutoMigrate(&entities.ExpenseCategory{}); err != nil {
        return err
    }
    return db.AutoMigrate(&entities.Expense{})
}

func DownCreateExpenseCategoriesAndExpensesTables(db *gorm.DB) error {
    if err := db.Migrator().DropTable(&entities.Expense{}); err != nil {
        return err
    }
    return db.Migrator().DropTable(&entities.ExpenseCategory{})
}
