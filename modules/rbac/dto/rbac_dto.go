package dto

const (
	MESSAGE_FAILED_GET_DATA_FROM_BODY = "failed get data from body"
	MESSAGE_SUCCESS_GET_DATA         = "success get data"
)

type (
	RoleCreateRequest struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
	}

	RoleUpdateRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	RoleResponse struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	PermissionCreateRequest struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
	}

	PermissionUpdateRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	PermissionResponse struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	AssignRoleRequest struct {
		RoleID string `json:"role_id" binding:"required"`
	}
)
