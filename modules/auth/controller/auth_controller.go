package controller

import (
	"io"
	"net/http"

	"github.com/Caknoooo/go-gin-clean-starter/modules/auth/dto"
	"github.com/Caknoooo/go-gin-clean-starter/modules/auth/service"
	"github.com/Caknoooo/go-gin-clean-starter/modules/auth/validation"
	userDto "github.com/Caknoooo/go-gin-clean-starter/modules/user/dto"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/constants"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type (
	AuthController interface {
		Register(ctx *gin.Context)
		Login(ctx *gin.Context)
		LoginByFace(ctx *gin.Context)
		EnrollFace(ctx *gin.Context)
		GetPerson(ctx *gin.Context)
		GetPhoto(ctx *gin.Context)
		RefreshToken(ctx *gin.Context)
		Logout(ctx *gin.Context)
		SendVerificationEmail(ctx *gin.Context)
		VerifyEmail(ctx *gin.Context)
		SendPasswordReset(ctx *gin.Context)
		ResetPassword(ctx *gin.Context)
	}

	authController struct {
		authService    service.AuthService
		authValidation *validation.AuthValidation
		db             *gorm.DB
	}
)

func NewAuthController(injector *do.Injector, as service.AuthService) AuthController {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	authValidation := validation.NewAuthValidation()
	return &authController{
		authService:    as,
		authValidation: authValidation,
		db:             db,
	}
}

func (c *authController) Register(ctx *gin.Context) {
	var req userDto.UserCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(userDto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Validate request
	if err := c.authValidation.ValidateRegisterRequest(req); err != nil {
		res := utils.BuildResponseFailed("Validation failed", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.authService.Register(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(userDto.MESSAGE_FAILED_REGISTER_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(userDto.MESSAGE_SUCCESS_REGISTER_USER, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *authController) Login(ctx *gin.Context) {
	var req userDto.UserLoginRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed(userDto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Validate request
	if err := c.authValidation.ValidateLoginRequest(req); err != nil {
		res := utils.BuildResponseFailed("Validation failed", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.authService.Login(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(userDto.MESSAGE_FAILED_LOGIN, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(userDto.MESSAGE_SUCCESS_LOGIN, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *authController) RefreshToken(ctx *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(userDto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.authService.RefreshToken(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REFRESH_TOKEN, err.Error(), nil)
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REFRESH_TOKEN, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *authController) Logout(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(string)

	err := c.authService.Logout(ctx.Request.Context(), userId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_LOGOUT, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_LOGOUT, nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *authController) SendVerificationEmail(ctx *gin.Context) {
	var req userDto.SendVerificationEmailRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(userDto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err := c.authService.SendVerificationEmail(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(userDto.MESSAGE_FAILED_PROSES_REQUEST, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(userDto.MESSAGE_SEND_VERIFICATION_EMAIL_SUCCESS, nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *authController) VerifyEmail(ctx *gin.Context) {
	var req userDto.VerifyEmailRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(userDto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.authService.VerifyEmail(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(userDto.MESSAGE_FAILED_VERIFY_EMAIL, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(userDto.MESSAGE_SUCCESS_VERIFY_EMAIL, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *authController) SendPasswordReset(ctx *gin.Context) {
	var req dto.SendPasswordResetRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(userDto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err := c.authService.SendPasswordReset(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_SEND_PASSWORD_RESET, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_SEND_PASSWORD_RESET, nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *authController) ResetPassword(ctx *gin.Context) {
	var req dto.ResetPasswordRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(userDto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err := c.authService.ResetPassword(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_RESET_PASSWORD, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_RESET_PASSWORD, nil)
	ctx.JSON(http.StatusOK, res)
}

// LoginByFace handles face ID login using external face-search API
func (c *authController) LoginByFace(ctx *gin.Context) {
	// receive uploaded image file from form field "image"
	fileHeader, err := ctx.FormFile("image")
	if err != nil {
		res := utils.BuildResponseFailed("failed get file", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		res := utils.BuildResponseFailed("failed open file", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	defer file.Close()

	// read file bytes
	imgBytes, err := io.ReadAll(file)
	if err != nil {
		res := utils.BuildResponseFailed("failed read file", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// call service to authenticate by face
	result, err := c.authService.LoginByFace(ctx.Request.Context(), imgBytes, fileHeader.Filename)
	if err != nil {
		res := utils.BuildResponseFailed("failed login by face", err.Error(), nil)
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	res := utils.BuildResponseSuccess("success", result)
	ctx.JSON(http.StatusOK, res)
}

// EnrollFace registers a new person/photo to external face service
func (c *authController) EnrollFace(ctx *gin.Context) {
	name := ctx.Query("name")
	if name == "" {
		res := utils.BuildResponseFailed("name query param required", "", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	fileHeader, err := ctx.FormFile("image")
	if err != nil {
		res := utils.BuildResponseFailed("failed get file", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		res := utils.BuildResponseFailed("failed open file", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	defer file.Close()

	imgBytes, err := io.ReadAll(file)
	if err != nil {
		res := utils.BuildResponseFailed("failed read file", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	resp, err := c.authService.EnrollFace(ctx.Request.Context(), imgBytes, fileHeader.Filename, name)
	if err != nil {
		res := utils.BuildResponseFailed("enroll failed", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("enroll success", resp)
	ctx.JSON(http.StatusOK, res)
}

// GetPerson returns person info from external face service
func (c *authController) GetPerson(ctx *gin.Context) {
	name := ctx.Param("name")
	if name == "" {
		res := utils.BuildResponseFailed("name required", "", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	resp, err := c.authService.GetPerson(ctx.Request.Context(), name)
	if err != nil {
		res := utils.BuildResponseFailed("failed get person", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success", resp)
	ctx.JSON(http.StatusOK, res)
}

// GetPhoto proxies photo bytes from external face service
func (c *authController) GetPhoto(ctx *gin.Context) {
	photoID := ctx.Param("id")
	if photoID == "" {
		res := utils.BuildResponseFailed("photo id required", "", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	contentType, data, err := c.authService.GetPhoto(ctx.Request.Context(), photoID)
	if err != nil {
		res := utils.BuildResponseFailed("failed get photo", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if contentType == "" {
		contentType = "application/octet-stream"
	}
	ctx.Data(http.StatusOK, contentType, data)
}
