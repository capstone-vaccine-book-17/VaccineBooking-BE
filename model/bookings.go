package model

import "time"

type Booking struct {
	BookingID uint    `gorm:"primaryKey;autoIncrement"`
	Citizen   Citizen `gorm:"foreignKey:CitizenId"`
	CitizenId uint
	Session   Session `gorm:"foreignKey:SessionId"`
	SessionId uint
	CreatedAT time.Time `json:"created_at"`
}
