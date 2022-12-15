package mocks

import "capstone_vaccine/dto/adminDto"

func (u *MockAdmin) CreateVaccine(input adminDto.VaccineRequest) (adminDto.VaccineResponse, error) {
	args := u.Called(input)

	return args.Get(0).(adminDto.VaccineResponse), args.Error(1)
}
func (u *MockAdmin) ViewAllVaccine(medicalId uint) ([]adminDto.VaccineDTO, error) {

	args := u.Called(medicalId)

	return args.Get(0).([]adminDto.VaccineDTO), args.Error(1)

}
func (u *MockAdmin) UpdateVaccine(payloads adminDto.VaccineDTO, medicalId uint) (adminDto.VaccineDTO, error) {

	args := u.Called(payloads)

	return args.Get(0).(adminDto.VaccineDTO), args.Error(1)
}
func (u *MockAdmin) DeleteVaccine(data adminDto.VaccineDTO, medicalId uint) error {
	args := u.Called(data.VaccineID, medicalId)

	return args.Error(0)
}
func (u *MockAdmin) GetVaccineById(vaccineId uint, medicalId uint) (adminDto.VaccineDTO, error) {

	args := u.Called(vaccineId, medicalId)

	return args.Get(0).(adminDto.VaccineDTO), args.Error(1)
}
