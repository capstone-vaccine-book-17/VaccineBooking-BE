package adminService

import (
	"capstone_vaccine/dto/adminDto"
	"capstone_vaccine/middleware"
	"capstone_vaccine/model"
	"capstone_vaccine/repository/adminRepository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	// TODO AUTH
	LoginAdmin(payloads adminDto.LoginDTO) (adminDto.LoginJWT, error)

	// TODO ROLES
	CreateRoles(payloads adminDto.RoleDTO) (adminDto.RoleDTO, error)

	// TODO CreateVaccine
	CreateVaccine(input adminDto.VaccineRequest) (adminDto.VaccineDTO, error)

	// TODO ViewAllVaccine
	ViewAllVaccine() ([]adminDto.VaccineDTO, error)

	// TODO GET SESSION BY ID
	GetVaccineById(payloads adminDto.VaccineRequest) (adminDto.VaccineDTO, error)

	// TODO UpdateVaccine
	UpdateVaccine(payloads adminDto.VaccineRequest) (adminDto.VaccineRequest, error)

	// TODO DeleteVaccine
	DeleteVaccine(data adminDto.VaccineRequest) ([]model.VaccineVarietie, error)
}

type adminService struct {
	adminRepository adminRepository.AdminRepository
}

func NewAdminService(adminRepo adminRepository.AdminRepository) *adminService {
	return &adminService{
		adminRepository: adminRepo,
	}
}

// TODO ADMIN SERVICE HERE

// TODO LOGIN ADMIN
func (s *adminService) LoginAdmin(payloads adminDto.LoginDTO) (adminDto.LoginJWT, error) {
	var temp adminDto.LoginJWT

	res, err := s.adminRepository.LoginAdmin(payloads)

	if errh := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(payloads.Password)); errh != nil {
		return temp, errors.New("username or password incorrect")
	}

	token, errt := middleware.CreateToken(res.AdminID, res.RoleId, res.MedicalFacilitysId, res.Username)

	temp = adminDto.LoginJWT{
		Username: res.Username,
		Token:    token,
	}

	if err != nil {
		return temp, err
	}

	if errt != nil {
		return temp, errt
	}

	return temp, nil
}
