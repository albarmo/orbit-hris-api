package tests

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	attendanceModule "github.com/Caknoooo/go-gin-clean-starter/modules/attendance"
	attendanceController "github.com/Caknoooo/go-gin-clean-starter/modules/attendance/controller"
	authService "github.com/Caknoooo/go-gin-clean-starter/modules/auth/service"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/constants"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/samber/do"
	"github.com/stretchr/testify/require"
)

const validTestToken = "valid-token"

type mockAttendanceController struct{}

func (m *mockAttendanceController) GetAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"handler": "GetAll"})
}

func (m *mockAttendanceController) GetTodayAttendances(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"handler": "GetTodayAttendances"})
}

func (m *mockAttendanceController) GetByID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"handler": "GetByID"})
}

func (m *mockAttendanceController) GetByEmployeeID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"handler": "GetByEmployeeID"})
}

func (m *mockAttendanceController) CheckIn(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"handler": "CheckIn"})
}

func (m *mockAttendanceController) CheckOut(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"handler": "CheckOut"})
}

func (m *mockAttendanceController) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"handler": "Update"})
}

func (m *mockAttendanceController) Approve(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"handler": "Approve"})
}

func (m *mockAttendanceController) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"handler": "Delete"})
}

type mockJWTService struct{}

func (m *mockJWTService) GenerateAccessToken(userID string, role string) string {
	return validTestToken
}

func (m *mockJWTService) GenerateRefreshToken() (string, time.Time) {
	return "refresh-token", time.Now().Add(15 * time.Minute)
}

func (m *mockJWTService) ValidateToken(token string) (*jwt.Token, error) {
	if token != validTestToken {
		return nil, errors.New("invalid token")
	}

	return &jwt.Token{Valid: true}, nil
}

func (m *mockJWTService) GetUserIDByToken(token string) (string, error) {
	if token != validTestToken {
		return "", errors.New("invalid token")
	}

	return "11111111-1111-1111-1111-111111111111", nil
}

func (m *mockJWTService) GetAccessExpiry() time.Duration {
	return 15 * time.Minute
}

func setupAttendanceRouter(t *testing.T) *gin.Engine {
	t.Helper()

	gin.SetMode(gin.TestMode)

	server := gin.New()
	injector := do.New()

	do.Provide(injector, func(i *do.Injector) (attendanceController.AttendanceController, error) {
		return &mockAttendanceController{}, nil
	})

	do.ProvideNamed(injector, constants.JWTService, func(i *do.Injector) (authService.JWTService, error) {
		return &mockJWTService{}, nil
	})

	attendanceModule.RegisterRoutes(server, injector)

	return server
}

func TestAttendanceRoutes_AuthorizedRequestsHitExpectedHandlers(t *testing.T) {
	server := setupAttendanceRouter(t)

	testCases := []struct {
		name       string
		method     string
		path       string
		body       string
		handlerTag string
	}{
		{
			name:       "GET /api/attendances",
			method:     http.MethodGet,
			path:       "/api/attendances",
			handlerTag: "GetAll",
		},
		{
			name:       "GET /api/attendances/today",
			method:     http.MethodGet,
			path:       "/api/attendances/today",
			handlerTag: "GetTodayAttendances",
		},
		{
			name:       "GET /api/attendances/:id",
			method:     http.MethodGet,
			path:       "/api/attendances/att-123",
			handlerTag: "GetByID",
		},
		{
			name:       "POST /api/attendances/check-in",
			method:     http.MethodPost,
			path:       "/api/attendances/check-in",
			body:       "{}",
			handlerTag: "CheckIn",
		},
		{
			name:       "PUT /api/attendances/check-out",
			method:     http.MethodPut,
			path:       "/api/attendances/check-out",
			body:       "{}",
			handlerTag: "CheckOut",
		},
		{
			name:       "PUT /api/attendances/:id",
			method:     http.MethodPut,
			path:       "/api/attendances/att-123",
			body:       "{}",
			handlerTag: "Update",
		},
		{
			name:       "PUT /api/attendances/:id/approve",
			method:     http.MethodPut,
			path:       "/api/attendances/att-123/approve",
			body:       "{}",
			handlerTag: "Approve",
		},
		{
			name:       "DELETE /api/attendances/:id",
			method:     http.MethodDelete,
			path:       "/api/attendances/att-123",
			handlerTag: "Delete",
		},
		{
			name:       "GET /api/attendances/employee/:employee_id",
			method:     http.MethodGet,
			path:       "/api/attendances/employee/emp-123",
			handlerTag: "GetByEmployeeID",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, tc.path, strings.NewReader(tc.body))
			req.Header.Set("Authorization", "Bearer "+validTestToken)

			if tc.body != "" {
				req.Header.Set("Content-Type", "application/json")
			}

			rec := httptest.NewRecorder()
			server.ServeHTTP(rec, req)

			require.Equal(t, http.StatusOK, rec.Code)
			require.Contains(t, rec.Body.String(), tc.handlerTag)
		})
	}
}

func TestAttendanceRoutes_RejectRequestWithoutToken(t *testing.T) {
	server := setupAttendanceRouter(t)

	testCases := []struct {
		name   string
		method string
		path   string
	}{
		{name: "GET list", method: http.MethodGet, path: "/api/attendances"},
		{name: "GET today", method: http.MethodGet, path: "/api/attendances/today"},
		{name: "GET by id", method: http.MethodGet, path: "/api/attendances/att-123"},
		{name: "POST check in", method: http.MethodPost, path: "/api/attendances/check-in"},
		{name: "PUT check out", method: http.MethodPut, path: "/api/attendances/check-out"},
		{name: "PUT update", method: http.MethodPut, path: "/api/attendances/att-123"},
		{name: "PUT approve", method: http.MethodPut, path: "/api/attendances/att-123/approve"},
		{name: "DELETE", method: http.MethodDelete, path: "/api/attendances/att-123"},
		{name: "GET by employee", method: http.MethodGet, path: "/api/attendances/employee/emp-123"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, tc.path, nil)
			rec := httptest.NewRecorder()

			server.ServeHTTP(rec, req)

			require.Equal(t, http.StatusUnauthorized, rec.Code)
		})
	}
}

func TestAttendanceRoutes_RejectRequestWithInvalidToken(t *testing.T) {
	server := setupAttendanceRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/api/attendances", nil)
	req.Header.Set("Authorization", "Bearer invalid-token")

	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	require.Equal(t, http.StatusUnauthorized, rec.Code)
}
