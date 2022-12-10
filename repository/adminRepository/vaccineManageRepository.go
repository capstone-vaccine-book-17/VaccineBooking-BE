package adminRepository

// // TODO CREATE Vaccine
// func (u *adminRepository) CreateVaccine(input adminDto.VaccineRequest) (adminDto.VaccineResponse, error) {

// 	temp := adminDto.VaccineResponse{
// 		Name:    input.Name,
// 		Kuota:   input.Kuota,
// 		Expired: input.Expired,
// 	}

// 	if err := u.db.Create(&model.VaccineVarietie{
// 		MedicalFacilitysId: input.MedicalFacilitysId,
// 		Name:               input.Name,
// 		Kuota:              input.Kuota,
// 		Expired:            input.Expired,
// 		CreatedAT:          time.Now(),
// 		UpdatedAT:          time.Now(),
// 	}).Error; err != nil {
// 		return temp, err
// 	}

// 	return temp, nil
// }

// // TODO View All Vaccine
// func (u *adminRepository) ViewAllVaccine() ([]adminDto.VaccineDTO, error) {
// 	vaccine := []adminDto.VaccineDTO{}

// 	if err := u.db.Model(&model.VaccineVarietie{}).Find(&vaccine).Error; err != nil {
// 		return nil, err
// 	}

// 	return vaccine, nil
// }

// // TODO Update Vaccine
// func (u *adminRepository) UpdateVaccine(payloads adminDto.VaccineDTO) (adminDto.VaccineDTO, error) {

// 	temp := adminDto.VaccineDTO{
// 		VaccineID: payloads.VaccineID,
// 		Name:      payloads.Name,
// 		Kuota:     payloads.Kuota,
// 		Expired:   payloads.Expired,
// 	}

// 	if err := u.db.Model(&model.VaccineVarietie{}).Where("vaccine_id = ?", payloads.VaccineID).Updates(&model.VaccineVarietie{
// 		Name:      payloads.Name,
// 		Kuota:     payloads.Kuota,
// 		Expired:   payloads.Expired,
// 		UpdatedAT: time.Now(),
// 	}).Error; err != nil {
// 		return temp, err
// 	}

// 	return temp, nil
// }

// // TODO DELETE Vaccine
// func (u adminRepository) DeleteVaccine(data adminDto.VaccineDTO) error {

// 	if err := u.db.Where("vaccine_id = ?", data.VaccineID).Delete(&model.VaccineVarietie{}).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }
