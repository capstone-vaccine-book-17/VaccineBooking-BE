package adminRepository

import (
	"capstone_vaccine/dto/adminDto"
	"capstone_vaccine/model"
	"time"
)

func (u *adminRepository) CreateRoles(payloads adminDto.RoleDTO) (adminDto.RoleDTO, error) {

	temp := adminDto.RoleDTO{
		Name: payloads.Name,
	}

	if err := u.db.Create(&model.Role{
		Name:      temp.Name,
		CreatedAT: time.Now(),
	}).Error; err != nil {
		return temp, err
	}

	return temp, nil
}
