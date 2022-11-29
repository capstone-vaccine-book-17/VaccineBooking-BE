package adminService

import (
	"capstone_vaccine/dto/adminDto"
	"errors"
	"strconv"
)

// TODO CREATE BOOKING
func (s *adminService) CreateBooking(payloads adminDto.BookingDto) (adminDto.BookingDto, error) {
	var dto adminDto.BookingDto

	// get max of queue of booking table by session_id
	res2, err := s.adminRepository.GetMaxQueue(payloads.SessionId)
	if err != nil {
		return dto, err
	}

	queue := strconv.Itoa(res2.TotalQ)
	// Check if queue nil or not
	if queue != "" {
		res2.TotalQ = res2.TotalQ + 1
	} else {
		res2.TotalQ = 1
	}

	// Set and get session by id
	session_dto := adminDto.SessionWithStatusDTO{
		SessionId: payloads.SessionId,
	}
	res3, err := s.adminRepository.GetSessionById(session_dto)
	if err != nil {
		return dto, err
	}

	if res3.Kuota < 1 {
		return dto, errors.New("kuota vaksin sudah habis")
	}

	// create citizen account
	res1, err := s.adminRepository.CreateCitizenBook(payloads.Nik, payloads.Nama, payloads.Address)
	if err != nil {
		return dto, err
	}

	temp := adminDto.BookingDto{
		CitizenId: res1.CitizenID,
		SessionId: payloads.SessionId,
		Nama:      payloads.Nama,
		Nik:       payloads.Nik,
		Address:   payloads.Address,
		Queue:     res2.TotalQ,
	}

	res, err := s.adminRepository.CreateBooking(temp)

	if err != nil {
		return res, err
	}

	// Update kuota vaksin on table session
	kuota := res3.Kuota - 1
	convKuota := strconv.Itoa(kuota)
	errK := s.adminRepository.UpdateSessionBooking(payloads.SessionId, convKuota)
	if errK != nil {
		return dto, errK
	}

	return res, nil
}
