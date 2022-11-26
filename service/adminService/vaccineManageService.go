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

func (s *adminService) ViewAllVaccine() ([]adminDto.VaccineDTO, error) {
	var vaccine []adminDto.VaccineDTO

	res, err := s.adminRepository.ViewAllVaccine()

	if err != nil {
		return nil, err
	}
	for _, v := range res {
		vaccine = append(vaccine, adminDto.VaccineDTO{
			Name:    v.Name,
			Kuota:   v.Kuota,
			Expired: v.Expired,
		})
	}

	return vaccine, nil
}

func(s *adminService) UpdateVaccine(updateReq adminDto.VaccineRequest) (adminDto.VaccineDTO, error){
	temp :=adminDto.VaccineRequest{}

	

}
