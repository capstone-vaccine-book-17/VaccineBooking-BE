package model

import "time"

type FamilyAs struct {
	FamilyAsID uint      `gorm:"primaryKey;autoIncrement"`
	Name       string    `gorm:"size:30;not null"`
	CreatedAT  time.Time `json:"created_at"`
}
