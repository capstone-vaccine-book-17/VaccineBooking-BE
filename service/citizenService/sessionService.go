package citizenService

import "capstone_vaccine/dto/citizenDto"

// TODO GET SESSION BY MEDICAL FACILITYS ID
func (s *citizenService) GetSessionByMedicalId(medicalID uint) ([]citizenDto.SessionDto, error) {
	return s.citizenRepository.GetSessionByMedicalId(medicalID)

}
