package seeds

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"gorm.io/gorm"
)

func DepartmentSeeder(db *gorm.DB) error {
	jsonFile, err := os.Open("./database/seeders/json/departments.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var listData []entities.Department
	if err := json.Unmarshal(jsonData, &listData); err != nil {
		return err
	}

	hasTable := db.Migrator().HasTable(&entities.Department{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entities.Department{}); err != nil {
			return err
		}
	}

	for _, data := range listData {
		var existingData entities.Department
		err := db.Where("name = ?", data.Name).First(&existingData).Error

		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			// An actual error occurred
			return err
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Record not found, create it using the data from JSON (including the ID)
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
		// If err is nil, it means the record was found, so we do nothing.
	}

	return nil
}
