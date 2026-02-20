package repository

import (
	"gorm.io/gorm"
)

type RbacRepository interface {
}

type rbacRepository struct {
	db *gorm.DB
}

func NewRbacRepository(db *gorm.DB) RbacRepository {
	return &rbacRepository{
		db: db,
	}
}
