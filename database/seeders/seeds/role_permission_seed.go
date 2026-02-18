package seeds

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"gorm.io/gorm"
)

type RolePermissionData struct {
	RoleName       string `json:"role_name"`
	PermissionName string `json:"permission_name"`
}

func RolePermissionSeeder(db *gorm.DB) error {
	jsonFile, err := os.Open("./database/seeders/json/role_permissions.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var listData []RolePermissionData
	if err := json.Unmarshal(jsonData, &listData); err != nil {
		return err
	}

	for _, data := range listData {
		var role entities.Role
		if err := db.Where("name = ?", data.RoleName).First(&role).Error; err != nil {
			// If role not found, we skip this entry, maybe log it
			fmt.Printf("Warning: Role '%s' not found, skipping permission assignment.\n", data.RoleName)
			continue
		}

		var permission entities.Permission
		if err := db.Where("name = ?", data.PermissionName).First(&permission).Error; err != nil {
			// If permission not found, we skip this entry, maybe log it
			fmt.Printf("Warning: Permission '%s' not found for role '%s', skipping.\n", data.PermissionName, data.RoleName)
			continue
		}

		rolePermission := entities.RolePermission{
			RoleID:       role.ID,
			PermissionID: permission.ID,
		}

		// Check if the association already exists
		var existing entities.RolePermission
		err := db.Where(&rolePermission).First(&existing).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		// If it doesn't exist, create it
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := db.Create(&rolePermission).Error; err != nil {
				return err
			}
		}
	}

	return nil
}