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

	err := u.db.Create(&address).Error
	if err != nil {
		return payloads, err
	}

	medical := model.MedicalFacilitys{
		Name:      payloads.Name,
		AddressId: address.AddressID,
		NoTlp:     payloads.NoTlp,
		CreatedAT: time.Now(),
		UpdatedAT: time.Now(),
	}

	err = u.db.Create(&medical).Error
	if err != nil {
		return payloads, err
	}

	return payloads, nil

}
