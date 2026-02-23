package migrations

import (
	"github.com/Caknoooo/go-gin-clean-starter/database"
	"gorm.io/gorm"
)

func init() {
	database.RegisterMigration("20240101000001_create_refresh_tokens_table", Up20240101000001CreateRefreshTokensTable, Down20240101000001CreateRefreshTokensTable)
}

func Up20240101000001CreateRefreshTokensTable(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`CREATE TABLE IF NOT EXISTS refresh_tokens (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			token varchar(255) NOT NULL UNIQUE,
			expires_at timestamptz NOT NULL,
			created_at timestamptz DEFAULT now()
		);`).Error; err != nil {
			return err
		}
		return nil
	})
}

func Down20240101000001CreateRefreshTokensTable(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`DROP TABLE IF EXISTS refresh_tokens CASCADE;`).Error; err != nil {
			return err
		}
		return nil
	})
}
