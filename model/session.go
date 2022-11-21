package model

import "time"

type Session struct {
	SessionID          uint             `gorm:"primaryKey;autoIncrement"`
	MedicalFacilitys   MedicalFacilitys `gorm:"foreignKey:MedicalFacilitysId"`
	MedicalFacilitysId uint
	VaccineVarietie    VaccineVarietie `gorm:"foreignKey:VaccineId"`
	VaccineId          uint
	Name               string    `gorm:"size:50;not null"`
	Kuota              int       `gorm:"size:10;not null"`
	Dosis              string    `gorm:"size:30;not null"`
	Date               string    `gorm:"size:30;not null"`
	StartTime          string    `gorm:"size:20;not null"`
	EndTime            string    `gorm:"size:20;not null"`
	Status             string    `gorm:"size:20;not null"`
	CreatedAT          time.Time `json:"created_at"`
	UpdatedAT          time.Time `json:"updated_at"`
}
