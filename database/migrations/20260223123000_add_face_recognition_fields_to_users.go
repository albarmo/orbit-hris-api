package migrations

import (
	"github.com/Caknoooo/go-gin-clean-starter/database"
	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"gorm.io/gorm"
)

func init() {
	database.RegisterMigration("20260223123000_add_face_recognition_fields_to_users", Up20260223123000AddFaceRecognitionFieldsToUsers, Down20260223123000AddFaceRecognitionFieldsToUsers)
}

func Up20260223123000AddFaceRecognitionFieldsToUsers(db *gorm.DB) error {
	// AutoMigrate will add the new columns on User struct
	return db.AutoMigrate(&entities.User{})
}

func Down20260223123000AddFaceRecognitionFieldsToUsers(db *gorm.DB) error {
	// Drop the added columns
	if err := db.Migrator().DropColumn(&entities.User{}, "FaceRecognitionModelID"); err != nil {
		return err
	}
	if err := db.Migrator().DropColumn(&entities.User{}, "FaceRecognitionUserID"); err != nil {
		return err
	}
	if err := db.Migrator().DropColumn(&entities.User{}, "FaceRecognitionBestPhotoPath"); err != nil {
		return err
	}
	if err := db.Migrator().DropColumn(&entities.User{}, "FaceRecognitionBestPhotoID"); err != nil {
		return err
	}
	if err := db.Migrator().DropColumn(&entities.User{}, "FaceRecognitionScore"); err != nil {
		return err
	}
	return nil
}
