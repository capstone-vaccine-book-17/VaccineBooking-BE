package model

import "time"

type Address struct {
	AddressID uint      `gorm:"primaryKey;autoIncrement" json:"address_id"`
	Address   string    `gorm:"size:200;not null" json:"address"`
	Province  string    `gorm:"size:50" json:"province"`
	PostCode  int       `gorm:"size:10" json:"post_code"`
	Country   string    `gorm:"size:50" json:"country"`
	City      string    `gorm:"size:50" json:"city"`
	CreatedAT time.Time `json:"created_at"`
	UpdatedAT time.Time `json:"updated_at"`
}
