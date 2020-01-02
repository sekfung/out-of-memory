package models

type Role struct {
	BaseModel
	Name   string `gorm:"not null;index:role_name" json:"name"`
	Desc   string `gorm:"not null" json:"desc"`
	Status bool   `json:"status"`
}
