package seeds

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"gorm.io/gorm"
)

func AttendanceSeeder(db *gorm.DB) error {
	jsonFile, err := os.Open("./database/seeders/json/attendances.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var listData []entities.Attendance
	if err := json.Unmarshal(jsonData, &listData); err != nil {
		return err
	}

	for _, data := range listData {
		var existingData entities.Attendance
		// Use ID for idempotency check as attendance records might not have other unique keys
		err := db.First(&existingData, "id = ?", data.ID).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
