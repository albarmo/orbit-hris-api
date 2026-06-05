package tests

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/Caknoooo/go-gin-clean-starter/modules/attendance/dto"
	attendanceService "github.com/Caknoooo/go-gin-clean-starter/modules/attendance/service"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/pagination"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

type attendanceRepositoryMock struct {
	findAllFn               func(ctx context.Context, db *gorm.DB, filter *pagination.Filter) (*pagination.Page[entities.Attendance], error)
	findByEmployeeIDFn      func(ctx context.Context, db *gorm.DB, filter *pagination.Filter, employeeID uuid.UUID) (*pagination.Page[entities.Attendance], error)
	findTodayFn             func(ctx context.Context, db *gorm.DB, userID uuid.UUID) (*entities.Attendance, error)
	findByIDFn              func(id uuid.UUID) (*entities.Attendance, error)
	findTodayByEmployeeIDFn func(employeeID uuid.UUID) (*entities.Attendance, error)
	createFn                func(attendance *entities.Attendance) (*entities.Attendance, error)
	updateFn                func(attendance *entities.Attendance) (*entities.Attendance, error)
	deleteFn                func(id uuid.UUID) error
}

func (m *attendanceRepositoryMock) FindAll(ctx context.Context, db *gorm.DB, filter *pagination.Filter) (*pagination.Page[entities.Attendance], error) {
	if m.findAllFn != nil {
		return m.findAllFn(ctx, db, filter)
	}
	return nil, nil
}

func (m *attendanceRepositoryMock) FindByEmployeeID(ctx context.Context, db *gorm.DB, filter *pagination.Filter, employeeID uuid.UUID) (*pagination.Page[entities.Attendance], error) {
	if m.findByEmployeeIDFn != nil {
		return m.findByEmployeeIDFn(ctx, db, filter, employeeID)
	}
	return nil, nil
}

func (m *attendanceRepositoryMock) FindToday(ctx context.Context, db *gorm.DB, userID uuid.UUID) (*entities.Attendance, error) {
	if m.findTodayFn != nil {
		return m.findTodayFn(ctx, db, userID)
	}
	return nil, nil
}

func (m *attendanceRepositoryMock) FindByID(id uuid.UUID) (*entities.Attendance, error) {
	if m.findByIDFn != nil {
		return m.findByIDFn(id)
	}
	return nil, nil
}

func (m *attendanceRepositoryMock) FindTodayByEmployeeID(employeeID uuid.UUID) (*entities.Attendance, error) {
	if m.findTodayByEmployeeIDFn != nil {
		return m.findTodayByEmployeeIDFn(employeeID)
	}
	return nil, nil
}

func (m *attendanceRepositoryMock) Create(attendance *entities.Attendance) (*entities.Attendance, error) {
	if m.createFn != nil {
		return m.createFn(attendance)
	}
	return attendance, nil
}

func (m *attendanceRepositoryMock) Update(attendance *entities.Attendance) (*entities.Attendance, error) {
	if m.updateFn != nil {
		return m.updateFn(attendance)
	}
	return attendance, nil
}

func (m *attendanceRepositoryMock) Delete(id uuid.UUID) error {
	if m.deleteFn != nil {
		return m.deleteFn(id)
	}
	return nil
}

func TestAttendanceService_GetByID_InvalidID(t *testing.T) {
	svc := attendanceService.NewAttendanceService(&attendanceRepositoryMock{}, nil)

	result, err := svc.GetByID("not-a-uuid")

	require.Error(t, err)
	require.Nil(t, result)
	assert.Equal(t, "invalid id", err.Error())
}

func TestAttendanceService_GetByID_Success(t *testing.T) {
	wantID := uuid.New()
	svc := attendanceService.NewAttendanceService(&attendanceRepositoryMock{
		findByIDFn: func(id uuid.UUID) (*entities.Attendance, error) {
			assert.Equal(t, wantID, id)
			return &entities.Attendance{ID: id, Status: "present"}, nil
		},
	}, nil)

	result, err := svc.GetByID(wantID.String())

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, wantID, result.ID)
	assert.Equal(t, "present", result.Status)
}

