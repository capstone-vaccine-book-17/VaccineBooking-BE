package model

import "time"

type Role struct {
	RoleID    uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"size:50;not null"`
	CreatedAT time.Time `json:"created_at"`
}
