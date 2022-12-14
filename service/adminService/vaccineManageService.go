package adminService

import (
	"capstone_vaccine/dto/adminDto"
	"errors"
)

// TODO Create Vaccine
func (s *adminService) CreateVaccine(input adminDto.VaccineRequest) (adminDto.VaccineResponse, error) {

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

// TODO View All Vaccine
func (s *adminService) ViewAllVaccine(medicalId uint) ([]adminDto.VaccineDTO, error) {
	var vaccine []adminDto.VaccineDTO

	res, err := s.adminRepository.ViewAllVaccine(medicalId)

	if err != nil {
		return nil, err
	}
	for _, v := range res {
		vaccine = append(vaccine, adminDto.VaccineDTO{
			VaccineID: v.VaccineID,
			Name:      v.Name,
			Kuota:     v.Kuota,
			Expired:   v.Expired,
		})
	}

	return vaccine, nil
}

// TODO Update Vaccine
func (s *adminService) UpdateVaccine(payloads adminDto.VaccineDTO, medicalId uint) (adminDto.VaccineDTO, error) {

	return s.adminRepository.UpdateVaccine(payloads, medicalId)
}

// TODO DELETE VACCINE
func (s *adminService) DeleteVaccine(data adminDto.VaccineDTO, medicalId uint) error {
	return s.adminRepository.DeleteVaccine(data, medicalId)

}

// TODO GET VACCINE
func (s *adminService) GetVaccineById(vaccineId uint, medicalId uint) (adminDto.VaccineDTO, error) {
	res, err := s.adminRepository.GetVaccineById(vaccineId, medicalId)

	if res.VaccineID < 1 {
		return res, errors.New("record not found")
	}

	if err != nil {
		return res, nil
	}

	return res, nil
}
