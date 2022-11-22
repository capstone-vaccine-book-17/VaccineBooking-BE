package model

import "time"

type Role struct {
	RoleID    uint      `gorm:"primaryKey;autoIncrement" json:"role_id"`
	Name      string    `gorm:"size:50;not null" json:"name"`
	CreatedAT time.Time `json:"created_at"`
}
