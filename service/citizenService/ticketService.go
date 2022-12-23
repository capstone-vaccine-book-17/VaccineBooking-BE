package citizenService

import "capstone_vaccine/dto/citizenDto"

// TODO GET All Ticket
func (s *citizenService) GetAllTicket(citizenId uint) ([]citizenDto.TicketDetails, error) {
	return s.citizenRepository.GetAllTicket(citizenId)

}

//GET Tiket By Status
func (s *citizenService) GetTicketOnStatus(citizenId uint, status string) ([]citizenDto.TicketDetails, error) {
	return s.citizenRepository.GetTicketOnStatus(citizenId, status)

}

//Get Ticket By Id
func (s *citizenService) GetTicket(bookingId uint) (citizenDto.TicketDetails, error) {
	return s.citizenRepository.GetTicket(bookingId)

}
