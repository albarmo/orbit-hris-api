package migrations

import (
	"github.com/Caknoooo/go-gin-clean-starter/database"
	"gorm.io/gorm"
)

func init() {
	database.RegisterMigration("20260223120000_drop_auth_and_refresh_tokens", UpDropAuthAndRefreshTokens, DownDropAuthAndRefreshTokens)
}

func UpDropAuthAndRefreshTokens(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`DROP TABLE IF EXISTS auth_tokens CASCADE;`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`DROP TABLE IF EXISTS refresh_tokens CASCADE;`).Error; err != nil {
			return err
		}
		return nil
	})
}

func DownDropAuthAndRefreshTokens(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`CREATE TABLE IF NOT EXISTS auth_tokens (
            id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
            user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
            access_uuid uuid UNIQUE,
            refresh_uuid uuid UNIQUE,
            expires_at timestamptz,
            created_at timestamptz DEFAULT now()
        );`).Error; err != nil {
			return err
		}

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
