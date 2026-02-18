package database

import (
	"fmt"
	"github.com/Caknoooo/go-gin-clean-starter/database/seeders/seeds"
	"gorm.io/gorm"
)

type SeederFunc func(db *gorm.DB) error

func Seeder(db *gorm.DB) error {
	seeders := []SeederFunc{
		seeds.ListUserSeeder,
		seeds.RoleSeeder,
		seeds.PermissionSeeder,
		seeds.RolePermissionSeeder,
		seeds.DepartmentSeeder,
		seeds.PositionSeeder,
		seeds.LocationSeeder,
		seeds.EmployeeSeeder,
		seeds.AttendanceSeeder,
	}

	for _, seeder := range seeders {
		if err := seeder(db); err != nil {
			// Instead of returning the error, we can log it and continue
			// This makes the seeder process more robust
			// For this example, we will return the error to be strict
			// To get the function name, we can use runtime, but it adds complexity
			fmt.Printf("Error running seeder: %v\n", err)
			return err
		}
	}

	return nil
}
