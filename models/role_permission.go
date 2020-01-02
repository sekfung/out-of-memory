package models

type RolePermission struct {
	BaseModel
	RoleId       uint32 `gorm:"not null;index:role_id" json:"role_id"`
	PermissionId uint32 `gorm:"not null;index:permission_id" json:"permission_id"`
}
