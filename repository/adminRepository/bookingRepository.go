package adminRepository

import (
	"capstone_vaccine/dto/adminDto"
	"capstone_vaccine/model"
	"strconv"
	"time"
)

// TODO GET MAX QUEUE
func (u *adminRepository) GetMaxQueue(session_id uint) (adminDto.MaxQueue, error) {
	var booking adminDto.MaxQueue
	if err := u.db.Raw("SELECT max(queue) as total_q FROM `bookings` WHERE session_id = ?", session_id).Scan(&booking).Error; err != nil {
		return booking, err
	}

	return booking, nil
}

// TODO CREATE CITIZEN FOR BOOKING
func (u *adminRepository) CreateCitizenBook(nik, nama, address string) (model.Citizen, error) {
	var dto model.Citizen
	citizen_address := model.Address{
		Address:   address,
		CreatedAT: time.Now(),
		UpdatedAT: time.Now(),
	}

	result1 := u.db.Create(&citizen_address)
	if result1.Error != nil {
		return dto, result1.Error
	}

	citizen := model.Citizen{
		AddressId: citizen_address.AddressID,
		Name:      nama,
		Nik:       nik,
		CreatedAT: time.Now(),
		UpdatedAT: time.Now(),
	}

	result2 := u.db.Create(&citizen)
	if result2.Error != nil {
		return citizen, result2.Error
	}

	return citizen, nil
}

// TODO CREATE BOOKING
func (u *adminRepository) CreateBooking(payloads adminDto.BookingDto) (adminDto.BookingDto, error) {
	if err := u.db.Create(&model.Booking{
		CitizenId: payloads.CitizenId,
		SessionId: payloads.SessionId,
		Queue:     payloads.Queue,
		Status:    "process",
		CreatedAT: time.Now(),
		UpdatedAT: time.Now(),
	}).Error; err != nil {
		return payloads, err
	}

	return payloads, nil
}

// TODO UPDATE KUOTA SESSION
func (u *adminRepository) UpdateSessionBooking(session_id uint, kuota string) error {
	if err := u.db.Model(&model.Session{}).Where("session_id = ?", session_id).Updates(&model.Session{
		Kuota:     kuota,
		UpdatedAT: time.Now(),
	}).Error; err != nil {
		return err
	}

	return nil
}

// TODO UPDATE BOOKING
func (u *adminRepository) UpdateBooking(payloads adminDto.UpdateBooking) (adminDto.UpdateBooking, error) {
	if payloads.Status == "process" {
		if err := u.db.Model(&model.Booking{}).Where("booking_id = ?", payloads.BookingId).Updates(&model.Booking{
			Status:    "process",
			UpdatedAT: time.Now(),
		}).Error; err != nil {
			return payloads, err
		}
	} else if payloads.Status == "batal" {
		var booking model.Booking
		if err := u.db.Model(&model.Booking{}).Where("booking_id = ? AND status='process'", payloads.BookingId).Find(&booking).Error; err != nil {
			return payloads, err
		}
		sessionID := adminDto.SessionWithStatusDTO{
			SessionId: booking.SessionId,
		}
		kuota, err := u.GetSessionById(sessionID)

		if err != nil {
			return payloads, err
		}
		addition := kuota.Kuota + 1
		conv_addition := strconv.Itoa(addition)
		if err := u.db.Model(&model.Session{}).Where("session_id = ?", booking.SessionId).Updates(&model.Session{
			Kuota:     conv_addition,
			UpdatedAT: time.Now(),
		}).Error; err != nil {
			return payloads, err
		}

		if err := u.db.Model(&model.Booking{}).Where("booking_id = ?", payloads.BookingId).Updates(&model.Booking{
			Status:    "batal",
			UpdatedAT: time.Now(),
		}).Error; err != nil {
			return payloads, err
		}
	} else if payloads.Status == "selesai" {
		if err := u.db.Model(&model.Booking{}).Where("booking_id = ?", payloads.BookingId).Updates(&model.Booking{
			Status:    "selesai",
			UpdatedAT: time.Now(),
		}).Error; err != nil {
			return payloads, err
		}
	}

	return payloads, nil
}

// TODO GET ALL BOOKING
func (u *adminRepository) GetAllBooking(medicalId uint) ([]adminDto.BookingAllDto, error) {

	booking := []adminDto.BookingAllDto{}

	if err := u.db.Model(&model.Booking{}).Select("citizens.name as citizen_name,citizens.nik,sessions.dosis,sessions.date,sessions.start_time,sessions.end_time,bookings.queue, bookings.booking_id,bookings.status").
		Joins("join citizens on citizens.citizen_id = bookings.citizen_id").Joins("join sessions on sessions.session_id = bookings.session_id").Where("sessions.medical_facilitys_id = ?", medicalId).Find(&booking).Error; err != nil {
		return nil, err
	}

	return booking, nil
}

// TODO GET BOOKING BY ID
func (u *adminRepository) GetBookingById(payloads adminDto.BookingAllDto) (adminDto.BookingAllDto, error) {

	var booking adminDto.BookingAllDto

	if err := u.db.Model(&model.Booking{}).Select("citizens.name as citizen_name,citizens.nik,sessions.dosis,sessions.date,sessions.start_time,sessions.end_time,bookings.queue, bookings.booking_id,bookings.status").
		Where("bookings.booking_id = ?", payloads.BookingId).
		Joins("join citizens on citizens.citizen_id = bookings.citizen_id").Joins("join sessions on sessions.session_id = bookings.session_id").Find(&booking).Error; err != nil {
		return booking, err
	}

	return booking, nil
}

// TODO DELETE BOOKING
func (u *adminRepository) DeleteBooking(payloads adminDto.BookingAllDto) error {
	var booking model.Booking
	if err := u.db.Model(&model.Booking{}).Where("booking_id = ? AND status='process'", payloads.BookingId).Find(&booking).Error; err != nil {
		return err
	}
	sessionID := adminDto.SessionWithStatusDTO{
		SessionId: booking.SessionId,
	}
	kuota, err := u.GetSessionById(sessionID)

	if err != nil {
		return err
	}
	addition := kuota.Kuota + 1
	conv_addition := strconv.Itoa(addition)
	if err := u.db.Model(&model.Session{}).Where("session_id = ?", booking.SessionId).Updates(&model.Session{
		Kuota:     conv_addition,
		UpdatedAT: time.Now(),
	}).Error; err != nil {
		return err
	}

	if err := u.db.Where("booking_id", payloads.BookingId).Delete(&model.Booking{}).Error; err != nil {
		return err
	}

	return nil
}
