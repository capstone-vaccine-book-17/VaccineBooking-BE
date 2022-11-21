package model

import "time"

type VaccineVarietie struct {
	VaccineID          uint             `gorm:"primaryKey;autoIncrement"`
	MedicalFacilitys   MedicalFacilitys `gorm:"foreignKey:MedicalFacilitysId"`
	MedicalFacilitysId uint
	Name               string    `gorm:"size:50;not null"`
	Kuota              int       `gorm:"size:10;not null"`
	Expired            string    `gorm:"size:50;not null"`
	CreatedAT          time.Time `json:"created_at"`
	UpdatedAT          time.Time `json:"updated_at"`
}
