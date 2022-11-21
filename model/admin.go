package model

import (
	"time"
)

type Admin struct {
	Role               Role             `gorm:"foreignKey:RoleId"`
	MedicalFacilitys   MedicalFacilitys `gorm:"foreignKey:MedicalFacilitysId"`
	AdminID            uint             `gorm:"primaryKey;autoIncrement"`
	RoleId             uint
	MedicalFacilitysId uint
	Username           string    `gorm:"size:50;not null"`
	Password           string    `gorm:"size:50;not null"`
	CreatedAT          time.Time `json:"created_at"`
}
