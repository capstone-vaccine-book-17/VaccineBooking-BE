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

// TODO GET Vaccine BY ID
func (u *adminRepository) GetVaccineById(payloads adminDto.VaccineRequest) (adminDto.VaccineDTO, error) {
	vaccine := adminDto.VaccineDTO{}

	if err := u.db.Model(&model.VaccineVarietie{}).Where("vaccine_id = ?", payloads.VaccineID).Find(&vaccine).Error; err != nil {
		return vaccine, err
	}

	return vaccine, nil
}

func (u *adminRepository) UpdateVaccine(payloads adminDto.VaccineRequest) (adminDto.VaccineRequest, error) {

	temp := adminDto.VaccineRequest{
		VaccineID: payloads.VaccineID,
		Name:      payloads.Name,
		Kuota:     payloads.Kuota,
		Expired:   payloads.Expired,
	}

	if err := u.db.Model(&model.VaccineVarietie{}).Where("vaccine_id = ?", payloads.VaccineID).Updates(&model.VaccineVarietie{
		Name:      payloads.Name,
		Kuota:     payloads.Kuota,
		Expired:   payloads.Expired,
		UpdatedAT: time.Now(),
	}).Error; err != nil {
		return temp, err
	}

	return temp, nil
}

// // TODO DELETE SESSION
// func (u *adminRepository) DeleteSession(payloads adminDto.SessionWithStatusDTO) error {
// 	if err := u.db.Where("session_id", payloads.SessionId).Delete(&model.Session{}).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (u *adminRepository) UpdateVaccine(updateReq adminDto.VaccineRequest) (adminDto.VaccineDTO, error) {
// 	temp := adminDto.VaccineDTO{
// 		Name:    updateReq.Name,
// 		Kuota:   updateReq.Kuota,
// 		Expired: updateReq.Expired,
// 	}

// 	if err := u.db.Model(&model.VaccineVarietie{}).Where("vaccine_id = ?", updateReq.VaccineID).Updates(&model.VaccineVarietie{
// 		Name:      updateReq.Name,
// 		Kuota:     updateReq.Kuota,
// 		Expired:   updateReq.Expired,
// 		UpdatedAT: time.Now(),
// 	}).Error; err != nil {
// 		return temp, err
// 	}

// 	return temp, nil

// }

// TODO DELETE KANDIDAT
func (u adminRepository) DeleteVaccine(data adminDto.VaccineRequest) ([]model.VaccineVarietie, error) {
	vaccine := []model.VaccineVarietie{}

	qry := u.db.Where("vaccine_id = ?", data.VaccineID).Delete(&vaccine)

	if qry.Error != nil {
		return nil, qry.Error
	}

	if err := u.db.Model(&model.VaccineVarietie{}).Find(&vaccine).Error; err != nil {
		return nil, err
	}

	return vaccine, nil
}
