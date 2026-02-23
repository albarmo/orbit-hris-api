package migrations

import (
	"github.com/Caknoooo/go-gin-clean-starter/database"
	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"gorm.io/gorm"
)

func init() {
	database.RegisterMigration("20260223090500_create_password_resets_table", UpCreatePasswordResetsTable, DownCreatePasswordResetsTable)
}

func UpCreatePasswordResetsTable(db *gorm.DB) error {
	return db.AutoMigrate(&entities.PasswordReset{})
}

func DownCreatePasswordResetsTable(db *gorm.DB) error {
	return db.Migrator().DropTable(&entities.PasswordReset{})
}
