package adminRepository

import (
	"capstone_vaccine/dto/adminDto"
	"capstone_vaccine/model"
	"errors"
	"time"

	"github.com/leekchan/accounting"
	"gorm.io/gorm"
)

type AdminRepository interface {
	// TODO AUTH
	RegisterAdmin(payloads adminDto.RegisterAdminDto) (adminDto.RegisterAdminDto, error)
	LoginAdmin(payloads adminDto.LoginDTO) (model.Admin, error)

	// TODO ROLES
	CreateRoles(payloads adminDto.RoleDTO) (adminDto.RoleDTO, error)

	// TODO MEDICAL FACILITYS
	CreateMedical(payloads adminDto.MedicalDto) (adminDto.MedicalDto, error)

	// TODO DASHBOARD

	GetDashboard() (adminDto.CountDashboard, error)

	// TODO SESSION
	CountKuota(vaccineID uint) (adminDto.CountKuota, error)
	AutoUpdateSession(dateR, timeR string) error
	CreateSession(payloads adminDto.SessionRequest) (adminDto.SessionDTO, error)
	GetAllSession() ([]adminDto.SessionWithStatusDTO, error)
	GetSessionById(payloads adminDto.SessionWithStatusDTO) (adminDto.SessionWithStatusDTO, error)
	UpdateSession(payloads adminDto.SessionRequestUpdate) (adminDto.SessionRequestUpdate, error)
	DeleteSession(payloads adminDto.SessionWithStatusDTO) error

	// TODO Manage Vaccine
	CreateVaccine(input adminDto.VaccineRequest) (adminDto.VaccineResponse, error)
	ViewAllVaccine() ([]adminDto.VaccineDTO, error)
	UpdateVaccine(payloads adminDto.VaccineDTO) (adminDto.VaccineDTO, error)
	DeleteVaccine(data adminDto.VaccineDTO) error
	GetVaccineById(vaccineId uint) (adminDto.VaccineDTO, error)

	// TODO Profile
	GetProfile(payload adminDto.ProfileRequest) ([]adminDto.ProfilDTO, error)
	UpdateProfile(payloads adminDto.ProfileRequest) (adminDto.Address, error)
	GetAdmin(payloads adminDto.ProfileRequest) (adminDto.Address, error)
	UpdateImage(payloads adminDto.ProfileRequest) (adminDto.ProfilDTO, error)
	// TODO BOOKING
	CreateCitizenBook(nik, nama, address string) (model.Citizen, error)
	CreateBooking(payloads adminDto.BookingDto) (adminDto.BookingDto, error)
	GetMaxQueue(session_id uint) (adminDto.MaxQueue, error)
	UpdateSessionBooking(session_id uint, kuota string) error
	UpdateBooking(payloads adminDto.UpdateBooking) (adminDto.UpdateBooking, error)
	GetAllBooking() ([]adminDto.BookingAllDto, error)
	GetBookingById(payloads adminDto.BookingAllDto) (adminDto.BookingAllDto, error)
	DeleteBooking(payloads adminDto.BookingAllDto) error
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{db}
}

// TODO ADMIN REPOSITORY HERE

// TODO LOGIN ADMIN
func (u *adminRepository) LoginAdmin(payloads adminDto.LoginDTO) (model.Admin, error) {
	var admin model.Admin

	query := u.db.Where("username = ?", payloads.Username).First(&admin)
	if query.Error != nil {
		return admin, query.Error
	}

	if query.RowsAffected < 1 {
		return admin, errors.New("username is incorrect")
	}

	return admin, nil
}

// TODO REGISTER ADMIN
func (u *adminRepository) RegisterAdmin(payloads adminDto.RegisterAdminDto) (adminDto.RegisterAdminDto, error) {

	if err := u.db.Create(&model.Admin{
		RoleId:             payloads.RoleId,
		MedicalFacilitysId: payloads.MedicalId,
		Username:           payloads.Username,
		Password:           payloads.Password,
		CreatedAT:          time.Now(),
	}).Error; err != nil {
		return payloads, err
	}

	return payloads, nil
}

// TODO DASHBOARD ADMIN
func (u *adminRepository) GetDashboard() (adminDto.CountDashboard, error) {
	dto := adminDto.CountDashboard{}
	var (
		vaccineAvail      int
		bookingToday      int64
		bookingRegistered int64
		convDate          string
	)
	today := time.Now()
	date := today.Format("2006-01-02")
	convDate = string(date)

	// QUERY GET VACCINE AVAILABLE
	if err := u.db.Model(&model.VaccineVarietie{}).Select("coalesce(sum(kuota), 0) as vaccine_available").Where("expired >= ?", convDate).Find(&vaccineAvail).Error; err != nil {
		return dto, err
	}

	// QUERY GET BOOKING TODAY
	if err := u.db.Model(&model.Booking{}).Where("created_at like ?", "%"+convDate+"%").Count(&bookingToday).Error; err != nil {
		return dto, err
	}

	// QUERY GET ALL BOOKING
	if err := u.db.Model(&model.Booking{}).Count(&bookingRegistered).Error; err != nil {
		return dto, err
	}

	// USING FORMAT MONEY FROM leekchan/accounting
	ac := accounting.Accounting{Precision: 0}

	dto.VaccineAvailable = ac.FormatMoney(vaccineAvail)
	dto.BookingToday = ac.FormatMoney(bookingToday)
	dto.BookingsRegistered = ac.FormatMoney(bookingRegistered)

	return dto, nil
}
