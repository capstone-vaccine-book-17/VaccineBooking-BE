package citizenRepository

import (
	"capstone_vaccine/dto/citizenDto"
	"capstone_vaccine/model"
)

// TODO GET SESSION BY MEDICAL FACILITYS ID
func (u *citizenRepository) GetSessionByMedicalId(medicalID uint) ([]citizenDto.SessionDto, error) {

	var session []citizenDto.SessionDto
	if err := u.db.Model(&model.Session{}).Select("sessions.session_id, sessions.name, sessions.kuota, sessions.dosis, sessions.date, DATE_FORMAT(sessions.date, '%d %M %Y') as conv_date, sessions.start_time, sessions.end_time, vaccine_varieties.name as vaccine").
		Joins("JOIN vaccine_varieties on vaccine_varieties.vaccine_id = sessions.vaccine_id").
		Where("sessions.medical_facilitys_id = ? AND sessions.status = 'process'", medicalID).Find(&session).Error; err != nil {
		return nil, err
	}

	return session, nil
}

// TODO GET SESSION BY ID
func (u *citizenRepository) GetSessionById(id uint) (citizenDto.SessionWithVaccineId, error) {
	var session citizenDto.SessionWithVaccineId

	if err := u.db.Model(&model.Session{}).Where("session_id = ?", id).Find(&session).Error; err != nil {
		return session, err
	}

	return session, nil
}
