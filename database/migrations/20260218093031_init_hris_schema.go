package migrations

import (
	"github.com/Caknoooo/go-gin-clean-starter/database"
	"gorm.io/gorm"
)

func init() {
	database.RegisterMigration(
		"20260218093031_init_hris_schema",
		UpInitHrisSchema,
		DownInitHrisSchema,
	)
}

func UpInitHrisSchema(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {

		// =====================================================
		// EXTENSIONS
		// =====================================================
		if err := tx.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
			return err
		}

		// =====================================================
		// RBAC
		// =====================================================
		if err := tx.Exec(`
		CREATE TABLE roles (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			name varchar UNIQUE NOT NULL,
			description varchar
		);`).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
		CREATE TABLE permissions (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			name varchar UNIQUE NOT NULL,
			description varchar
		);`).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
		CREATE TABLE user_roles (
			user_id uuid REFERENCES users(id) ON DELETE CASCADE,
			role_id uuid REFERENCES roles(id) ON DELETE CASCADE,
			UNIQUE(user_id, role_id)
		);`).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
		CREATE TABLE role_permissions (
			role_id uuid REFERENCES roles(id) ON DELETE CASCADE,
			permission_id uuid REFERENCES permissions(id) ON DELETE CASCADE,
			UNIQUE(role_id, permission_id)
		);`).Error; err != nil {
			return err
		}

		// =====================================================
		// HR MASTER
		// =====================================================
		if err := tx.Exec(`
		CREATE TABLE departments (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			name varchar UNIQUE NOT NULL,
			description varchar
		);`).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
		CREATE TABLE positions (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			department_id uuid REFERENCES departments(id),
			name varchar NOT NULL,
			level varchar
		);`).Error; err != nil {
			return err
		}

		// =====================================================
		// EMPLOYEES
		// =====================================================
		if err := tx.Exec(`
		CREATE TABLE employees (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			user_id uuid UNIQUE REFERENCES users(id),
			employee_code varchar UNIQUE NOT NULL,
			supervisor_id uuid REFERENCES employees(id),
			department_id uuid REFERENCES departments(id),
			position_id uuid REFERENCES positions(id),
			join_date date,
			end_date date,
			employment_type varchar,
			employment_status varchar,
			probation_end_date date,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now()
		);`).Error; err != nil {
			return err
		}

		// =====================================================
		// EMPLOYEE DETAIL
		// =====================================================
		if err := tx.Exec(`
		CREATE TABLE employee_personal_infos (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			employee_id uuid UNIQUE REFERENCES employees(id),
			nik varchar UNIQUE,
			gender varchar,
			birth_place varchar,
			birth_date date,
			marital_status varchar,
			religion varchar,
			nationality varchar,
			personal_email varchar,
			personal_phone varchar,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now()
		);`).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
		CREATE TABLE employee_addresses (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			employee_id uuid REFERENCES employees(id),
			type varchar,
			address text,
			city varchar,
			province varchar,
			postal_code varchar,
			created_at timestamptz DEFAULT now()
		);`).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
		CREATE TABLE employee_legal_infos (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			employee_id uuid UNIQUE REFERENCES employees(id),
			npwp varchar,
			bpjs_kesehatan varchar,
			bpjs_ketenagakerjaan varchar,
			passport_number varchar,
			passport_expired_date date,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now()
		);`).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
		CREATE TABLE employee_payroll_profiles (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			employee_id uuid UNIQUE REFERENCES employees(id),
			basic_salary numeric(15,2),
			bank_name varchar,
			bank_account_number varchar,
			bank_account_holder varchar,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now()
		);`).Error; err != nil {
			return err
		}

		// =====================================================
		// LOCATIONS (NON POSTGIS)
		// =====================================================
		if err := tx.Exec(`
		CREATE TABLE locations (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			name varchar NOT NULL,

			latitude decimal,
			longitude decimal,
			radius_meters int,

			polygon jsonb,

			is_active boolean DEFAULT true,
			created_at timestamptz DEFAULT now()
		);`).Error; err != nil {
			return err
		}

		// =====================================================
		// ATTENDANCE
		// =====================================================
		if err := tx.Exec(`
		CREATE TABLE attendance (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			employee_id uuid REFERENCES employees(id),
			location_id uuid REFERENCES locations(id),
			check_in_time timestamptz,
			check_out_time timestamptz,
			status varchar,
			created_at timestamptz DEFAULT now()
		);`).Error; err != nil {
			return err
		}

		// =====================================================
		// PAYROLL
		// =====================================================
		if err := tx.Exec(`
		CREATE TABLE payroll_periods (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			month int,
			year int,
			start_date date,
			end_date date,
			is_closed boolean DEFAULT false
		);`).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
		CREATE TABLE payrolls (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			employee_id uuid REFERENCES employees(id),
			payroll_period_id uuid REFERENCES payroll_periods(id),
			basic_salary numeric(15,2),
			total_allowance numeric(15,2),
			total_deduction numeric(15,2),
			net_salary numeric(15,2),
			generated_at timestamptz DEFAULT now()
		);`).Error; err != nil {
			return err
		}

		// =====================================================
		// AUDIT LOG (ISO 27001 READY)
		// =====================================================
		if err := tx.Exec(`
		CREATE TABLE audit_logs (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			user_id uuid REFERENCES users(id),
			action varchar NOT NULL,
			entity varchar NOT NULL,
			entity_id uuid,
			old_values jsonb,
			new_values jsonb,
			ip_address varchar,
			user_agent text,
			request_id uuid,
			source varchar,
			severity varchar,
			created_at timestamptz NOT NULL DEFAULT now()
		);`).Error; err != nil {
			return err
		}

		return nil
	})
}

func DownInitHrisSchema(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		tables := []string{
			"audit_logs",
			"payrolls",
			"payroll_periods",
			"attendance",
			"locations",
			"employee_payroll_profiles",
			"employee_legal_infos",
			"employee_addresses",
			"employee_personal_infos",
			"employees",
			"positions",
			"departments",
			"role_permissions",
			"user_roles",
			"permissions",
			"roles",
		}

		for _, table := range tables {
			if err := tx.Exec(`DROP TABLE IF EXISTS ` + table + ` CASCADE;`).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
