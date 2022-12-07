package citizenService

import "capstone_vaccine/dto/citizenDto"

// TODO GET SESSION BY MEDICAL FACILITYS ID
func (s *citizenService) GetSessionByMedicalId(medicalID uint) ([]citizenDto.SessionDto, error) {
	res, err := s.citizenRepository.GetSessionByMedicalId(medicalID)

	if err != nil {
		return nil, err
	}

	return res, nil
}
