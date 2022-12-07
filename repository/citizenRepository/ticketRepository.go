package citizenRepository

import (
	"capstone_vaccine/dto/citizenDto"
	"capstone_vaccine/model"
)

func (u *citizenRepository) GetTicket(citizenId uint) ([]citizenDto.TicketDetails, error) {

	var ticket []citizenDto.TicketDetails

	// dik := "proses"
	if err := u.db.Model(&model.Booking{}).Select("bookings.*, citizens.name, citizens.nik, citizens.gender, vaccine_varieties.name as vaccine, medical_facilitys.name as rs_name,sessions.dosis,sessions.date,sessions.start_time,sessions.end_time").
		Joins("JOIN sessions ON sessions.session_id = bookings.session_id").
		Joins("JOIN citizens ON citizens.citizen_id = bookings.citizen_id").
		Joins("JOIN vaccine_varieties ON vaccine_varieties.vaccine_id = sessions.vaccine_id").
		Joins("JOIN medical_facilitys ON medical_facilitys.medical_facilitys_id = sessions.medical_facilitys_id").
		Where("bookings.citizen_id = ?", citizenId).
		Find(&ticket).Error; err != nil {
		return ticket, err
	}

	return ticket, nil
}
