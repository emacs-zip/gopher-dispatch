package dto

type CreateRoleModel struct {
    RoleName string `json:"roleName" binding:"required"`
}
