package citizenService

import (
	"capstone_vaccine/dto/citizenDto"
)

// TODO Get Profile
func (s *citizenService) GetProfile(payloads citizenDto.ProfileReq) (citizenDto.ProfileDTO, error) {

	res, err := s.citizenRepository.GetProfile(payloads)

	if err != nil {
		return res, err
	}

	return res, nil
}
