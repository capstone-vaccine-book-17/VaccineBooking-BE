package adminRepository

import (
	"capstone_vaccine/dto/adminDto"
	"capstone_vaccine/model"
)

// TODO Get Profile
func (u *adminRepository) GetProfile() ([]adminDto.ProfilDTO, error) {
	profile := []adminDto.ProfilDTO{}

	if err := u.db.Model(&model.Admin{}).Select("admins.*, medical_facilitys.name as name,medical_facilitys.image as image,addresses.address as address").Joins("join medical_facilitys on medical_facilitys.medical_facilitys_id= admins.medical_facilitys_id").Joins("join addresses on addresses.address_id=medical_facilitys.address_id").Find(&profile).Error; err != nil {
		return nil, err
	}

	return profile, nil
}
