package citizenService

import "capstone_vaccine/dto/citizenDto"


// TODO GET All Ticket
func (s *citizenService) GetAllTicket(citizenId uint) ([]citizenDto.TicketDetails, error) {
	res, err := s.citizenRepository.GetAllTicket(citizenId)

	if err != nil {
		return res, err
	}

	return res, nil

}

//GET Tiket By Status 
func (s *citizenService) GetTicketOnStatus(citizenId uint, status string) ([]citizenDto.TicketDetails, error) {
	res, err := s.citizenRepository.GetTicketOnStatus(citizenId, status)

	if err != nil {
		return res, err
	}

	return res, nil

}

//Get Ticket By Id
func (s *citizenService) GetTicket(bookingId uint) (citizenDto.TicketDetails, error) {
	res, err := s.citizenRepository.GetTicket(bookingId)

	if err != nil {
		return res, err
	}

	return res, nil

}
