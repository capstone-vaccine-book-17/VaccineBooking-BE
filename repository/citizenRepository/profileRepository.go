package citizenRepository

import (
	"capstone_vaccine/dto/citizenDto"
	"capstone_vaccine/model"
)

// TODO GET Profile
func (u *citizenRepository) GetProfile(payloads citizenDto.ProfileReq) (citizenDto.ProfileDTO, error) {
	profile := citizenDto.ProfileDTO{}

	if err := u.db.Model(&model.Citizen{}).Select("citizens.Name,nik").Where("citizen_id = ?", payloads.CitizenID).Find(&profile).Error; err != nil {
		return profile, err
	}
	return profile, nil
}
