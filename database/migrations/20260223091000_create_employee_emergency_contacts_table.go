package migrations

import (
	"github.com/Caknoooo/go-gin-clean-starter/database"
	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"gorm.io/gorm"
)

func init() {
	database.RegisterMigration("20260223091000_create_employee_emergency_contacts_table", UpCreateEmployeeEmergencyContactsTable, DownCreateEmployeeEmergencyContactsTable)
}

func UpCreateEmployeeEmergencyContactsTable(db *gorm.DB) error {
	return db.AutoMigrate(&entities.EmployeeEmergencyContact{})
}

func DownCreateEmployeeEmergencyContactsTable(db *gorm.DB) error {
	return db.Migrator().DropTable(&entities.EmployeeEmergencyContact{})
}
