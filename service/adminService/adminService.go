package adminService

import (
	"capstone_vaccine/repository/adminRepository"
)

type AdminService interface {
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
