package model

import "time"

type FamilyMember struct {
	FamilyID  uint      `gorm:"primaryKey;autoIncrement" json:"family_id"`
	Citizen   Citizen   `gorm:"foreignKey:CitizenId"`
	CitizenId uint      `json:"citizen_id"`
	FamilyAs  string    `gorm:"size:15;not null" json:"family_as"`
	Name      string    `gorm:"size:50;not null" json:"name"`
	Nik       string    `gorm:"size:16;not null" json:"nik"`
	Age       uint      `gorm:"age"`
	Gender    string    `gorm:"size:15;not null" json:"gender"`
	CreatedAT time.Time `json:"created_at"`
	UpdatedAT time.Time `json:"updated_at"`
}
