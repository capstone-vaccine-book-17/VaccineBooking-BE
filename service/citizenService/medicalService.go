package citizenService

import (
	"capstone_vaccine/dto/citizenDto"
)

// TODO GET ALL MEDICAL BY CITY AND BY SEARCH
func (s *citizenService) GetMedicalByCity(payloads citizenDto.SearchKey) ([]citizenDto.SearchDto, error) {

	city, errC := s.citizenRepository.GetCityCitizen(payloads.CitizenId)

	if errC != nil {
		return nil, errC
	}

	payloads.City = city.City

	res, errM := s.citizenRepository.GetMedicalByCity(payloads)

	if errM != nil {
		return nil, errM
	}

	return res, nil
}

// TODO GET MEDICAL BY ID
func (s *citizenService) GetMedicalById(medicalID uint) (citizenDto.MedicalDto, error) {
	res, err := s.citizenRepository.GetMedicalById(medicalID)

	if err != nil {
		return res, err
	}

	return res, nil
}
