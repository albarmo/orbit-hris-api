package dto

const (
	MESSAGE_FAILED_GET_DATA_FROM_BODY = "failed get data from body"
	MESSAGE_SUCCESS_GET_DATA         = "success get data"
)

type (
	MasterCreateRequest struct {
	}

	MasterResponse struct {
	}
)

// Department DTOs
type DepartmentCreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type DepartmentUpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DepartmentResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Location DTOs
type LocationCreateRequest struct {
	Name         string  `json:"name" binding:"required"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	RadiusMeters int     `json:"radius_meters"`
}

type LocationUpdateRequest struct {
	Name         string  `json:"name"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	RadiusMeters int     `json:"radius_meters"`
}

type LocationResponse struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	RadiusMeters int     `json:"radius_meters"`
}

// Position DTOs
type PositionCreateRequest struct {
	DepartmentID string `json:"department_id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Level        string `json:"level"`
}

type PositionUpdateRequest struct {
	DepartmentID string `json:"department_id"`
	Name         string `json:"name"`
	Level        string `json:"level"`
}

type PositionResponse struct {
	ID           string `json:"id"`
	DepartmentID string `json:"department_id"`
	Name         string `json:"name"`
	Level        string `json:"level"`
}
