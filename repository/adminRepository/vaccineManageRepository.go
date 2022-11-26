package adminRepository

import (
	"capstone_vaccine/dto/adminDto"
	"capstone_vaccine/model"
	"time"
)

func (u *adminRepository) CreateVaccine(input adminDto.VaccineRequest) (adminDto.VaccineDTO, error) {

	temp := adminDto.VaccineDTO{
		Name:    input.Name,
		Kuota:   input.Kuota,
		Expired: input.Expired,
	}

	if err := u.db.Create(&model.VaccineVarietie{
		MedicalFacilitysId: input.MedicalFacilitysId,
		Name:               input.Name,
		Kuota:              input.Kuota,
		Expired:            input.Expired,
		CreatedAT:          time.Now(),
		UpdatedAT:          time.Now(),
	}).Error; err != nil {
		return temp, err
	}

	return temp, nil
}
