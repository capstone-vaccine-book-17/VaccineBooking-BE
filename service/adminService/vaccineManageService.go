package adminService

import (
	"capstone_vaccine/dto/adminDto"
)

func (s *adminService) CreateVaccine(input adminDto.VaccineRequest) (adminDto.VaccineDTO, error) {
	// var dto adminDto.VaccineDTO

	// kuota, _ := s.adminRepository.CountKuota(payloads.VaccineId)
	// if kuota.TotalS >= kuota.TotalV {
	// 	return dto, errors.New("kuota vaksin yang di input melebihi batas")
	// } else if payloads.Kuota > kuota.TotalV {
	// 	return dto, errors.New("kuota vaksin yang di input melebihi batas")
	// }

	temp := adminDto.VaccineRequest{
		Name:               input.Name,
		MedicalFacilitysId: input.MedicalFacilitysId,
		Kuota:              input.Kuota,
		Expired:            input.Expired,
	}

	res, err := s.adminRepository.CreateVaccine(temp)
	if err != nil {
		return res, err
	}
	return res, nil
}
