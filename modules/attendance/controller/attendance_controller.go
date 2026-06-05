package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/Caknoooo/go-gin-clean-starter/modules/attendance/dto"
	"github.com/Caknoooo/go-gin-clean-starter/modules/attendance/service"
	"github.com/Caknoooo/go-gin-clean-starter/modules/attendance/validation"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/constants"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/pagination"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/samber/do"
	"gorm.io/gorm"
)

const (
	checkInIdempotencyHeader   = "Idempotency-Key"
	idempotencyStateProcessing = "processing"
	idempotencyStateCompleted  = "completed"
)

var (
	checkInProcessingTTL = 90 * time.Second
	checkInCompletedTTL  = 24 * time.Hour
)

type checkInIdempotencyRecord struct {
	State       string          `json:"state"`
	RequestHash string          `json:"request_hash"`
	StatusCode  int             `json:"status_code,omitempty"`
	Response    json.RawMessage `json:"response,omitempty"`
}

type (
	AttendanceController interface {
		GetAll(ctx *gin.Context)
		GetTodayAttendances(ctx *gin.Context)
		GetByID(ctx *gin.Context)
		GetByEmployeeID(ctx *gin.Context)
		CheckIn(ctx *gin.Context)
		CheckOut(ctx *gin.Context)
		Update(ctx *gin.Context)
		Approve(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	attendanceController struct {
		service    service.AttendanceService
		validation *validation.AttendanceValidation
		db         *gorm.DB
		redis      *redis.Client
	}
)

func NewAttendanceController(injector *do.Injector, s service.AttendanceService) AttendanceController {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	var redisClient *redis.Client
	if c, err := do.InvokeNamed[*redis.Client](injector, constants.REDISClient); err == nil {
		redisClient = c
	}

	return &attendanceController{
		service:    s,
		validation: validation.NewAttendanceValidation(),
		db:         db,
		redis:      redisClient,
	}
}

// GetAll godoc
// @Summary Get all attendances
// @Description Get all attendances
// @Tags attendances
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Success 200 {object} utils.Response
// @Router /attendances [get]
func (c *attendanceController) GetAll(ctx *gin.Context) {
	var filter = pagination.Filter{}
	filter.Bind(ctx)

	page, err := c.service.FindAll(ctx.Request.Context(), &filter)
	if err != nil {
		res := utils.BuildResponseFailed("failed get attendances", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success", page)
	ctx.JSON(http.StatusOK, res)
}

func (c *attendanceController) GetTodayAttendances(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	attendance, err := c.service.FindToday(ctx.Request.Context(), userID)
	if err != nil {
		res := utils.BuildResponseFailed("failed get today attendances", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success", attendance)
	ctx.JSON(http.StatusOK, res)
}

// GetByID godoc
// @Summary Get attendance by ID
// @Description Get attendance by ID
// @Tags attendances
// @Accept json
// @Produce json
// @Param id path string true "Attendance ID"
// @Success 200 {object} utils.Response
// @Router /attendances/{id} [get]
func (c *attendanceController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := c.service.GetByID(id)
	if err != nil {
		res := utils.BuildResponseFailed("not found", err.Error(), nil)
		ctx.JSON(http.StatusNotFound, res)
		return
	}
	res := utils.BuildResponseSuccess("success", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *attendanceController) GetByEmployeeID(ctx *gin.Context) {
	employeeID := ctx.Param("employee_id")

	var filter = pagination.Filter{}
	filter.Bind(ctx)

	page, err := c.service.FindByEmployeeID(ctx.Request.Context(), employeeID, &filter)
	if err != nil {
		res := utils.BuildResponseFailed("failed get attendances by employee", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success", page)
	ctx.JSON(http.StatusOK, res)
}

// CheckIn godoc
// @Summary Check in
// @Description Check in
// @Tags attendances
// @Accept json
// @Produce json
// @Param body body dto.CheckInDTO true "CheckIn DTO"
// @Success 201 {object} utils.Response
// @Router /attendances/check-in [post]
func (c *attendanceController) CheckIn(ctx *gin.Context) {
	var req dto.CheckInDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("failed get data from body", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.validation.CheckIn(req); err != nil {
		res := utils.BuildResponseFailed("validation failed", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	idempotencyKey := strings.TrimSpace(ctx.GetHeader(checkInIdempotencyHeader))
	redisKey := ""
	useIdempotency := idempotencyKey != "" && c.redis != nil
	reqHash := ""

	if useIdempotency {
		reqHash = hashCheckInRequest(req)
		redisKey = buildCheckInIdempotencyRedisKey(req.EmployeeID, idempotencyKey)

		record := checkInIdempotencyRecord{
			State:       idempotencyStateProcessing,
			RequestHash: reqHash,
		}
		recordPayload, err := json.Marshal(record)
		if err != nil {
			useIdempotency = false
		} else {
			isNew, err := c.redis.SetNX(ctx.Request.Context(), redisKey, recordPayload, checkInProcessingTTL).Result()
			if err != nil {
				useIdempotency = false
			} else if !isNew {
				// Existing idempotency key found: replay completed response or reject conflicting/in-progress requests.
				cachedPayload, err := c.redis.Get(ctx.Request.Context(), redisKey).Bytes()
				if err != nil {
					res := utils.BuildResponseFailed("failed check-in", "idempotency key is being processed", nil)
					ctx.JSON(http.StatusConflict, res)
					return
				}

				var cached checkInIdempotencyRecord
				if err := json.Unmarshal(cachedPayload, &cached); err != nil {
					res := utils.BuildResponseFailed("failed check-in", "idempotency record is invalid", nil)
					ctx.JSON(http.StatusConflict, res)
					return
				}

				if cached.RequestHash != reqHash {
					res := utils.BuildResponseFailed("failed check-in", "idempotency key already used with a different request", nil)
					ctx.JSON(http.StatusConflict, res)
					return
				}

				if cached.State == idempotencyStateCompleted && cached.StatusCode != 0 && len(cached.Response) > 0 {
					ctx.Header("Idempotency-Replayed", "true")
					ctx.Data(cached.StatusCode, "application/json; charset=utf-8", cached.Response)
					return
				}

				res := utils.BuildResponseFailed("failed check-in", "idempotency key is being processed", nil)
				ctx.JSON(http.StatusConflict, res)
				return
			}
		}
	}

	result, err := c.service.CheckIn(req)
	if err != nil {
		if useIdempotency && redisKey != "" {
			// Release the key on failure so the same idempotency key can be retried.
			_ = c.redis.Del(ctx.Request.Context(), redisKey).Err()
		}

		res := utils.BuildResponseFailed("failed check-in", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("check-in successful", result)

	if useIdempotency && redisKey != "" {
		responsePayload, err := json.Marshal(res)
		if err == nil {
			finalRecord := checkInIdempotencyRecord{
				State:       idempotencyStateCompleted,
				RequestHash: reqHash,
				StatusCode:  http.StatusCreated,
				Response:    responsePayload,
			}
			if finalPayload, err := json.Marshal(finalRecord); err == nil {
				_ = c.redis.Set(ctx.Request.Context(), redisKey, finalPayload, checkInCompletedTTL).Err()
			}
		}
	}

	ctx.JSON(http.StatusCreated, res)
}

// CheckOut godoc
// @Summary Check out
// @Description Check out
// @Tags attendances
// @Accept json
// @Produce json
// @Param body body dto.CheckOutDTO true "CheckOut DTO"
// @Success 200 {object} utils.Response
// @Router /attendances/check-out [put]
func (c *attendanceController) CheckOut(ctx *gin.Context) {
	var req dto.CheckOutDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("failed get data from body", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.validation.CheckOut(req); err != nil {
		res := utils.BuildResponseFailed("validation failed", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.service.CheckOut(req)
	if err != nil {
		res := utils.BuildResponseFailed("failed check-out", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("check-out successful", result)
	ctx.JSON(http.StatusOK, res)
}

// Update godoc
// @Summary Update attendance
// @Description Update attendance
// @Tags attendances
// @Accept json
// @Produce json
// @Param id path string true "Attendance ID"
// @Param body body dto.UpdateAttendanceDTO true "UpdateAttendance DTO"
// @Success 200 {object} utils.Response
// @Router /attendances/{id} [put]
func (c *attendanceController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var req dto.UpdateAttendanceDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("failed get data from body", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.validation.UpdateAttendance(req); err != nil {
		res := utils.BuildResponseFailed("validation failed", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.service.Update(id, req)
	if err != nil {
		res := utils.BuildResponseFailed("failed update attendance", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("update successful", result)
	ctx.JSON(http.StatusOK, res)
}

// Approve godoc
// @Summary Approve attendance
// @Description Approve attendance
// @Tags attendances
// @Accept json
// @Produce json
// @Param id path string true "Attendance ID"
// @Param body body dto.ApproveAttendanceDTO true "ApproveAttendance DTO"
// @Success 200 {object} utils.Response
// @Router /attendances/{id}/approve [put]
func (c *attendanceController) Approve(ctx *gin.Context) {
	id := ctx.Param("id")
	var req dto.ApproveAttendanceDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("failed get data from body", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.validation.ApproveAttendance(req); err != nil {
		res := utils.BuildResponseFailed("validation failed", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.service.Approve(id, req)
	if err != nil {
		res := utils.BuildResponseFailed("failed approval", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("approval successful", result)
	ctx.JSON(http.StatusOK, res)
}

// Delete godoc
// @Summary Delete attendance
// @Description Delete attendance
// @Tags attendances
// @Accept json
// @Produce json
// @Param id path string true "Attendance ID"
// @Success 200 {object} utils.Response
// @Router /attendances/{id} [delete]
func (c *attendanceController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.service.Delete(id)
	if err != nil {
		res := utils.BuildResponseFailed("failed delete attendance", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("delete successful", nil)
	ctx.JSON(http.StatusOK, res)
}

func buildCheckInIdempotencyRedisKey(employeeID uuid.UUID, idempotencyKey string) string {
	// Scope the key by employee and date so a reused key across different days does not collide forever.
	datePart := time.Now().UTC().Format("2006-01-02")
	return "idem:attendance:checkin:" + employeeID.String() + ":" + datePart + ":" + idempotencyKey
}

func hashCheckInRequest(req dto.CheckInDTO) string {
	sum := sha256.Sum256([]byte(req.EmployeeID.String() + "|" + req.LocationID.String()))
	return hex.EncodeToString(sum[:])
}
