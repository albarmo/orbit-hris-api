package migrations

import (
	"github.com/Caknoooo/go-gin-clean-starter/database"
	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"gorm.io/gorm"
)

func init() {
    database.RegisterMigration("20260223091500_create_shifts_and_employee_shifts_tables", UpCreateShiftsAndEmployeeShiftsTables, DownCreateShiftsAndEmployeeShiftsTables)
}

func UpCreateShiftsAndEmployeeShiftsTables(db *gorm.DB) error {
    if err := db.AutoMigrate(&entities.Shift{}); err != nil {
        return err
    }
    return db.AutoMigrate(&entities.EmployeeShift{})
}

func DownCreateShiftsAndEmployeeShiftsTables(db *gorm.DB) error {
    if err := db.Migrator().DropTable(&entities.EmployeeShift{}); err != nil {
        return err
    }
    return db.Migrator().DropTable(&entities.Shift{})
}
