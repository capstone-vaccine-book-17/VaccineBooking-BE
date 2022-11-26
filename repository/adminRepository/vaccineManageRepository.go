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

func (u *adminRepository) ViewAllVaccine() ([]adminDto.VaccineDTO, error) {
	vaccine := []adminDto.VaccineDTO{}

	if err := u.db.Model(&model.VaccineVarietie{}).Find(&vaccine).Error; err != nil {
		return nil, err
	}

	return vaccine, nil
}

func (u *adminRepository) UpdateVaccine(updateReq adminDto.VaccineRequest) (adminDto.VaccineDTO, error) {
	temp := adminDto.VaccineDTO{
		Name:    updateReq.Name,
		Kuota:   updateReq.Kuota,
		Expired: updateReq.Expired,
	}

	if err := u.db.Model(&model.VaccineVarietie{}).Where("vaccine_id = ?", updateReq.VaccineID).Updates(&model.VaccineVarietie{
		Name:      updateReq.Name,
		Kuota:     updateReq.Kuota,
		Expired:   updateReq.Expired,
		UpdatedAT: time.Now(),
	}).Error; err != nil {
		return temp, nil
	}

	return temp, nil

}
