package model

import "time"

type Booking struct {
	BookingID uint      `gorm:"primaryKey;autoIncrement" json:"booking_id"`
	Citizen   Citizen   `gorm:"foreignKey:CitizenId"`
	CitizenId uint      `json:"citizen_id"`
	Session   Session   `gorm:"foreignKey:SessionId"`
	SessionId uint      `json:"session_id"`
	CreatedAT time.Time `json:"created_at"`
	Queue     string    `json:"queue"`
}
