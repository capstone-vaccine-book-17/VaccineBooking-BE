package adminService

import "capstone_vaccine/dto/adminDto"

func (s *adminService) CreateMedical(payloads adminDto.MedicalDto) (adminDto.MedicalDto, error) {
	return s.adminRepository.CreateMedical(payloads)

}
