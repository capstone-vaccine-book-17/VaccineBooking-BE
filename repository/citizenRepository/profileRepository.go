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

	res := citizenDto.ProfileReq{
		Image: payloads.Image,
	}

	if err := u.db.Model(&model.Citizen{}).Where("citizen_id = ?", payloads.CitizenID).Updates(&model.Citizen{
		Image:     payloads.Image,
		UpdatedAT: time.Now(),
	}).Error; err != nil {
		return res, err

	}
	return res, nil

}

func (u *citizenRepository) GetPersonalData(payload citizenDto.ProfileReq) ([]citizenDto.PersonalData, error) {

	Personal := []citizenDto.PersonalData{}

	if err := u.db.Model(&model.Citizen{}).Select("citizens.*, addresses.*").Joins("join addresses on addresses.address_id= citizens.address_id").Where("citizen_id=?", payload.CitizenID).Find(&Personal).Error; err != nil {
		return nil, err
	}
	return Personal, nil
}

func (u *citizenRepository) UpdateAddress(payloads citizenDto.AddressCitizenReq) error {

	temp := citizenDto.AddressCitizenReq{}
	temp2 := citizenDto.AddressCitizenReq{}

	if errh := u.db.Model(&model.Citizen{}).Select("citizens.*").Where("citizen_id=?", payloads.CitizenID).Find(&temp).Error; errh != nil {
		return errh
	}

	if erre := u.db.Model(&model.Address{}).Select("addresses.*").Where("address_id=?", temp.AddressID).Find(&temp2).Error; erre != nil {
		return erre
	}
	if err := u.db.Model(&model.Address{}).Where("address_id = ?", temp.AddressID).Updates(&model.Address{
		Address:   payloads.NewAddress,
		Province:  payloads.Province,
		City:      payloads.City,
		PostCode:  payloads.PostCode,
		UpdatedAT: time.Now(),
	}).Error; err != nil {
		return err

	}
	return nil
}

func (u *citizenRepository) GetAddress(payload citizenDto.ProfileReq) (citizenDto.AddressResp, error) {
	temp := citizenDto.AddressCitizenReq{}
	out := citizenDto.AddressResp{}

	if err := u.db.Model(&model.Citizen{}).Select("citizens.*").Where("citizen_id=?", payload.CitizenID).Find(&temp).Error; err != nil {
		return out, err
	}

	if err := u.db.Model(&model.Address{}).Select("addresses.*").Where("address_id=?", temp.AddressID).Find(&out).Error; err != nil {
		return out, err
	}

	return out, nil
}

func (u *citizenRepository) GetEmail(payloads citizenDto.ProfileReq) (citizenDto.PersonalData, error) {
	profile := citizenDto.PersonalData{}

	if err := u.db.Model(&model.Citizen{}).Select("citizens.email").Where("citizen_id = ?", payloads.CitizenID).Find(&profile).Error; err != nil {
		return profile, err
	}
	return profile, nil
}

func (u *citizenRepository) UpdateEmail(payloads citizenDto.UpdateEmail) error {

	if err := u.db.Model(&model.Citizen{}).Where("citizen_id = ?", payloads.CitizenID).Updates(&model.Citizen{
		Email: payloads.Email,
	}).Error; err != nil {
		return err

	}
	return nil
}
