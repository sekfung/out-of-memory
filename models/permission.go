package models

type Permission struct {
	BaseModel
	Name   string `gorm:"not null;index:name" json:"name"`
	Desc   string `gorm:"not null" json:"desc"`
	Method string `gorm:"not null" json:"method"`
	Status bool   `json:"status"`
}
