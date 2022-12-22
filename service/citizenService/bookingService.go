package citizenService

import (
	"capstone_vaccine/dto/citizenDto"
	"errors"
	"strconv"
)

// TODO CREATE BOOKING
func (s *citizenService) CreateBooking(payloads citizenDto.BookingDto) (citizenDto.BookingDto, error) {
	// get max of queue of booking table by session_id
	max, err := s.citizenRepository.GetMaxQueue(payloads.CitizenID)
	if err != nil {
		return payloads, err
	}

	queue := strconv.Itoa(max.TotalQ)
	// Check if queue nil or not
	if queue != "" {
		max.TotalQ = max.TotalQ + 1
	} else {
		max.TotalQ = 1
	}

	session, err := s.citizenRepository.GetSessionById(payloads.SessionID)

	if err != nil {
		return payloads, err
	}

	if session.Kuota < 1 {
		return payloads, errors.New("kuota vaksin sudah habis")
	}

	temp := citizenDto.BookingDto{
		CitizenID: payloads.CitizenID,
		SessionID: payloads.SessionID,
		Queue:     max.TotalQ,
	}

	res, errR := s.citizenRepository.CreateBooking(temp)

	if errR != nil {
		return res, errR
	}

	// Update kuota vaksin on table session
	kuota := session.Kuota - 1
	convKuota := strconv.Itoa(kuota)
	errK := s.citizenRepository.UpdateKuotaSession(payloads.SessionID, convKuota)
	if errK != nil {
		return res, errK
	}

	return res, nil
}

// TODO GET LAST BOOKING FOR TICKET
func (s *citizenService) GetLastBooking(citizenId uint) (citizenDto.TicketBooking, error) {
	return s.citizenRepository.GetLastBooking(citizenId)

}
