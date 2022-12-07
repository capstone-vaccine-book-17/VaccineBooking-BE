package citizenService

import "capstone_vaccine/dto/citizenDto"

func (s *citizenService) GetTicket(citizenId uint) ([]citizenDto.TicketDetails, error) {
	res, err := s.citizenRepository.GetTicket(citizenId)

	if err != nil {
		return res, err
	}

	

	return res, nil
}
