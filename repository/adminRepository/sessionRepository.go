package adminRepository

// TODO SESSION

// // TODO SUM ALL KUOTA FROM SESSION
// func (u *adminRepository) CountKuota(vaccineID uint) (adminDto.CountKuota, error) {
// 	var result adminDto.CountKuota
// 	var (
// 		totals int
// 		totalv int
// 	)

// 	if err := u.db.Model(&model.Session{}).Select("coalesce(sum(kuota), 0) as total_s").Where("vaccine_id = ? AND sessions.status = 'process'", vaccineID).Find(&totals).Error; err != nil {
// 		return result, err
// 	}

// 	if err := u.db.Model(&model.VaccineVarietie{}).Select("coalesce(kuota, 0) as total_v").Where("vaccine_id = ?", vaccineID).Find(&totalv).Error; err != nil {
// 		return result, err
// 	}

// 	result.TotalS = totals
// 	result.TotalV = totalv

// 	return result, nil
// }

// // TODO AUTOMATIC UPDATE STATUS SESSION
// func (u *adminRepository) AutoUpdateSession(dateR, timeR string) error {

// 	if err := u.db.Model(&model.Session{}).Where("date = ? AND end_time = ? AND status='process'", dateR, timeR).Updates(&model.Session{
// 		Status:    "selesai",
// 		UpdatedAT: time.Now(),
// 	}).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }

// // TODO CREATE SESSION
// func (u *adminRepository) CreateSession(payloads adminDto.SessionRequest) (adminDto.SessionDTO, error) {

// 	temp := adminDto.SessionDTO{
// 		Name:      payloads.Name,
// 		VaccineId: payloads.VaccineId,
// 		StartTime: payloads.StartTime,
// 		Kuota:     payloads.Kuota,
// 		Dosis:     payloads.Dosis,
// 		EndTime:   payloads.EndTime,
// 		Date:      payloads.Date,
// 	}
// 	convKuota := strconv.Itoa(payloads.Kuota)
// 	if err := u.db.Create(&model.Session{
// 		Name:               payloads.Name,
// 		MedicalFacilitysId: payloads.MedicalFacilitysId,
// 		VaccineId:          payloads.VaccineId,
// 		StartTime:          payloads.StartTime,
// 		Kuota:              convKuota,
// 		Dosis:              payloads.Dosis,
// 		EndTime:            payloads.EndTime,
// 		Date:               payloads.Date,
// 		Status:             "process",
// 		CreatedAT:          time.Now(),
// 		UpdatedAT:          time.Now(),
// 	}).Error; err != nil {
// 		return temp, err
// 	}

// 	return temp, nil
// }

// // TODO GET SESSION
// func (u *adminRepository) GetAllSession() ([]adminDto.SessionWithStatusDTO, error) {
// 	session := []adminDto.SessionWithStatusDTO{}

// 	if err := u.db.Model(&model.Session{}).Select("sessions.*, vaccine_varieties.name as vaccine_name").Joins("join vaccine_varieties on vaccine_varieties.vaccine_id = sessions.vaccine_id").Find(&session).Error; err != nil {
// 		return nil, err
// 	}

// 	return session, nil
// }

// // TODO GET SESSION BY ID
// func (u *adminRepository) GetSessionById(payloads adminDto.SessionWithStatusDTO) (adminDto.SessionWithStatusDTO, error) {
// 	session := adminDto.SessionWithStatusDTO{}

// 	if err := u.db.Model(&model.Session{}).Where("session_id = ?", payloads.SessionId).Select("sessions.*, vaccine_varieties.name as vaccine_name").Joins("join vaccine_varieties on vaccine_varieties.vaccine_id = sessions.vaccine_id").Find(&session).Error; err != nil {
// 		return session, err
// 	}

// 	return session, nil
// }

// // TODO UPDATE SESSION
// func (u *adminRepository) UpdateSession(payloads adminDto.SessionRequestUpdate) (adminDto.SessionRequestUpdate, error) {

// 	temp := adminDto.SessionRequestUpdate{
// 		SessionId: payloads.SessionId,
// 		Name:      payloads.Name,
// 		VaccineId: payloads.VaccineId,
// 		StartTime: payloads.StartTime,
// 		Kuota:     payloads.Kuota,
// 		Dosis:     payloads.Dosis,
// 		EndTime:   payloads.EndTime,
// 	}
// 	convKuota := strconv.Itoa(payloads.Kuota)
// 	if err := u.db.Model(&model.Session{}).Where("session_id = ?", payloads.SessionId).Updates(&model.Session{
// 		Name:      temp.Name,
// 		VaccineId: payloads.VaccineId,
// 		StartTime: payloads.StartTime,
// 		Kuota:     convKuota,
// 		Dosis:     payloads.Dosis,
// 		EndTime:   payloads.EndTime,
// 		UpdatedAT: time.Now(),
// 	}).Error; err != nil {
// 		return temp, err
// 	}

// 	return temp, nil
// }

// // TODO DELETE SESSION
// func (u *adminRepository) DeleteSession(payloads adminDto.SessionWithStatusDTO) error {
// 	if err := u.db.Where("session_id", payloads.SessionId).Delete(&model.Session{}).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }
