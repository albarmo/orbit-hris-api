package service

import (
	"context"
	"errors"
	"time"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/Caknoooo/go-gin-clean-starter/modules/attendance/dto"
	"github.com/Caknoooo/go-gin-clean-starter/modules/attendance/repository"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/pagination"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AttendanceService interface {
	GetByID(id string) (*entities.Attendance, error)
	CheckIn(req dto.CheckInDTO) (*entities.Attendance, error)
	CheckOut(req dto.CheckOutDTO) (*entities.Attendance, error)
	Update(id string, req dto.UpdateAttendanceDTO) (*entities.Attendance, error)
	Approve(id string, req dto.ApproveAttendanceDTO) (*entities.Attendance, error)
	Delete(id string) error
	FindAll(ctx context.Context, filter *pagination.Filter) (*pagination.Page[entities.Attendance], error)
}

type attendanceService struct {
	attendanceRepository repository.AttendanceRepository
	db                   *gorm.DB
}

func NewAttendanceService(
	attendanceRepo repository.AttendanceRepository,
	db *gorm.DB,
) AttendanceService {
	return &attendanceService{
		attendanceRepository: attendanceRepo,
		db:                   db,
	}
}

func (s *attendanceService) GetByID(id string) (*entities.Attendance, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}

	data, err := s.attendanceRepository.FindByID(uid)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *attendanceService) FindAll(ctx context.Context, filter *pagination.Filter) (*pagination.Page[entities.Attendance], error) {
	page, err := s.attendanceRepository.FindAll(ctx, s.db, filter)
	if err != nil {
		return nil, err
	}
	return page, nil
}

func (s *attendanceService) CheckIn(req dto.CheckInDTO) (*entities.Attendance, error) {
	// Check if already checked in today
	_, err := s.attendanceRepository.FindTodayByEmployeeID(req.EmployeeID)
	if err == nil {
		return nil, errors.New("already checked in today")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	newAttendance := &entities.Attendance{
		EmployeeID:  req.EmployeeID,
		LocationID:  req.LocationID,
		CheckInTime: time.Now(),
		Status:      "present",
	}

	return s.attendanceRepository.Create(newAttendance)
}

func (s *attendanceService) CheckOut(req dto.CheckOutDTO) (*entities.Attendance, error) {
	// Find today's attendance
	attendance, err := s.attendanceRepository.FindTodayByEmployeeID(req.EmployeeID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no check-in record found for today")
		}
		return nil, err
	}

	if attendance.CheckOutTime != nil {
		return nil, errors.New("already checked out today")
	}

	now := time.Now()
	attendance.CheckOutTime = &now

	return s.attendanceRepository.Update(attendance)
}

func (s *attendanceService) Update(id string, req dto.UpdateAttendanceDTO) (*entities.Attendance, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}

	attendance, err := s.attendanceRepository.FindByID(uid)
	if err != nil {
		return nil, err
	}

	attendance.Status = req.Status

	return s.attendanceRepository.Update(attendance)
}

func (s *attendanceService) Approve(id string, req dto.ApproveAttendanceDTO) (*entities.Attendance, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}

	attendance, err := s.attendanceRepository.FindByID(uid)
	if err != nil {
		return nil, err
	}

	attendance.Status = req.Status

	return s.attendanceRepository.Update(attendance)
}

func (s *attendanceService) Delete(id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return errors.New("invalid id")
	}

	_, err = s.attendanceRepository.FindByID(uid)
	if err != nil {
		return err
	}

	return s.attendanceRepository.Delete(uid)
}
