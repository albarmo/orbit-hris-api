package service

import (
	"context"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/Caknoooo/go-gin-clean-starter/modules/expenses/dto"
	"github.com/Caknoooo/go-gin-clean-starter/modules/expenses/repository"
	"github.com/google/uuid"
)

type ExpenseService interface {
    // categories
    CreateCategory(ctx context.Context, req dto.ExpenseCategoryCreateRequest) (dto.ExpenseCategoryResponse, error)
    FindAllCategories(ctx context.Context) ([]dto.ExpenseCategoryResponse, error)
    FindCategoryByID(ctx context.Context, id uuid.UUID) (dto.ExpenseCategoryResponse, error)
    UpdateCategory(ctx context.Context, id uuid.UUID, req dto.ExpenseCategoryCreateRequest) (dto.ExpenseCategoryResponse, error)
    DeleteCategory(ctx context.Context, id uuid.UUID) error

    // expenses
    CreateExpense(ctx context.Context, req dto.ExpenseCreateRequest) (dto.ExpenseResponse, error)
    FindAllExpenses(ctx context.Context) ([]dto.ExpenseResponse, error)
    FindExpenseByID(ctx context.Context, id uuid.UUID) (dto.ExpenseResponse, error)
    UpdateExpense(ctx context.Context, id uuid.UUID, req dto.ExpenseUpdateRequest) (dto.ExpenseResponse, error)
    DeleteExpense(ctx context.Context, id uuid.UUID) error
}

type expenseService struct{
    repo repository.ExpenseRepository
}

func NewExpenseService(r repository.ExpenseRepository) ExpenseService {
    return &expenseService{repo: r}
}

func (s *expenseService) CreateCategory(ctx context.Context, req dto.ExpenseCategoryCreateRequest) (dto.ExpenseCategoryResponse, error) {
    c := &entities.ExpenseCategory{Name: req.Name}
    if err := s.repo.CreateCategory(ctx, c); err != nil {
        return dto.ExpenseCategoryResponse{}, err
    }
    return dto.ExpenseCategoryResponse{ID: c.ID.String(), Name: c.Name}, nil
}

func (s *expenseService) FindAllCategories(ctx context.Context) ([]dto.ExpenseCategoryResponse, error) {
    list, err := s.repo.FindAllCategories(ctx)
    if err != nil { return nil, err }
    var out []dto.ExpenseCategoryResponse
    for _, c := range list { out = append(out, dto.ExpenseCategoryResponse{ID: c.ID.String(), Name: c.Name}) }
    return out, nil
}

func (s *expenseService) FindCategoryByID(ctx context.Context, id uuid.UUID) (dto.ExpenseCategoryResponse, error) {
    c, err := s.repo.FindCategoryByID(ctx, id)
    if err != nil { return dto.ExpenseCategoryResponse{}, err }
    return dto.ExpenseCategoryResponse{ID: c.ID.String(), Name: c.Name}, nil
}

func (s *expenseService) UpdateCategory(ctx context.Context, id uuid.UUID, req dto.ExpenseCategoryCreateRequest) (dto.ExpenseCategoryResponse, error) {
    c, err := s.repo.FindCategoryByID(ctx, id)
    if err != nil { return dto.ExpenseCategoryResponse{}, err }
    c.Name = req.Name
    if err := s.repo.UpdateCategory(ctx, c); err != nil { return dto.ExpenseCategoryResponse{}, err }
    return s.FindCategoryByID(ctx, id)
}

func (s *expenseService) DeleteCategory(ctx context.Context, id uuid.UUID) error {
    return s.repo.DeleteCategory(ctx, id)
}

func (s *expenseService) CreateExpense(ctx context.Context, req dto.ExpenseCreateRequest) (dto.ExpenseResponse, error) {
    e := &entities.Expense{EmployeeID: uuid.MustParse(req.EmployeeID), CategoryID: uuid.MustParse(req.CategoryID), Amount: req.Amount, Description: req.Description, ReceiptURL: req.ReceiptURL}
    if err := s.repo.CreateExpense(ctx, e); err != nil { return dto.ExpenseResponse{}, err }
    return dto.ExpenseResponse{ID: e.ID.String(), EmployeeID: e.EmployeeID.String(), CategoryID: e.CategoryID.String(), Amount: e.Amount, Description: e.Description, ReceiptURL: e.ReceiptURL, Status: e.Status}, nil
}

func (s *expenseService) FindAllExpenses(ctx context.Context) ([]dto.ExpenseResponse, error) {
    list, err := s.repo.FindAllExpenses(ctx)
    if err != nil { return nil, err }
    var out []dto.ExpenseResponse
    for _, e := range list { out = append(out, dto.ExpenseResponse{ID: e.ID.String(), EmployeeID: e.EmployeeID.String(), CategoryID: e.CategoryID.String(), Amount: e.Amount, Description: e.Description, ReceiptURL: e.ReceiptURL, Status: e.Status}) }
    return out, nil
}

func (s *expenseService) FindExpenseByID(ctx context.Context, id uuid.UUID) (dto.ExpenseResponse, error) {
    e, err := s.repo.FindExpenseByID(ctx, id)
    if err != nil { return dto.ExpenseResponse{}, err }
    return dto.ExpenseResponse{ID: e.ID.String(), EmployeeID: e.EmployeeID.String(), CategoryID: e.CategoryID.String(), Amount: e.Amount, Description: e.Description, ReceiptURL: e.ReceiptURL, Status: e.Status}, nil
}

func (s *expenseService) UpdateExpense(ctx context.Context, id uuid.UUID, req dto.ExpenseUpdateRequest) (dto.ExpenseResponse, error) {
    e, err := s.repo.FindExpenseByID(ctx, id)
    if err != nil { return dto.ExpenseResponse{}, err }
    if req.CategoryID != nil { e.CategoryID = uuid.MustParse(*req.CategoryID) }
    if req.Amount != nil { e.Amount = *req.Amount }
    if req.Description != nil { e.Description = *req.Description }
    if req.ReceiptURL != nil { e.ReceiptURL = *req.ReceiptURL }
    if req.Status != nil { e.Status = *req.Status }
    if err := s.repo.UpdateExpense(ctx, e); err != nil { return dto.ExpenseResponse{}, err }
    return s.FindExpenseByID(ctx, id)
}

func (s *expenseService) DeleteExpense(ctx context.Context, id uuid.UUID) error {
    return s.repo.DeleteExpense(ctx, id)
}
