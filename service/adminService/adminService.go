package adminService

import (
	"capstone_vaccine/dto/adminDto"
	"capstone_vaccine/middleware"
	"capstone_vaccine/repository/adminRepository"
	"capstone_vaccine/utils"
	"errors"
	"mime/multipart"

	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	// TODO AUTH
	RegisterAdmin(payloads adminDto.RegisterAdminDto) (adminDto.RegisterAdminDto, error)
	LoginAdmin(payloads adminDto.LoginDTO) (adminDto.LoginJWT, error)

	// TODO ROLES
	CreateRoles(payloads adminDto.RoleDTO) (adminDto.RoleDTO, error)

	// TODO MEDICAL FACILITYS
	CreateMedical(payloads adminDto.MedicalDto) (adminDto.MedicalDto, error)

	// TODO DASHBOARD

	GetDashboard(medicalId uint) (adminDto.CountDashboard, error)

	// TODO SESSION
	CreateSession(payloads adminDto.SessionRequest) (adminDto.SessionDTO, error)
	GetAllSession(medicalId uint) ([]adminDto.SessionWithStatusDTO, error)
	GetSessionById(payloads adminDto.SessionWithStatusDTO) (adminDto.SessionWithStatusDTO, error)
	UpdateSession(payloads adminDto.SessionRequestUpdate) (adminDto.SessionRequestUpdate, error)
	DeleteSession(payloads adminDto.SessionWithStatusDTO) error

	// TODO Manage Vaccine
	CreateVaccine(input adminDto.VaccineRequest) (adminDto.VaccineResponse, error)
	ViewAllVaccine(medicalId uint) ([]adminDto.VaccineDTO, error)
	UpdateVaccine(payloads adminDto.VaccineDTO, medicalId uint) (adminDto.VaccineDTO, error)
	DeleteVaccine(data adminDto.VaccineDTO, medicalId uint) error
	GetVaccineById(vaccineId uint, medicalId uint) (adminDto.VaccineDTO, error)

	// TODO Profile
	GetProfile(payloads adminDto.ProfileRequest) ([]adminDto.ProfilDTO, error)
	UpdateProfile(payloads adminDto.ProfileRequest) (adminDto.ProfileRequest, error)
	UpdateImage(payloads adminDto.ProfileRequest, file multipart.File) error

	// TODO BOOKING
	CreateBooking(payloads adminDto.BookingDto) (adminDto.BookingDto, error)
	UpdateBooking(payloads adminDto.UpdateBooking) (adminDto.UpdateBooking, error)
	GetAllBooking(medicalId uint) ([]adminDto.BookingAllDto, error)
	GetBookingById(payloads adminDto.BookingAllDto) (adminDto.BookingAllDto, error)
	DeleteBooking(payloads adminDto.BookingAllDto) error
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

// TODO REGISTER ADMIN
func (s *adminService) RegisterAdmin(payloads adminDto.RegisterAdminDto) (adminDto.RegisterAdminDto, error) {

	// encrypt password
	pw, err := utils.HashBcrypt(payloads.Password)

	if err != nil {
		return payloads, err
	}

	payloads.Password = pw

	res, err := s.adminRepository.RegisterAdmin(payloads)

	if err != nil {
		return res, err
	}

	return res, nil
}

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

// TODO DASHBOARD
// TODO GET DASHBOARD
func (s *adminService) GetDashboard(medicalId uint) (adminDto.CountDashboard, error) {
	res, err := s.adminRepository.GetDashboard(medicalId)

	if err != nil {
		return res, err
	}

	return res, nil
}
