package adminRepository

import (
	"capstone_vaccine/dto/adminDto"
	"capstone_vaccine/model"
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