package migrations

import (
	"github.com/Caknoooo/go-gin-clean-starter/database"
	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"gorm.io/gorm"
)

func init() {
	database.RegisterMigration("20260223093000_create_payroll_items_table", UpCreatePayrollItemsTable, DownCreatePayrollItemsTable)
}

func UpCreatePayrollItemsTable(db *gorm.DB) error {
	return db.AutoMigrate(&entities.PayrollItem{})
}

func DownCreatePayrollItemsTable(db *gorm.DB) error {
	return db.Migrator().DropTable(&entities.PayrollItem{})
}
