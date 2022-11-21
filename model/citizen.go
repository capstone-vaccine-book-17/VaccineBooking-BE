package model

import "time"

type Citizen struct {
	CitizenID uint    `gorm:"primaryKey;autoIncrement"`
	Address   Address `gorm:"foreignKey:AddressId"`
	AddressId uint
	Name      string    `gorm:"size:50;not null"`
	Nik       string    `gorm:"size:16;not null"`
	Dob       string    `gorm:"size:50;not null"`
	Gender    string    `gorm:"size:15;not null"`
	Email     string    `gorm:"size:50;not null"`
	Password  string    `gorm:"not null"`
	Image     string    `gorm:"size:50;"`
	CreatedAT time.Time `json:"created_at"`
	UpdatedAT time.Time `json:"updated_at"`
}
