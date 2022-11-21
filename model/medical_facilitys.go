package model

import "time"

type MedicalFacilitys struct {
	Address            Address `gorm:"foreignKey:AddressId"`
	MedicalFacilitysID uint    `gorm:"primaryKey;autoIncrement" json:"medical_facilitys_id"`
	Name               string  `gorm:"size:125;not null" form:"name" json:"name"`
	AddressId          uint
	NoTlp              string `gorm:"size:14;not null"`
	Image              string
	CreatedAT          time.Time `json:"created_at"`
	UpdatedAT          time.Time `json:"updated_at"`
}
