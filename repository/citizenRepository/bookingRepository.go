package citizenRepository

import (
	"capstone_vaccine/dto/citizenDto"
	"capstone_vaccine/model"
	"time"
)

// TODO GET MAX QUEUE
func (u *citizenRepository) GetMaxQueue(session_id uint) (citizenDto.MaxQueue, error) {
	var booking citizenDto.MaxQueue
	if err := u.db.Raw("SELECT max(queue) as total_q FROM `bookings` WHERE session_id = ?", session_id).Scan(&booking).Error; err != nil {
		return booking, err
	}

	return booking, nil
}

// TODO UPDATE KUOTA SESSION
func (u *citizenRepository) UpdateKuotaSession(session_id uint, kuota string) error {
	if err := u.db.Model(&model.Session{}).Where("session_id = ?", session_id).Updates(&model.Session{
		Kuota:     kuota,
		UpdatedAT: time.Now(),
	}).Error; err != nil {
		return err
	}

	return nil
}

// TODO CREATE BOOKING
func (u *citizenRepository) CreateBooking(payloads citizenDto.BookingDto) (citizenDto.BookingDto, error) {
	if err := u.db.Create(&model.Booking{
		CitizenId: payloads.CitizenID,
		SessionId: payloads.SessionID,
		Queue:     payloads.Queue,
		Status:    "process",
		CreatedAT: time.Now(),
		UpdatedAT: time.Now(),
	}).Error; err != nil {
		return payloads, err
	}

	return payloads, nil
}

// TODO GET LAST BOOKING BY CITIZEN
func (u *citizenRepository) GetLastBooking(citizenId uint) (citizenDto.TicketBooking, error) {

	var ticket citizenDto.TicketBooking
	if err := u.db.Model(&model.Booking{}).Select("sessions.*, citizens.name, citizens.nik, citizens.gender, vaccine_varieties.name as vaccine, medical_facilitys.name as rs_name, bookings.queue").
		Joins("JOIN sessions ON sessions.session_id = bookings.session_id").
		Joins("JOIN citizens ON citizens.citizen_id = bookings.citizen_id").
		Joins("JOIN vaccine_varieties ON vaccine_varieties.vaccine_id = sessions.vaccine_id").
		Joins("JOIN medical_facilitys ON medical_facilitys.medical_facilitys_id = sessions.medical_facilitys_id").
		Where("bookings.citizen_id = ?", citizenId).
		Order("bookings.created_at").
		Limit(1).
		Find(&ticket).Error; err != nil {
		return ticket, err
	}

	return ticket, nil
}
