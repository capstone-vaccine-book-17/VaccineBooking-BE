package mocks

import (
	"capstone_vaccine/dto/adminDto"

	"github.com/stretchr/testify/mock"
)

type MockAdmin struct {
	mock.Mock
}

// MOCK Register
func (u *MockAdmin) RegisterAdmin(payloads adminDto.RegisterAdminDto) (adminDto.RegisterAdminDto, error) {
	args := u.Called(payloads)

	return args.Get(0).(adminDto.RegisterAdminDto), args.Error(1)
}

// MOCK Login
func (u *MockAdmin) LoginAdmin(payloads adminDto.LoginDTO) (adminDto.LoginJWT, error) {
	args := u.Called(payloads)

	return args.Get(0).(adminDto.LoginJWT), args.Error(1)
}

// MOCK Dashboard
func (u *MockAdmin) GetDashboard() (adminDto.CountDashboard, error) {
	args := u.Called()

	return args.Get(0).(adminDto.CountDashboard), args.Error(1)
}