func TestAttendanceService_FindAll_Success(t *testing.T) {
	wantPage := &pagination.Page[entities.Attendance]{
		Page:  1,
		Limit: 10,
		Total: 1,
		Data: []entities.Attendance{
			{ID: uuid.New()},
		},
	}

	svc := attendanceService.NewAttendanceService(&attendanceRepositoryMock{
		findAllFn: func(ctx context.Context, db *gorm.DB, filter *pagination.Filter) (*pagination.Page[entities.Attendance], error) {
			require.NotNil(t, filter)
			return wantPage, nil
		},
	}, nil)

	result, err := svc.FindAll(context.Background(), &pagination.Filter{Page: 1, Limit: 10})

	require.NoError(t, err)
	assert.Equal(t, wantPage, result)
}

func TestAttendanceService_FindByEmployeeID_InvalidID(t *testing.T) {
	svc := attendanceService.NewAttendanceService(&attendanceRepositoryMock{}, nil)

	result, err := svc.FindByEmployeeID(context.Background(), "invalid-id", &pagination.Filter{Page: 1, Limit: 10})

	require.Error(t, err)
	require.Nil(t, result)
	assert.Equal(t, "invalid employee id", err.Error())
}

func TestAttendanceService_CheckIn_AlreadyCheckedIn(t *testing.T) {
	employeeID := uuid.New()

	svc := attendanceService.NewAttendanceService(&attendanceRepositoryMock{
		findTodayByEmployeeIDFn: func(id uuid.UUID) (*entities.Attendance, error) {
			assert.Equal(t, employeeID, id)
			return &entities.Attendance{ID: uuid.New(), EmployeeID: id}, nil
		},
	}, nil)

	result, err := svc.CheckIn(dto.CheckInDTO{
		EmployeeID: employeeID,
		LocationID: uuid.New(),
	})

	require.Error(t, err)
	require.Nil(t, result)
	assert.Equal(t, "already checked in today", err.Error())
}

func TestAttendanceService_CheckIn_Success(t *testing.T) {
	employeeID := uuid.New()
	locationID := uuid.New()
	createCalled := false

	svc := attendanceService.NewAttendanceService(&attendanceRepositoryMock{
		findTodayByEmployeeIDFn: func(id uuid.UUID) (*entities.Attendance, error) {
			return nil, gorm.ErrRecordNotFound
		},
		createFn: func(attendance *entities.Attendance) (*entities.Attendance, error) {
			createCalled = true
			assert.Equal(t, employeeID, attendance.EmployeeID)
			assert.Equal(t, locationID, attendance.LocationID)
			assert.Equal(t, "present", attendance.Status)
			return attendance, nil
		},
	}, nil)

	result, err := svc.CheckIn(dto.CheckInDTO{
		EmployeeID: employeeID,
		LocationID: locationID,
	})

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.True(t, createCalled)
	assert.Equal(t, "present", result.Status)
}

func TestAttendanceService_CheckOut_NoCheckInRecord(t *testing.T) {
	svc := attendanceService.NewAttendanceService(&attendanceRepositoryMock{
		findTodayByEmployeeIDFn: func(employeeID uuid.UUID) (*entities.Attendance, error) {
			return nil, gorm.ErrRecordNotFound
		},
	}, nil)

	result, err := svc.CheckOut(dto.CheckOutDTO{EmployeeID: uuid.New()})

	require.Error(t, err)
	require.Nil(t, result)
	assert.Equal(t, "no check-in record found for today", err.Error())
}

