package migrations

import (
	"github.com/Caknoooo/go-gin-clean-starter/database"
	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"gorm.io/gorm"
)

func init() {
    database.RegisterMigration("20260223092000_create_leave_types_and_leave_balances_tables", UpCreateLeaveTypesAndLeaveBalancesTables, DownCreateLeaveTypesAndLeaveBalancesTables)
}

func UpCreateLeaveTypesAndLeaveBalancesTables(db *gorm.DB) error {
    if err := db.AutoMigrate(&entities.LeaveType{}); err != nil {
        return err
    }
    return db.AutoMigrate(&entities.LeaveBalance{})
}

func DownCreateLeaveTypesAndLeaveBalancesTables(db *gorm.DB) error {
    if err := db.Migrator().DropTable(&entities.LeaveBalance{}); err != nil {
        return err
    }
    return db.Migrator().DropTable(&entities.LeaveType{})
}
