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

	LoginAdmin(payloads adminDto.LoginDTO) (model.Admin, error)

	// TODO ROLES
	CreateRoles(payloads adminDto.RoleDTO) (adminDto.RoleDTO, error)

	// TODO DASHBOARD

	GetDashboard() (adminDto.CountDashboard, error)

	// TODO SESSION
	CountKuota(vaccineID uint) (adminDto.CountKuota, error)
	CreateSession(payloads adminDto.SessionRequest) (adminDto.SessionDTO, error)
	GetAllSession() ([]adminDto.SessionWithStatusDTO, error)
	GetSessionById(payloads adminDto.SessionWithStatusDTO) (adminDto.SessionWithStatusDTO, error)
	UpdateSession(payloads adminDto.SessionRequestUpdate) (adminDto.SessionRequestUpdate, error)
	DeleteSession(payloads adminDto.SessionWithStatusDTO) error
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
	if err := u.db.Model(&model.VaccineVarietie{}).Select("sum(kuota) as vaccine_available").Where("expired <= ?", convDate).Find(&vaccineAvail).Error; err != nil {
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