func TestAttendanceService_CheckOut_Success(t *testing.T) {
	employeeID := uuid.New()
	existing := &entities.Attendance{
		ID:         uuid.New(),
		EmployeeID: employeeID,
		Status:     "present",
	}
	updateCalled := false

	svc := attendanceService.NewAttendanceService(&attendanceRepositoryMock{
		findTodayByEmployeeIDFn: func(id uuid.UUID) (*entities.Attendance, error) {
			return existing, nil
		},
		updateFn: func(attendance *entities.Attendance) (*entities.Attendance, error) {
			updateCalled = true
			require.NotNil(t, attendance.CheckOutTime)
			return attendance, nil
		},
	}, nil)

	result, err := svc.CheckOut(dto.CheckOutDTO{EmployeeID: employeeID})

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.True(t, updateCalled)
	assert.NotNil(t, result.CheckOutTime)
}

func TestAttendanceService_Update_Success(t *testing.T) {
	attendanceID := uuid.New()
	updateCalled := false

	svc := attendanceService.NewAttendanceService(&attendanceRepositoryMock{
		findByIDFn: func(id uuid.UUID) (*entities.Attendance, error) {
			assert.Equal(t, attendanceID, id)
			return &entities.Attendance{
				ID:     id,
				Status: "present",
			}, nil
		},
		updateFn: func(attendance *entities.Attendance) (*entities.Attendance, error) {
			updateCalled = true
			assert.Equal(t, "approved", attendance.Status)
			return attendance, nil
		},
	}, nil)

	result, err := svc.Update(attendanceID.String(), dto.UpdateAttendanceDTO{Status: "approved"})

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.True(t, updateCalled)
	assert.Equal(t, "approved", result.Status)
}

func TestAttendanceService_Delete_Success(t *testing.T) {
	attendanceID := uuid.New()
	deleteCalled := false

	svc := attendanceService.NewAttendanceService(&attendanceRepositoryMock{
		findByIDFn: func(id uuid.UUID) (*entities.Attendance, error) {
			return &entities.Attendance{ID: id}, nil
		},
		deleteFn: func(id uuid.UUID) error {
			deleteCalled = true
			assert.Equal(t, attendanceID, id)
			return nil
		},
	}, nil)

	err := svc.Delete(attendanceID.String())

	require.NoError(t, err)
	assert.True(t, deleteCalled)
}

func TestAttendanceService_Delete_InvalidID(t *testing.T) {
	svc := attendanceService.NewAttendanceService(&attendanceRepositoryMock{}, nil)

	err := svc.Delete("invalid-id")

	require.Error(t, err)
	assert.Equal(t, "invalid id", err.Error())
}

func TestAttendanceService_CheckIn_RepositoryFailure(t *testing.T) {
	employeeID := uuid.New()
	expectedErr := errors.New("database unavailable")

	svc := attendanceService.NewAttendanceService(&attendanceRepositoryMock{
		findTodayByEmployeeIDFn: func(id uuid.UUID) (*entities.Attendance, error) {
			return nil, expectedErr
		},
	}, nil)

	result, err := svc.CheckIn(dto.CheckInDTO{
		EmployeeID: employeeID,
		LocationID: uuid.New(),
	})

	require.Error(t, err)
	require.Nil(t, result)
	assert.Equal(t, expectedErr, err)
}

func TestAttendanceService_CheckOut_AlreadyCheckedOut(t *testing.T) {
	employeeID := uuid.New()
	now := time.Now()

	svc := attendanceService.NewAttendanceService(&attendanceRepositoryMock{
		findTodayByEmployeeIDFn: func(id uuid.UUID) (*entities.Attendance, error) {
			return &entities.Attendance{
				ID:           uuid.New(),
				EmployeeID:   employeeID,
				CheckOutTime: &now,
			}, nil
		},
	}, nil)

	result, err := svc.CheckOut(dto.CheckOutDTO{EmployeeID: employeeID})

	require.Error(t, err)
	require.Nil(t, result)
	assert.Equal(t, "already checked out today", err.Error())
}
