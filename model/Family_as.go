package model

import "time"

type FamilyAs struct {
	FamilyAsID uint      `gorm:"primaryKey;autoIncrement" json:"family_as_id"`
	Name       string    `gorm:"size:30;not null" json:"name"`
	CreatedAT  time.Time `json:"created_at"`
}
