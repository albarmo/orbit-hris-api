package migrations

import (
	"github.com/Caknoooo/go-gin-clean-starter/database"
	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"gorm.io/gorm"
)

func init() {
	database.RegisterMigration("20260223235247_create_leaves_table", UpCreateLeavesTable, DownCreateLeavesTable)
}

func UpCreateLeavesTable(db *gorm.DB) error {
	return db.AutoMigrate(&entities.Leave{})
}

func DownCreateLeavesTable(db *gorm.DB) error {
	return db.Migrator().DropTable(&entities.Leave{})
}
