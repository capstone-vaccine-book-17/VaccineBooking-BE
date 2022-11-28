package adminService

import (
	"capstone_vaccine/dto/adminDto"
	"capstone_vaccine/model"
	"errors"
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

func (s *adminService) GetVaccineById(payloads adminDto.VaccineRequest) (adminDto.VaccineDTO, error) {
	res, err := s.adminRepository.GetVaccineById(payloads)

	if res.VaccineID < 1 {
		return res, errors.New("record not found")
	}

	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *adminService) UpdateVaccine(payloads adminDto.VaccineRequest) (adminDto.VaccineRequest, error) {
	// DTO for get data kuota on session table by id
	// dto_vaccineById, _:= s.adminRepository.GetVaccineById(payloads)

	// // dto_update := adminDto.VaccineRequest{}

	// // kuotaById, _ := s.adminRepository.GetVaccineById()(dto_vaccineById)

	// // kuota, _ := s.adminRepository.CountKuota(payloads.VaccineId)
	// // // Check if kuota already maximum
	// // kuota.TotalS = kuota.TotalS - kuotaById.Kuota

	// // if kuota.TotalS >= kuota.TotalV {
	// // 	return dto_update, errors.New("kuota vaksin yang di input melebihi batas")
	// // } else if payloads.Kuota+kuota.TotalS > kuota.TotalV {
	// // 	return dto_update, errors.New("kuota vaksin yang di input melebihi batas")
	// // }

	res, err := s.adminRepository.UpdateVaccine(payloads)

	if err != nil {
		return res, err
	}

	return res, nil
}

// /func (s *adminService) UpdateVaccine(updateReq adminDto.VaccineRequest) (adminDto.VaccineDTO, error) {
// 	//temp := adminDto.VaccineRequest{}

// 	// temp.MedicalFacilitysId = updateReq.MedicalFacilitysId,
// 	// temp.VaccineID = updateReq.VaccineID,
// 	// temp.Name = updateReq.Name,
// 	// temp.Kuota = updateReq.Kuota,
// 	// temp.Expired =updateReq.Expired,

// 	res, err := s.adminRepository.UpdateVaccine(updateReq)

// 	if err != nil {
// 		return res, err
// 	}

// 	return res, nil

// }

func (s *adminService) DeleteVaccine(data adminDto.VaccineRequest) ([]model.VaccineVarietie, error) {
	res, err := s.adminRepository.DeleteVaccine(data)

	if err != nil {
		return nil, err
	}

	return res, nil
}
