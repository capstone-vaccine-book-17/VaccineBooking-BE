package adminRepository

import (
	"capstone_vaccine/dto/adminDto"
	"capstone_vaccine/model"
	"time"
)

func (u *adminRepository) CreateMedical(payloads adminDto.MedicalDto) (adminDto.MedicalDto, error) {

	address := model.Address{
		Address:   payloads.Address,
		Province:  payloads.Province,
		PostCode:  payloads.PostCode,
		Country:   payloads.Country,
		City:      payloads.City,
		CreatedAT: time.Now(),
		UpdatedAT: time.Now(),
	}

	errA := u.db.Create(&address).Error
	if errA != nil {
		return payloads, errA
	}

	medical := model.MedicalFacilitys{
		Name:      payloads.Name,
		AddressId: address.AddressID,
		NoTlp:     payloads.NoTlp,
		CreatedAT: time.Now(),
		UpdatedAT: time.Now(),
	}

	errM := u.db.Create(&medical).Error
	if errM != nil {
		return payloads, errM
	}

	return payloads, nil

}
