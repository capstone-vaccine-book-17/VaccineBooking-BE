package citizenRepository

import (
	"capstone_vaccine/dto/citizenDto"
	"capstone_vaccine/model"
	"time"
)

// TODO GET Profile
func (u *citizenRepository) GetProfile(payloads citizenDto.ProfileReq) (citizenDto.ProfileDTO, error) {
	profile := citizenDto.ProfileDTO{}

	if err := u.db.Model(&model.Citizen{}).Select("citizens.Name,nik,image").Where("citizen_id = ?", payloads.CitizenID).Find(&profile).Error; err != nil {
		return profile, err
	}
	return profile, nil
}

// TODO Upload Image
func (u *citizenRepository) UploadImage(payloads citizenDto.ProfileReq) (citizenDto.ProfileReq, error) {
	UpImage := citizenDto.ProfileReq{}

	if err := u.db.Model(&model.Citizen{}).Where("citizen_id = ?", payloads.CitizenID).Updates(&model.Citizen{
		Image:     payloads.Image,
		UpdatedAT: time.Now(),
	}).Error; err != nil {
		return UpImage, err
	}

	return UpImage, nil

}
