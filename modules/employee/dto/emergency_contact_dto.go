package dto

type EmergencyContactCreateRequest struct {
    Name         string `json:"name" binding:"required"`
    Relationship string `json:"relationship" binding:"required"`
    Phone        string `json:"phone" binding:"required"`
}

type EmergencyContactResponse struct {
    ID           string `json:"id"`
    EmployeeID   string `json:"employee_id"`
    Name         string `json:"name"`
    Relationship string `json:"relationship"`
    Phone        string `json:"phone"`
}
