package adminRepository

import (
	"capstone_vaccine/dto/adminDto"
	"capstone_vaccine/model"
)

// TODO Get Profile
func (u *adminRepository) GetProfile(payloads adminDto.ProfileRequest) ([]adminDto.ProfilDTO, error) {
	profile := []adminDto.ProfilDTO{}

	if err := u.db.Model(&model.Admin{}).Select("admins.*, medical_facilitys.name as name,medical_facilitys.image as image,addresses.address as address").Joins("join medical_facilitys on medical_facilitys.medical_facilitys_id= admins.medical_facilitys_id").Where("admin_id=?", payloads.AdminID).Joins("join addresses on addresses.address_id=medical_facilitys.address_id").Find(&profile).Error; err != nil {
		return nil, err
	}

	return profile, nil
}

func (u *adminRepository) UpdateProfile(payloads adminDto.ProfileRequest) (adminDto.ProfileRequest, error) {

	temp := adminDto.ProfileRequest{
		AdminID:            payloads.AdminID,
		MedicalFacilitysId: payloads.MedicalFacilitysId,
		Name:               payloads.Name,
		Image:              payloads.Image,
		Address:            payloads.Address,
		Username:           payloads.Username,
	}
	profile := adminDto.Address{}
	if err := u.db.Model(&model.MedicalFacilitys{}).Select("medical_facilitys.*").Where("medical_facilitys_id = ?", payloads.MedicalFacilitysId).Find(&profile).Error; err != nil {
		return temp, err
	}

	if err := u.db.Model(&model.Admin{}).Where("admin_id = ?", payloads.AdminID).Updates(&model.Admin{
		Username: temp.Username,
	}).Error; err != nil {
		return temp, err

	}
	if err := u.db.Model(&model.MedicalFacilitys{}).Where("medical_facilitys_id = ?", payloads.MedicalFacilitysId).Updates(&model.MedicalFacilitys{
		Name:  temp.Name,
		Image: payloads.Image,
	}).Error; err != nil {
		return temp, err

	}
	if err := u.db.Model(&model.Address{}).Where("address_id = ?", profile.AddressID).Updates(&model.Address{
		Address: payloads.Address,
	}).Error; err != nil {
		return temp, err

	}

	return temp, nil
}
