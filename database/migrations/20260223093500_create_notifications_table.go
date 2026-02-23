package migrations

import (
	"github.com/Caknoooo/go-gin-clean-starter/database"
	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"gorm.io/gorm"
)

func init() {
    database.RegisterMigration("20260223093500_create_notifications_table", UpCreateNotificationsTable, DownCreateNotificationsTable)
}

func UpCreateNotificationsTable(db *gorm.DB) error {
    return db.AutoMigrate(&entities.Notification{})
}

func DownCreateNotificationsTable(db *gorm.DB) error {
    return db.Migrator().DropTable(&entities.Notification{})
}
