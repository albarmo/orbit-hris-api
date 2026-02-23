package entities

import (
	"time"

	"github.com/google/uuid"
)

type ExpenseCategory struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name        string    `gorm:"type:varchar;not null" json:"name"`
	LimitAmount float64   `gorm:"type:numeric(15,2)" json:"limit_amount"`
}

func (ExpenseCategory) TableName() string {
	return "expense_categories"
}

type Expense struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	EmployeeID  uuid.UUID  `gorm:"type:uuid;not null;index" json:"employee_id"`
	CategoryID  uuid.UUID  `gorm:"type:uuid;not null;index" json:"category_id"`
	Amount      float64    `gorm:"type:numeric(15,2)" json:"amount"`
	Description string     `gorm:"type:text" json:"description"`
	ReceiptURL  string     `gorm:"type:varchar" json:"receipt_url"`
	Status      string     `gorm:"type:varchar" json:"status"`
	SubmittedAt time.Time  `gorm:"type:timestamptz" json:"submitted_at"`
	ApprovedBy  *uuid.UUID `gorm:"type:uuid" json:"approved_by"`
	ApprovedAt  *time.Time `gorm:"type:timestamptz" json:"approved_at"`

	Employee Employee        `gorm:"foreignKey:EmployeeID;references:ID" json:"-"`
	Category ExpenseCategory `gorm:"foreignKey:CategoryID;references:ID" json:"-"`
}

func (Expense) TableName() string {
	return "expenses"
}
