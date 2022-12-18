package mocks

import "capstone_vaccine/dto/adminDto"

func (u *MockAdmin) CreateRoles(payloads adminDto.RoleDTO) (adminDto.RoleDTO, error) {
	args := u.Called(payloads)

	return args.Get(0).(adminDto.RoleDTO), args.Error(1)
}
