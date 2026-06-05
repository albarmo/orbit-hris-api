package tests

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	attendanceRepository "github.com/Caknoooo/go-gin-clean-starter/modules/attendance/repository"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/pagination"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type attendanceRepoSeed struct {
	userID      uuid.UUID
	user2ID     uuid.UUID
	employeeID  uuid.UUID
	employee2ID uuid.UUID
	locationID  uuid.UUID
}

func setupAttendanceRepository(t *testing.T) (attendanceRepository.AttendanceRepository, *gorm.DB, attendanceRepoSeed) {
	t.Helper()

	dsn := "file:" + uuid.NewString() + "?mode=memory&cache=shared"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	require.NoError(t, err)

	require.NoError(t, db.Exec(`
		CREATE TABLE employees (
			id TEXT PRIMARY KEY,
			user_id TEXT,
			employee_code TEXT,
			supervisor_id TEXT,
			department_id TEXT,
			position_id TEXT,
			join_date DATETIME,
			end_date DATETIME,
			employment_type TEXT,
			employment_status TEXT,
			probation_end_date DATETIME,
			created_at DATETIME,
			updated_at DATETIME
		)
	`).Error)

	require.NoError(t, db.Exec(`
		CREATE TABLE locations (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			latitude REAL,
			longitude REAL,
			radius_meters INTEGER,
			polygon TEXT,
			is_active BOOLEAN,
			created_at DATETIME
		)
	`).Error)

	require.NoError(t, db.Exec(`
		CREATE TABLE attendance (
			id TEXT PRIMARY KEY,
			employee_id TEXT,
			location_id TEXT,
			check_in_time DATETIME,
			check_out_time DATETIME,
			status TEXT,
			created_at DATETIME
		)
	`).Error)

	seed := attendanceRepoSeed{
		userID:      uuid.New(),
		user2ID:     uuid.New(),
		employeeID:  uuid.New(),
		employee2ID: uuid.New(),
		locationID:  uuid.New(),
	}

	require.NoError(t, db.Exec(
		`INSERT INTO locations (id, name) VALUES (?, ?)`,
		seed.locationID.String(), "Head Office",
	).Error)

	require.NoError(t, db.Exec(
		`INSERT INTO employees (id, user_id, employee_code) VALUES (?, ?, ?)`,
		seed.employeeID.String(), seed.userID.String(), "EMP-001",
	).Error)

	require.NoError(t, db.Exec(
		`INSERT INTO employees (id, user_id, employee_code) VALUES (?, ?, ?)`,
		seed.employee2ID.String(), seed.user2ID.String(), "EMP-002",
	).Error)

	return attendanceRepository.NewAttendanceRepository(db), db, seed
}

func insertAttendanceRow(t *testing.T, db *gorm.DB, attendanceID, employeeID, locationID uuid.UUID, checkInTime time.Time, status string) {
	t.Helper()

	require.NoError(t, db.Exec(
		`INSERT INTO attendance (id, employee_id, location_id, check_in_time, status, created_at) VALUES (?, ?, ?, ?, ?, ?)`,
		attendanceID.String(), employeeID.String(), locationID.String(), checkInTime, status, time.Now(),
	).Error)
}

func TestAttendanceRepository_CRUD(t *testing.T) {
	repo, _, seed := setupAttendanceRepository(t)
	attendanceID := uuid.New()
	checkIn := time.Now()

	created, err := repo.Create(&entities.Attendance{
		ID:          attendanceID,
		EmployeeID:  seed.employeeID,
		LocationID:  seed.locationID,
		CheckInTime: checkIn,
		Status:      "present",
		CreatedAt:   time.Now(),
	})

	require.NoError(t, err)
	require.NotNil(t, created)
	assert.Equal(t, attendanceID, created.ID)
	assert.Equal(t, seed.employeeID, created.EmployeeID)
	assert.Equal(t, seed.locationID, created.LocationID)

	found, err := repo.FindByID(attendanceID)
	require.NoError(t, err)
	assert.Equal(t, attendanceID, found.ID)
	assert.Equal(t, "present", found.Status)

	now := time.Now()
	found.Status = "approved"
	found.CheckOutTime = &now

	updated, err := repo.Update(found)
	require.NoError(t, err)
	assert.Equal(t, "approved", updated.Status)
	require.NotNil(t, updated.CheckOutTime)

	err = repo.Delete(attendanceID)
	require.NoError(t, err)

	_, err = repo.FindByID(attendanceID)
	require.Error(t, err)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}

func TestAttendanceRepository_FindTodayByEmployeeID(t *testing.T) {
	repo, db, seed := setupAttendanceRepository(t)
	oldAttendanceID := uuid.New()
	todayAttendanceID := uuid.New()

	insertAttendanceRow(t, db, oldAttendanceID, seed.employeeID, seed.locationID, time.Now().Add(-48*time.Hour), "present")
	insertAttendanceRow(t, db, todayAttendanceID, seed.employeeID, seed.locationID, time.Now(), "present")

	result, err := repo.FindTodayByEmployeeID(seed.employeeID)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, todayAttendanceID, result.ID)
	assert.Equal(t, seed.employeeID, result.EmployeeID)
}

func TestAttendanceRepository_FindAll_And_FindByEmployeeID(t *testing.T) {
	repo, db, seed := setupAttendanceRepository(t)

	insertAttendanceRow(t, db, uuid.New(), seed.employeeID, seed.locationID, time.Now(), "present")
	insertAttendanceRow(t, db, uuid.New(), seed.employee2ID, seed.locationID, time.Now(), "present")

	allPage, err := repo.FindAll(context.Background(), nil, &pagination.Filter{Page: 1, Limit: 10})
	require.NoError(t, err)
	require.NotNil(t, allPage)
	assert.Equal(t, int64(2), allPage.Total)
	assert.Len(t, allPage.Data, 2)

	employeePage, err := repo.FindByEmployeeID(context.Background(), nil, &pagination.Filter{Page: 1, Limit: 10}, seed.employeeID)
	require.NoError(t, err)
	require.NotNil(t, employeePage)
	assert.Equal(t, int64(1), employeePage.Total)
	assert.Len(t, employeePage.Data, 1)
	assert.Equal(t, seed.employeeID, employeePage.Data[0].EmployeeID)
}

func TestAttendanceRepository_FindToday_ByUserID(t *testing.T) {
	repo, db, seed := setupAttendanceRepository(t)
	expectedID := uuid.New()

	insertAttendanceRow(t, db, expectedID, seed.employeeID, seed.locationID, time.Now(), "present")
	insertAttendanceRow(t, db, uuid.New(), seed.employee2ID, seed.locationID, time.Now(), "present")

	result, err := repo.FindToday(context.Background(), nil, seed.userID)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, expectedID, result.ID)
	assert.Equal(t, seed.employeeID, result.EmployeeID)
}
