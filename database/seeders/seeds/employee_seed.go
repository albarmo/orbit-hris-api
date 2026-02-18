package seeds

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"gorm.io/gorm"
)

func EmployeeSeeder(db *gorm.DB) error {
	if err := seedEmployees(db); err != nil {
		return err
	}
	if err := seedEmployeePersonalInfos(db); err != nil {
		return err
	}
	if err := seedEmployeeAddresses(db); err != nil {
		return err
	}
	if err := seedEmployeeLegalInfos(db); err != nil {
		return err
	}
	if err := seedEmployeePayrollProfiles(db); err != nil {
		return err
	}
	return nil
}

func seedEmployees(db *gorm.DB) error {
	jsonFile, err := os.Open("./database/seeders/json/employees.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var listData []entities.Employee
	if err := json.Unmarshal(jsonData, &listData); err != nil {
		return err
	}

	for _, data := range listData {
		var existingData entities.Employee
		err := db.Where("employee_code = ?", data.EmployeeCode).First(&existingData).Error
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

func seedEmployeePersonalInfos(db *gorm.DB) error {
	jsonFile, err := os.Open("./database/seeders/json/employee_personal_infos.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var listData []entities.EmployeePersonalInfo
	if err := json.Unmarshal(jsonData, &listData); err != nil {
		return err
	}

	for _, data := range listData {
		var existingData entities.EmployeePersonalInfo
		err := db.Where("employee_id = ?", data.EmployeeID).First(&existingData).Error
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

func seedEmployeeAddresses(db *gorm.DB) error {
	jsonFile, err := os.Open("./database/seeders/json/employee_addresses.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var listData []entities.EmployeeAddress
	if err := json.Unmarshal(jsonData, &listData); err != nil {
		return err
	}

	for _, data := range listData {
		var existingData entities.EmployeeAddress
		// Using ID for 1-to-many relationship to ensure idempotency
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

func seedEmployeeLegalInfos(db *gorm.DB) error {
	jsonFile, err := os.Open("./database/seeders/json/employee_legal_infos.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var listData []entities.EmployeeLegalInfo
	if err := json.Unmarshal(jsonData, &listData); err != nil {
		return err
	}

	for _, data := range listData {
		var existingData entities.EmployeeLegalInfo
		err := db.Where("employee_id = ?", data.EmployeeID).First(&existingData).Error
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

func seedEmployeePayrollProfiles(db *gorm.DB) error {
	jsonFile, err := os.Open("./database/seeders/json/employee_payroll_profiles.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var listData []entities.EmployeePayrollProfile
	if err := json.Unmarshal(jsonData, &listData); err != nil {
		return err
	}

	for _, data := range listData {
		var existingData entities.EmployeePayrollProfile
		err := db.Where("employee_id = ?", data.EmployeeID).First(&existingData).Error
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
