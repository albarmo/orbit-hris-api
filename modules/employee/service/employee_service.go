package service

import (
	"context"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/Caknoooo/go-gin-clean-starter/modules/employee/dto"
	"github.com/Caknoooo/go-gin-clean-starter/modules/employee/repository"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/pagination"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EmployeeService interface {
	Create(ctx context.Context, req dto.EmployeeCreateRequest) (dto.EmployeeResponse, error)
	FindAll(ctx context.Context, filter *pagination.Filter) (*pagination.Page[dto.EmployeeResponse], error)
	FindByID(ctx context.Context, id uuid.UUID) (dto.EmployeeResponse, error)
	Update(ctx context.Context, id uuid.UUID, req dto.EmployeeUpdateRequest) (dto.EmployeeResponse, error)
	Delete(ctx context.Context, id uuid.UUID) error

	// Child table service methods
	CreatePersonalInfo(ctx context.Context, employeeID uuid.UUID, req dto.EmployeePersonalInfoCreateRequest) (dto.EmployeePersonalInfoCreateRequest, error)
	GetPersonalInfo(ctx context.Context, employeeID uuid.UUID) (dto.EmployeePersonalInfoCreateRequest, error)
	UpdatePersonalInfo(ctx context.Context, employeeID uuid.UUID, req dto.EmployeePersonalInfoUpdateRequest) (dto.EmployeePersonalInfoUpdateRequest, error)
	DeletePersonalInfo(ctx context.Context, employeeID uuid.UUID) error

	CreateAddress(ctx context.Context, employeeID uuid.UUID, req dto.EmployeeAddressCreateRequest) (dto.EmployeeAddressCreateRequest, error)
	GetAddresses(ctx context.Context, employeeID uuid.UUID) ([]dto.EmployeeAddressCreateRequest, error)
	GetAddress(ctx context.Context, id uuid.UUID) (dto.EmployeeAddressCreateRequest, error)
	UpdateAddress(ctx context.Context, id uuid.UUID, req dto.EmployeeAddressUpdateRequest) (dto.EmployeeAddressUpdateRequest, error)
	DeleteAddress(ctx context.Context, id uuid.UUID) error

	CreateLegalInfo(ctx context.Context, employeeID uuid.UUID, req dto.EmployeeLegalInfoCreateRequest) (dto.EmployeeLegalInfoCreateRequest, error)
	GetLegalInfo(ctx context.Context, employeeID uuid.UUID) (dto.EmployeeLegalInfoCreateRequest, error)
	UpdateLegalInfo(ctx context.Context, employeeID uuid.UUID, req dto.EmployeeLegalInfoUpdateRequest) (dto.EmployeeLegalInfoUpdateRequest, error)
	DeleteLegalInfo(ctx context.Context, employeeID uuid.UUID) error

	CreatePayrollProfile(ctx context.Context, employeeID uuid.UUID, req dto.EmployeePayrollProfileCreateRequest) (dto.EmployeePayrollProfileCreateRequest, error)
	GetPayrollProfile(ctx context.Context, employeeID uuid.UUID) (dto.EmployeePayrollProfileCreateRequest, error)
	UpdatePayrollProfile(ctx context.Context, employeeID uuid.UUID, req dto.EmployeePayrollProfileUpdateRequest) (dto.EmployeePayrollProfileUpdateRequest, error)
	DeletePayrollProfile(ctx context.Context, employeeID uuid.UUID) error
}

type employeeService struct {
	employeeRepository repository.EmployeeRepository
	db                 *gorm.DB
}

func NewEmployeeService(employeeRepository repository.EmployeeRepository,
	db *gorm.DB,
	 ) EmployeeService {
	return &employeeService{
		employeeRepository: employeeRepository,
		db:                 db,
	}
}

