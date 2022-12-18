package model

import (
	"time"

	"gorm.io/gorm"
)

type VaccineVarietie struct {
	VaccineID          uint             `gorm:"primaryKey;autoIncrement" json:"vaccine_id"`
	MedicalFacilitys   MedicalFacilitys `gorm:"foreignKey:MedicalFacilitysId"`
	MedicalFacilitysId uint             `json:"medical_facilitys_id"`
	Name               string           `gorm:"size:50;not null" json:"name"`
	Kuota              int              `gorm:"size:10;not null" json:"kuota"`
	Expired            string           `gorm:"size:50;not null" json:"expired"`
	CreatedAT          time.Time        `json:"created_at"`
	UpdatedAT          time.Time        `json:"updated_at"`
	DeletedAT          gorm.DeletedAt   `gorm:"index"`
}
