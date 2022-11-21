package model

import "time"

type FamilyMember struct {
	FamilyID   uint     `gorm:"primaryKey;autoIncrement"`
	FamilyAs   FamilyAs `gorm:"foreignKey:FamilyAsId"`
	FamilyAsId uint
	Name       string    `gorm:"size:50;not null"`
	Nik        string    `gorm:"size:16;not null"`
	CreatedAT  time.Time `json:"created_at"`
	UpdatedAT  time.Time `json:"updated_at"`
}
