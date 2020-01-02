package models

type UserRole struct {
	BaseModel
	Uid    uint32 `gorm:"not null;index:uid" json:"uid"`
	RoleId uint32 `gorm:"not null;index:role_id" json:"role_id"`
}
