package model

import "time"

type Session struct {
	SessionID          uint             `gorm:"primaryKey;autoIncrement" json:"session_id"`
	MedicalFacilitys   MedicalFacilitys `gorm:"foreignKey:MedicalFacilitysId"`
	MedicalFacilitysId uint             `json:"medical_facilitys_id"`
	VaccineVarietie    VaccineVarietie  `gorm:"foreignKey:VaccineId"`
	VaccineId          uint             `json:"vaccine_id"`
	Name               string           `gorm:"size:50;not null" json:"name"`
	Kuota              string           `gorm:"size:10;not null" json:"kuota"`
	Dosis              string           `gorm:"size:30;not null" json:"dosis"`
	Date               string           `gorm:"size:30;not null" json:"date"`
	StartTime          string           `gorm:"size:20;not null" json:"startTime"`
	EndTime            string           `gorm:"size:20;not null" json:"endTime"`
	Status             string           `gorm:"size:20;not null" json:"status"`
	CreatedAT          time.Time        `json:"created_at"`
	UpdatedAT          time.Time        `json:"updated_at"`
}
