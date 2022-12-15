package adminService

import (
	"capstone_vaccine/dto/adminDto"
	"errors"
)

func (s *adminService) CreateRoles(payloads adminDto.RoleDTO) (adminDto.RoleDTO, error) {

	temp := adminDto.RoleDTO{
		Name: payloads.Name,
	}

	if temp.Name == "" {
		return temp, errors.New("name cannot be empty")
	}

	res, err := s.adminRepository.CreateRoles(payloads)

	if err != nil {
		return res, err
	}

	return temp, nil
}
