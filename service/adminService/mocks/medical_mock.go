package mocks

import "capstone_vaccine/dto/adminDto"

// MOCK Create Medical
func (u *MockAdmin) CreateMedical(payloads adminDto.MedicalDto) (adminDto.MedicalDto, error) {
	args := u.Called(payloads)

	return args.Get(0).(adminDto.MedicalDto), args.Error(1)
}
