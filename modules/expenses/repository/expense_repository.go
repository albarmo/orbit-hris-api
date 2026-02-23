package repository

import (
	"context"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ExpenseRepository interface {
	// categories
	CreateCategory(ctx context.Context, c *entities.ExpenseCategory) error
	FindAllCategories(ctx context.Context) ([]entities.ExpenseCategory, error)
	FindCategoryByID(ctx context.Context, id uuid.UUID) (*entities.ExpenseCategory, error)
	UpdateCategory(ctx context.Context, c *entities.ExpenseCategory) error
	DeleteCategory(ctx context.Context, id uuid.UUID) error

	// expenses
	CreateExpense(ctx context.Context, e *entities.Expense) error
	FindAllExpenses(ctx context.Context) ([]entities.Expense, error)
	FindExpenseByID(ctx context.Context, id uuid.UUID) (*entities.Expense, error)
	UpdateExpense(ctx context.Context, e *entities.Expense) error
	DeleteExpense(ctx context.Context, id uuid.UUID) error
}

type expenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) ExpenseRepository {
	return &expenseRepository{db: db}
}

func (r *expenseRepository) CreateCategory(ctx context.Context, c *entities.ExpenseCategory) error {
	return r.db.WithContext(ctx).Create(c).Error
}

func (r *expenseRepository) FindAllCategories(ctx context.Context) ([]entities.ExpenseCategory, error) {
	var list []entities.ExpenseCategory
	if err := r.db.WithContext(ctx).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *expenseRepository) FindCategoryByID(ctx context.Context, id uuid.UUID) (*entities.ExpenseCategory, error) {
	var c entities.ExpenseCategory
	if err := r.db.WithContext(ctx).First(&c, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *expenseRepository) UpdateCategory(ctx context.Context, c *entities.ExpenseCategory) error {
	return r.db.WithContext(ctx).Save(c).Error
}

func (r *expenseRepository) DeleteCategory(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&entities.ExpenseCategory{}, "id = ?", id).Error
}

func (r *expenseRepository) CreateExpense(ctx context.Context, e *entities.Expense) error {
	return r.db.WithContext(ctx).Create(e).Error
}

func (r *expenseRepository) FindAllExpenses(ctx context.Context) ([]entities.Expense, error) {
	var list []entities.Expense
	if err := r.db.WithContext(ctx).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *expenseRepository) FindExpenseByID(ctx context.Context, id uuid.UUID) (*entities.Expense, error) {
	var e entities.Expense
	if err := r.db.WithContext(ctx).First(&e, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *expenseRepository) UpdateExpense(ctx context.Context, e *entities.Expense) error {
	return r.db.WithContext(ctx).Save(e).Error
}

func (r *expenseRepository) DeleteExpense(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&entities.Expense{}, "id = ?", id).Error
}
