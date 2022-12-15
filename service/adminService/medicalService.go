package adminService

import "capstone_vaccine/dto/adminDto"

func (s *adminService) CreateMedical(payloads adminDto.MedicalDto) (adminDto.MedicalDto, error) {
	res, err := s.adminRepository.CreateMedical(payloads)

	if err != nil {
		return res, err
	}

	return res, nil
}
