package model

import "time"

type Citizen struct {
	CitizenID uint      `gorm:"primaryKey;autoIncrement" json:"citizen_id"`
	Address   Address   `gorm:"foreignKey:AddressId"`
	AddressId uint      `json:"address_id"`
	Name      string    `gorm:"size:50;not null" json:"name"`
	Nik       string    `gorm:"size:16;not null" json:"nik"`
	Age       uint      `gorm:"age"`
	Dob       string    `gorm:"size:50;not null" json:"dob"`
	Gender    string    `gorm:"size:15;not null" json:"gender"`
	Email     string    `gorm:"size:50;not null" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	Image     string    `json:"image"`
	CreatedAT time.Time `json:"created_at"`
	UpdatedAT time.Time `json:"updated_at"`
}
