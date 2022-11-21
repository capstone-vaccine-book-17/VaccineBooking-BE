package model

import "time"

type Address struct {
	AddressID uint      `gorm:"primaryKey;autoIncrement"`
	Address   string    `gorm:"size:200;not null"`
	Province  string    `gorm:"size:50"`
	PostCode  int       `gorm:"size:10"`
	Country   string    `gorm:"size:50"`
	City      string    `gorm:"size:50"`
	CreatedAT time.Time `json:"created_at"`
	UpdatedAT time.Time `json:"updated_at"`
}
