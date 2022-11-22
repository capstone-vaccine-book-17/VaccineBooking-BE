package model

import "time"

type FamilyMember struct {
	FamilyID   uint      `gorm:"primaryKey;autoIncrement" json:"family_id"`
	FamilyAs   FamilyAs  `gorm:"foreignKey:FamilyAsId"`
	FamilyAsId uint      `json:"family_as_id"`
	Name       string    `gorm:"size:50;not null" json:"name"`
	Nik        string    `gorm:"size:16;not null" json:"nik"`
	CreatedAT  time.Time `json:"created_at"`
	UpdatedAT  time.Time `json:"updated_at"`
}
