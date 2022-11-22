package model

import (
	"time"
)

type Admin struct {
	Role               Role             `gorm:"foreignKey:RoleId"`
	MedicalFacilitys   MedicalFacilitys `gorm:"foreignKey:MedicalFacilitysId"`
	AdminID            uint             `gorm:"primaryKey;autoIncrement" json:"admin_id"`
	RoleId             uint             `json:"role_id"`
	MedicalFacilitysId uint             `json:"medical_facilitys_id"`
	Username           string           `gorm:"size:50;not null" json:"username"`
	Password           string           `gorm:"size:50;not null" json:"password"`
	CreatedAT          time.Time        `json:"created_at"`
}
