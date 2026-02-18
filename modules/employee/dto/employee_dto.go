package dto

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Failed
	MESSAGE_FAILED_GET_DATA_FROM_BODY = "failed get data from body"
	MESSAGE_FAILED_CREATE_EMPLOYEE    = "failed create employee"
	MESSAGE_FAILED_GET_LIST_EMPLOYEE  = "failed get list employee"
	MESSAGE_FAILED_GET_EMPLOYEE       = "failed get employee"
	MESSAGE_FAILED_UPDATE_EMPLOYEE    = "failed update employee"
	MESSAGE_FAILED_DELETE_EMPLOYEE    = "failed delete employee"

	// Success
	MESSAGE_SUCCESS_CREATE_EMPLOYEE = "success create employee"
	MESSAGE_SUCCESS_GET_LIST_EMPLOYEE = "success get list employee"
	MESSAGE_SUCCESS_GET_EMPLOYEE    = "success get employee"
	MESSAGE_SUCCESS_UPDATE_EMPLOYEE = "success update employee"
	MESSAGE_SUCCESS_DELETE_EMPLOYEE = "success delete employee"
)

type (
	EmployeeCreateRequest struct {
		UserID           uuid.UUID  `json:"user_id" binding:"required"`
		EmployeeCode     string     `json:"employee_code" binding:"required"`
		SupervisorID     *uuid.UUID `json:"supervisor_id"`
		DepartmentID     uuid.UUID  `json:"department_id" binding:"required"`
		PositionID       uuid.UUID  `json:"position_id" binding:"required"`
		JoinDate         time.Time  `json:"join_date" binding:"required"`
		EndDate          time.Time  `json:"end_date"`
		EmploymentType   string     `json:"employment_type" binding:"required"`
		EmploymentStatus string     `json:"employment_status" binding:"required"`
		ProbationEndDate time.Time  `json:"probation_end_date"`

		PersonalInfo EmployeePersonalInfoCreateRequest `json:"personal_info" binding:"required"`
		Addresses    []EmployeeAddressCreateRequest    `json:"addresses" binding:"required"`
		LegalInfo    EmployeeLegalInfoCreateRequest    `json:"legal_info" binding:"required"`
		PayrollInfo  EmployeePayrollProfileCreateRequest `json:"payroll_info" binding:"required"`
	}

	EmployeePersonalInfoCreateRequest struct {
		NIK           string    `json:"nik" binding:"required"`
		Gender        string    `json:"gender" binding:"required"`
		BirthPlace    string    `json:"birth_place" binding:"required"`
		BirthDate     time.Time `json:"birth_date" binding:"required"`
		MaritalStatus string    `json:"marital_status" binding:"required"`
		Religion      string    `json:"religion" binding:"required"`
		Nationality   string    `json:"nationality" binding:"required"`
		PersonalEmail string    `json:"personal_email" binding:"required,email"`
		PersonalPhone string    `json:"personal_phone" binding:"required"`
	}

	EmployeeAddressCreateRequest struct {
		Type       string `json:"type" binding:"required"`
		Address    string `json:"address" binding:"required"`
		City       string `json:"city" binding:"required"`
		Province   string `json:"province" binding:"required"`
		PostalCode string `json:"postal_code" binding:"required"`
	}

	EmployeeLegalInfoCreateRequest struct {
		NPWP                string    `json:"npwp" binding:"required"`
		BPJSKesehatan       string    `json:"bpjs_kesehatan" binding:"required"`
		BPJSKetenagakerjaan string    `json:"bpjs_ketenagakerjaan" binding:"required"`
		PassportNumber      string    `json:"passport_number"`
		PassportExpiredDate time.Time `json:"passport_expired_date"`
	}

	EmployeePayrollProfileCreateRequest struct {
		BasicSalary       float64 `json:"basic_salary" binding:"required"`
		BankName          string  `json:"bank_name" binding:"required"`
		BankAccountNumber string  `json:"bank_account_number" binding:"required"`
		BankAccountHolder string  `json:"bank_account_holder" binding:"required"`
	}

	EmployeeUpdateRequest struct {
		SupervisorID     *uuid.UUID `json:"supervisor_id"`
		DepartmentID     uuid.UUID  `json:"department_id"`
		PositionID       uuid.UUID  `json:"position_id"`
		EndDate          time.Time  `json:"end_date"`
		EmploymentType   string     `json:"employment_type"`
		EmploymentStatus string     `json:"employment_status"`
		ProbationEndDate time.Time  `json:"probation_end_date"`

		PersonalInfo EmployeePersonalInfoUpdateRequest `json:"personal_info"`
		Addresses    []EmployeeAddressUpdateRequest    `json:"addresses"`
		LegalInfo    EmployeeLegalInfoUpdateRequest    `json:"legal_info"`
		PayrollInfo  EmployeePayrollProfileUpdateRequest `json:"payroll_info"`
	}

	EmployeePersonalInfoUpdateRequest struct {
		MaritalStatus string `json:"marital_status"`
		PersonalEmail string `json:"personal_email" binding:"omitempty,email"`
		PersonalPhone string `json:"personal_phone"`
	}

	EmployeeAddressUpdateRequest struct {
		ID         *uuid.UUID `json:"id"`
		Type       string     `json:"type"`
		Address    string     `json:"address"`
		City       string     `json:"city"`
		Province   string     `json:"province"`
		PostalCode string     `json:"postal_code"`
	}

	EmployeeLegalInfoUpdateRequest struct {
		NPWP                string    `json:"npwp"`
		BPJSKesehatan       string    `json:"bpjs_kesehatan"`
		BPJSKetenagakerjaan string    `json:"bpjs_ketenagakerjaan"`
		PassportNumber      string    `json:"passport_number"`
		PassportExpiredDate time.Time `json:"passport_expired_date"`
	}

	EmployeePayrollProfileUpdateRequest struct {
		BasicSalary       float64 `json:"basic_salary"`
		BankName          string  `json:"bank_name"`
		BankAccountNumber string  `json:"bank_account_number"`
		BankAccountHolder string  `json:"bank_account_holder"`
	}

	UserResponse struct {
		ID         uuid.UUID `json:"id"`
		Name       string    `json:"name"`
		Email      string    `json:"email"`
		TelpNumber string    `json:"telp_number"`
		Role       string    `json:"role"`
		ImageUrl   string    `json:"image_url"`
		IsVerified bool      `json:"is_verified"`
	}

	EmployeeResponse struct {
		ID               uuid.UUID `json:"id"`
		UserID           uuid.UUID `json:"user_id"`
		EmployeeCode     string    `json:"employee_code"`
		SupervisorID     *uuid.UUID `json:"supervisor_id"`
		DepartmentID     uuid.UUID `json:"department_id"`
		PositionID       uuid.UUID `json:"position_id"`
		JoinDate         time.Time `json:"join_date"`
		EndDate          time.Time `json:"end_date"`
		EmploymentType   string    `json:"employment_type"`
		EmploymentStatus string    `json:"employment_status"`
		ProbationEndDate time.Time `json:"probation_end_date"`

		User       UserResponse `json:"user"`
		Department struct {
			ID   uuid.UUID `json:"id"`
			Name string    `json:"name"`
		} `json:"department"`
		Position struct {
			ID   uuid.UUID `json:"id"`
			Name string    `json:"name"`
		} `json:"position"`
	}
)