func (s *employeeService) Create(ctx context.Context, req dto.EmployeeCreateRequest) (dto.EmployeeResponse, error) {
	tx := s.db.Begin()
	defer tx.Rollback()

	employee := entities.Employee{
		UserID:           req.UserID,
		EmployeeCode:     req.EmployeeCode,
		SupervisorID:     req.SupervisorID,
		DepartmentID:     req.DepartmentID,
		PositionID:       req.PositionID,
		JoinDate:         req.JoinDate,
		EndDate:          req.EndDate,
		EmploymentType:   req.EmploymentType,
		EmploymentStatus: req.EmploymentStatus,
		ProbationEndDate: req.ProbationEndDate,
	}

	personalInfo := entities.EmployeePersonalInfo{
		NIK:           req.PersonalInfo.NIK,
		Gender:        req.PersonalInfo.Gender,
		BirthPlace:    req.PersonalInfo.BirthPlace,
		BirthDate:     req.PersonalInfo.BirthDate,
		MaritalStatus: req.PersonalInfo.MaritalStatus,
		Religion:      req.PersonalInfo.Religion,
		Nationality:   req.PersonalInfo.Nationality,
		PersonalEmail: req.PersonalInfo.PersonalEmail,
		PersonalPhone: req.PersonalInfo.PersonalPhone,
	}

	var addresses []entities.EmployeeAddress
	for _, addr := range req.Addresses {
		addresses = append(addresses, entities.EmployeeAddress{
			Type:       addr.Type,
			Address:    addr.Address,
			City:       addr.City,
			Province:   addr.Province,
			PostalCode: addr.PostalCode,
		})
	}

	legalInfo := entities.EmployeeLegalInfo{
		NPWP:                req.LegalInfo.NPWP,
		BPJSKesehatan:       req.LegalInfo.BPJSKesehatan,
		BPJSKetenagakerjaan: req.LegalInfo.BPJSKetenagakerjaan,
		PassportNumber:      req.LegalInfo.PassportNumber,
		PassportExpiredDate: req.LegalInfo.PassportExpiredDate,
	}

	payrollInfo := entities.EmployeePayrollProfile{
		BasicSalary:       req.PayrollInfo.BasicSalary,
		BankName:          req.PayrollInfo.BankName,
		BankAccountNumber: req.PayrollInfo.BankAccountNumber,
		BankAccountHolder: req.PayrollInfo.BankAccountHolder,
	}

	createdEmployee, err := s.employeeRepository.Create(ctx, tx, employee, personalInfo, addresses, legalInfo, payrollInfo)
	if err != nil {
		return dto.EmployeeResponse{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return dto.EmployeeResponse{}, err
	}

	return s.FindByID(ctx, createdEmployee.ID)
}

func (s *employeeService) FindAll(ctx context.Context, filter *pagination.Filter) (*pagination.Page[dto.EmployeeResponse], error) {
	page, err := s.employeeRepository.FindAll(ctx, s.db, filter)
	if err != nil {
		return nil, err
	}

	var employeeResponses []dto.EmployeeResponse
	for _, employee := range page.Data {
		employeeResponses = append(employeeResponses, dto.EmployeeResponse{
			ID:               employee.ID,
			UserID:           employee.UserID,
			EmployeeCode:     employee.EmployeeCode,
			SupervisorID:     employee.SupervisorID,
			DepartmentID:     employee.DepartmentID,
			PositionID:       employee.PositionID,
			JoinDate:         employee.JoinDate,
			EndDate:          employee.EndDate,
			EmploymentType:   employee.EmploymentType,
			EmploymentStatus: employee.EmploymentStatus,
			ProbationEndDate: employee.ProbationEndDate,
			User: dto.UserResponse{
				ID:         employee.User.ID,
				Name:       employee.User.Name,
				Email:      employee.User.Email,
				TelpNumber: employee.User.TelpNumber,
				Role:       employee.User.Role,
				ImageUrl:   employee.User.ImageUrl,
				IsVerified: employee.User.IsVerified,
			},
			Department: struct {
				ID   uuid.UUID `json:"id"`
				Name string    `json:"name"`
			}{
				ID:   employee.Department.ID,
				Name: employee.Department.Name,
			},
			Position: struct {
				ID   uuid.UUID `json:"id"`
				Name string    `json:"name"`
			}{
				ID:   employee.Position.ID,
				Name: employee.Position.Name,
			},
		})
	}

	return &pagination.Page[dto.EmployeeResponse]{
		Page:       page.Page,
		Limit:      page.Limit,
		Total:      page.Total,
		TotalPages: page.TotalPages,
		Data:       employeeResponses,
	}, nil
}

func (s *employeeService) FindByID(ctx context.Context, id uuid.UUID) (dto.EmployeeResponse, error) {
	employee, err := s.employeeRepository.FindByID(ctx, s.db, id)
	if err != nil {
		return dto.EmployeeResponse{}, err
	}

	return dto.EmployeeResponse{
		ID:               employee.ID,
		UserID:           employee.UserID,
		EmployeeCode:     employee.EmployeeCode,
		SupervisorID:     employee.SupervisorID,
		DepartmentID:     employee.DepartmentID,
		PositionID:       employee.PositionID,
		JoinDate:         employee.JoinDate,
		EndDate:          employee.EndDate,
		EmploymentType:   employee.EmploymentType,
		EmploymentStatus: employee.EmploymentStatus,
		ProbationEndDate: employee.ProbationEndDate,
		User: dto.UserResponse{
			ID:         employee.User.ID,
			Name:       employee.User.Name,
			Email:      employee.User.Email,
			TelpNumber: employee.User.TelpNumber,
			Role:       employee.User.Role,
			ImageUrl:   employee.User.ImageUrl,
			IsVerified: employee.User.IsVerified,
		},
		Department: struct {
			ID   uuid.UUID `json:"id"`
			Name string    `json:"name"`
		}{
			ID:   employee.Department.ID,
			Name: employee.Department.Name,
		},
		Position: struct {
			ID   uuid.UUID `json:"id"`
			Name string    `json:"name"`
		}{
			ID:   employee.Position.ID,
			Name: employee.Position.Name,
		},
	}, nil
}

// Child table service implementations
func (s *employeeService) CreatePersonalInfo(ctx context.Context, employeeID uuid.UUID, req dto.EmployeePersonalInfoCreateRequest) (dto.EmployeePersonalInfoCreateRequest, error) {
	tx := s.db.Begin()
	defer tx.Rollback()

	info := entities.EmployeePersonalInfo{
		EmployeeID:    employeeID,
		NIK:           req.NIK,
		Gender:        req.Gender,
		BirthPlace:    req.BirthPlace,
		BirthDate:     req.BirthDate,
		MaritalStatus: req.MaritalStatus,
		Religion:      req.Religion,
		Nationality:   req.Nationality,
		PersonalEmail: req.PersonalEmail,
		PersonalPhone: req.PersonalPhone,
	}

	if _, err := s.employeeRepository.CreatePersonalInfo(ctx, tx, info); err != nil {
		return dto.EmployeePersonalInfoCreateRequest{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return dto.EmployeePersonalInfoCreateRequest{}, err
	}

	return req, nil
}

func (s *employeeService) GetPersonalInfo(ctx context.Context, employeeID uuid.UUID) (dto.EmployeePersonalInfoCreateRequest, error) {
	info, err := s.employeeRepository.GetPersonalInfoByEmployeeID(ctx, s.db, employeeID)
	if err != nil {
		return dto.EmployeePersonalInfoCreateRequest{}, err
	}
	return dto.EmployeePersonalInfoCreateRequest{
		NIK:           info.NIK,
		Gender:        info.Gender,
		BirthPlace:    info.BirthPlace,
		BirthDate:     info.BirthDate,
		MaritalStatus: info.MaritalStatus,
		Religion:      info.Religion,
		Nationality:   info.Nationality,
		PersonalEmail: info.PersonalEmail,
		PersonalPhone: info.PersonalPhone,
	}, nil
}

func (s *employeeService) UpdatePersonalInfo(ctx context.Context, employeeID uuid.UUID, req dto.EmployeePersonalInfoUpdateRequest) (dto.EmployeePersonalInfoUpdateRequest, error) {
	tx := s.db.Begin()
	defer tx.Rollback()

	info := entities.EmployeePersonalInfo{
		EmployeeID:    employeeID,
		MaritalStatus: req.MaritalStatus,
		PersonalEmail: req.PersonalEmail,
		PersonalPhone: req.PersonalPhone,
	}

	if _, err := s.employeeRepository.UpdatePersonalInfo(ctx, tx, info); err != nil {
		return dto.EmployeePersonalInfoUpdateRequest{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return dto.EmployeePersonalInfoUpdateRequest{}, err
	}

	return req, nil
}

func (s *employeeService) DeletePersonalInfo(ctx context.Context, employeeID uuid.UUID) error {
	tx := s.db.Begin()
	defer tx.Rollback()

	if err := s.employeeRepository.DeletePersonalInfo(ctx, tx, employeeID); err != nil {
		return err
	}

	return tx.Commit().Error
}

// Addresses
func (s *employeeService) CreateAddress(ctx context.Context, employeeID uuid.UUID, req dto.EmployeeAddressCreateRequest) (dto.EmployeeAddressCreateRequest, error) {
	tx := s.db.Begin()
	defer tx.Rollback()

	addr := entities.EmployeeAddress{
		EmployeeID: employeeID,
		Type:       req.Type,
		Address:    req.Address,
		City:       req.City,
		Province:   req.Province,
		PostalCode: req.PostalCode,
	}

	if _, err := s.employeeRepository.CreateAddress(ctx, tx, addr); err != nil {
		return dto.EmployeeAddressCreateRequest{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return dto.EmployeeAddressCreateRequest{}, err
	}

	return req, nil
}

func (s *employeeService) GetAddresses(ctx context.Context, employeeID uuid.UUID) ([]dto.EmployeeAddressCreateRequest, error) {
	addrs, err := s.employeeRepository.GetAddressesByEmployeeID(ctx, s.db, employeeID)
	if err != nil {
		return nil, err
	}
	var res []dto.EmployeeAddressCreateRequest
	for _, a := range addrs {
		res = append(res, dto.EmployeeAddressCreateRequest{
			Type:       a.Type,
			Address:    a.Address,
			City:       a.City,
			Province:   a.Province,
			PostalCode: a.PostalCode,
		})
	}
	return res, nil
}

func (s *employeeService) GetAddress(ctx context.Context, id uuid.UUID) (dto.EmployeeAddressCreateRequest, error) {
	a, err := s.employeeRepository.GetAddressByID(ctx, s.db, id)
	if err != nil {
		return dto.EmployeeAddressCreateRequest{}, err
	}
	return dto.EmployeeAddressCreateRequest{
		Type:       a.Type,
		Address:    a.Address,
		City:       a.City,
		Province:   a.Province,
		PostalCode: a.PostalCode,
	}, nil
}

func (s *employeeService) UpdateAddress(ctx context.Context, id uuid.UUID, req dto.EmployeeAddressUpdateRequest) (dto.EmployeeAddressUpdateRequest, error) {
	tx := s.db.Begin()
	defer tx.Rollback()

	addr := entities.EmployeeAddress{
		ID:         uuid.Nil,
		Type:       req.Type,
		Address:    req.Address,
		City:       req.City,
		Province:   req.Province,
		PostalCode: req.PostalCode,
	}
	// set ID
	if req.ID != nil {
		addr.ID = *req.ID
	} else {
		addr.ID = id
	}

	if _, err := s.employeeRepository.UpdateAddress(ctx, tx, addr); err != nil {
		return dto.EmployeeAddressUpdateRequest{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return dto.EmployeeAddressUpdateRequest{}, err
	}

	return req, nil
}

func (s *employeeService) DeleteAddress(ctx context.Context, id uuid.UUID) error {
	tx := s.db.Begin()
	defer tx.Rollback()

	if err := s.employeeRepository.DeleteAddress(ctx, tx, id); err != nil {
		return err
	}
	return tx.Commit().Error
}

// LegalInfo
func (s *employeeService) CreateLegalInfo(ctx context.Context, employeeID uuid.UUID, req dto.EmployeeLegalInfoCreateRequest) (dto.EmployeeLegalInfoCreateRequest, error) {
	tx := s.db.Begin()
	defer tx.Rollback()

	info := entities.EmployeeLegalInfo{
		EmployeeID:          employeeID,
		NPWP:                req.NPWP,
		BPJSKesehatan:       req.BPJSKesehatan,
		BPJSKetenagakerjaan: req.BPJSKetenagakerjaan,
		PassportNumber:      req.PassportNumber,
		PassportExpiredDate: req.PassportExpiredDate,
	}

	if _, err := s.employeeRepository.CreateLegalInfo(ctx, tx, info); err != nil {
		return dto.EmployeeLegalInfoCreateRequest{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return dto.EmployeeLegalInfoCreateRequest{}, err
	}

	return req, nil
}

func (s *employeeService) GetLegalInfo(ctx context.Context, employeeID uuid.UUID) (dto.EmployeeLegalInfoCreateRequest, error) {
	info, err := s.employeeRepository.GetLegalInfoByEmployeeID(ctx, s.db, employeeID)
	if err != nil {
		return dto.EmployeeLegalInfoCreateRequest{}, err
	}
	return dto.EmployeeLegalInfoCreateRequest{
		NPWP:                info.NPWP,
		BPJSKesehatan:       info.BPJSKesehatan,
		BPJSKetenagakerjaan: info.BPJSKetenagakerjaan,
		PassportNumber:      info.PassportNumber,
		PassportExpiredDate: info.PassportExpiredDate,
	}, nil
}

func (s *employeeService) UpdateLegalInfo(ctx context.Context, employeeID uuid.UUID, req dto.EmployeeLegalInfoUpdateRequest) (dto.EmployeeLegalInfoUpdateRequest, error) {
	tx := s.db.Begin()
	defer tx.Rollback()

	info := entities.EmployeeLegalInfo{
		EmployeeID:          employeeID,
		NPWP:                req.NPWP,
		BPJSKesehatan:       req.BPJSKesehatan,
		BPJSKetenagakerjaan: req.BPJSKetenagakerjaan,
		PassportNumber:      req.PassportNumber,
		PassportExpiredDate: req.PassportExpiredDate,
	}

	if _, err := s.employeeRepository.UpdateLegalInfo(ctx, tx, info); err != nil {
		return dto.EmployeeLegalInfoUpdateRequest{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return dto.EmployeeLegalInfoUpdateRequest{}, err
	}

	return req, nil
}

func (s *employeeService) DeleteLegalInfo(ctx context.Context, employeeID uuid.UUID) error {
	tx := s.db.Begin()
	defer tx.Rollback()

	if err := s.employeeRepository.DeleteLegalInfo(ctx, tx, employeeID); err != nil {
		return err
	}
	return tx.Commit().Error
}

// Payroll
func (s *employeeService) CreatePayrollProfile(ctx context.Context, employeeID uuid.UUID, req dto.EmployeePayrollProfileCreateRequest) (dto.EmployeePayrollProfileCreateRequest, error) {
	tx := s.db.Begin()
	defer tx.Rollback()

	p := entities.EmployeePayrollProfile{
		EmployeeID:        employeeID,
		BasicSalary:       req.BasicSalary,
		BankName:          req.BankName,
		BankAccountNumber: req.BankAccountNumber,
		BankAccountHolder: req.BankAccountHolder,
	}

	if _, err := s.employeeRepository.CreatePayrollProfile(ctx, tx, p); err != nil {
		return dto.EmployeePayrollProfileCreateRequest{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return dto.EmployeePayrollProfileCreateRequest{}, err
	}

	return req, nil
}

func (s *employeeService) GetPayrollProfile(ctx context.Context, employeeID uuid.UUID) (dto.EmployeePayrollProfileCreateRequest, error) {
	p, err := s.employeeRepository.GetPayrollByEmployeeID(ctx, s.db, employeeID)
	if err != nil {
		return dto.EmployeePayrollProfileCreateRequest{}, err
	}
	return dto.EmployeePayrollProfileCreateRequest{
		BasicSalary:       p.BasicSalary,
		BankName:          p.BankName,
		BankAccountNumber: p.BankAccountNumber,
		BankAccountHolder: p.BankAccountHolder,
	}, nil
}

func (s *employeeService) UpdatePayrollProfile(ctx context.Context, employeeID uuid.UUID, req dto.EmployeePayrollProfileUpdateRequest) (dto.EmployeePayrollProfileUpdateRequest, error) {
	tx := s.db.Begin()
	defer tx.Rollback()

	p := entities.EmployeePayrollProfile{
		EmployeeID:        employeeID,
		BasicSalary:       req.BasicSalary,
		BankName:          req.BankName,
		BankAccountNumber: req.BankAccountNumber,
		BankAccountHolder: req.BankAccountHolder,
	}

	if _, err := s.employeeRepository.UpdatePayrollProfile(ctx, tx, p); err != nil {
		return dto.EmployeePayrollProfileUpdateRequest{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return dto.EmployeePayrollProfileUpdateRequest{}, err
	}

	return req, nil
}

func (s *employeeService) DeletePayrollProfile(ctx context.Context, employeeID uuid.UUID) error {
	tx := s.db.Begin()
	defer tx.Rollback()

	if err := s.employeeRepository.DeletePayrollProfile(ctx, tx, employeeID); err != nil {
		return err
	}
	return tx.Commit().Error
}

func (s *employeeService) Update(ctx context.Context, id uuid.UUID, req dto.EmployeeUpdateRequest) (dto.EmployeeResponse, error) {
	tx := s.db.Begin()
	defer tx.Rollback()

	employee, err := s.employeeRepository.FindByID(ctx, tx, id)
	if err != nil {
		return dto.EmployeeResponse{}, err
	}

	if req.SupervisorID != nil {
		employee.SupervisorID = req.SupervisorID
	}
	if req.DepartmentID != uuid.Nil {
		employee.DepartmentID = req.DepartmentID
	}
	if req.PositionID != uuid.Nil {
		employee.PositionID = req.PositionID
	}
	if !req.EndDate.IsZero() {
		employee.EndDate = req.EndDate
	}
	if req.EmploymentType != "" {
		employee.EmploymentType = req.EmploymentType
	}
	if req.EmploymentStatus != "" {
		employee.EmploymentStatus = req.EmploymentStatus
	}
	if !req.ProbationEndDate.IsZero() {
		employee.ProbationEndDate = req.ProbationEndDate
	}

	updatedEmployee, err := s.employeeRepository.Update(ctx, tx, employee)
	if err != nil {
		return dto.EmployeeResponse{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return dto.EmployeeResponse{}, err
	}

	return s.FindByID(ctx, updatedEmployee.ID)
}

func (s *employeeService) Delete(ctx context.Context, id uuid.UUID) error {
	tx := s.db.Begin()
	defer tx.Rollback()

	if err := s.employeeRepository.Delete(ctx, tx, id); err != nil {
		return err
	}

	return tx.Commit().Error
}
