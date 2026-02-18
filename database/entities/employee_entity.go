package entities

import (
	"time"

	"github.com/google/uuid"
)

type Employee struct {
	ID               uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID           uuid.UUID `gorm:"type:uuid;unique" json:"user_id"`
	EmployeeCode     string    `gorm:"type:varchar;unique;not null" json:"employee_code"`
	SupervisorID     *uuid.UUID `gorm:"type:uuid" json:"supervisor_id"`
	DepartmentID     uuid.UUID `gorm:"type:uuid" json:"department_id"`
	PositionID       uuid.UUID `gorm:"type:uuid" json:"position_id"`
	JoinDate         time.Time `gorm:"type:date" json:"join_date"`
	EndDate          time.Time `gorm:"type:date" json:"end_date"`
	EmploymentType   string    `gorm:"type:varchar" json:"employment_type"`
	EmploymentStatus string    `gorm:"type:varchar" json:"employment_status"`
	ProbationEndDate time.Time `gorm:"type:date" json:"probation_end_date"`

	User       User       `gorm:"foreignKey:UserID;references:ID" json:"user"`
	Supervisor *Employee  `gorm:"foreignKey:SupervisorID;references:ID" json:"supervisor"`
	Department Department `gorm:"foreignKey:DepartmentID;references:ID" json:"department"`
	Position   Position   `gorm:"foreignKey:PositionID;references:ID" json:"position"`

	Timestamp
}

type EmployeePersonalInfo struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	EmployeeID     uuid.UUID `gorm:"type:uuid;unique" json:"employee_id"`
	NIK            string    `gorm:"type:varchar;unique" json:"nik"`
	Gender         string    `gorm:"type:varchar" json:"gender"`
	BirthPlace     string    `gorm:"type:varchar" json:"birth_place"`
	BirthDate      time.Time `gorm:"type:date" json:"birth_date"`
	MaritalStatus  string    `gorm:"type:varchar" json:"marital_status"`
	Religion       string    `gorm:"type:varchar" json:"religion"`
	Nationality    string    `gorm:"type:varchar" json:"nationality"`
	PersonalEmail  string    `gorm:"type:varchar" json:"personal_email"`
	PersonalPhone  string    `gorm:"type:varchar" json:"personal_phone"`

	Timestamp
}

func (EmployeePersonalInfo) TableName() string {
	return "employee_personal_infos"
}

type EmployeeAddress struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	EmployeeID uuid.UUID `gorm:"type:uuid" json:"employee_id"`
	Type       string    `gorm:"type:varchar" json:"type"`
	Address    string    `gorm:"type:text" json:"address"`
	City       string    `gorm:"type:varchar" json:"city"`
	Province   string    `gorm:"type:varchar" json:"province"`
	PostalCode string    `gorm:"type:varchar" json:"postal_code"`
	CreatedAt  time.Time `gorm:"type:timestamp with time zone;default:now()" json:"created_at"`
}

type EmployeeLegalInfo struct {
	ID                  uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	EmployeeID          uuid.UUID `gorm:"type:uuid;unique" json:"employee_id"`
	NPWP                string    `gorm:"type:varchar" json:"npwp"`
	BPJSKesehatan       string    `gorm:"type:varchar" json:"bpjs_kesehatan"`
	BPJSKetenagakerjaan string    `gorm:"type:varchar" json:"bpjs_ketenagakerjaan"`
	PassportNumber      string    `gorm:"type:varchar" json:"passport_number"`
	PassportExpiredDate time.Time `gorm:"type:date" json:"passport_expired_date"`

	Timestamp
}

func (EmployeeLegalInfo) TableName() string {
	return "employee_legal_infos"
}

type EmployeePayrollProfile struct {
	ID                  uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	EmployeeID          uuid.UUID `gorm:"type:uuid;unique" json:"employee_id"`
	BasicSalary         float64   `gorm:"type:numeric(15,2)" json:"basic_salary"`
	BankName            string    `gorm:"type:varchar" json:"bank_name"`
	BankAccountNumber   string    `gorm:"type:varchar" json:"bank_account_number"`
	BankAccountHolder   string    `gorm:"type:varchar" json:"bank_account_holder"`

	Timestamp
}

func (EmployeePayrollProfile) TableName() string {
	return "employee_payroll_profiles"
}
