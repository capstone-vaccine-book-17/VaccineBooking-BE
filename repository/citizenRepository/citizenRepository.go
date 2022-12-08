package citizenRepository

import (
	"capstone_vaccine/dto/citizenDto"
	"capstone_vaccine/model"
	"time"

	"gorm.io/gorm"
)

type CitizenRepository interface {
	// TODO AUTH
	LoginCitizen(payloads citizenDto.LoginDto) (model.Citizen, error)
	RegisterCitizen(payloads citizenDto.RegisterDto) (citizenDto.RegisterDto, error)

	//TODO Profile
	GetProfile(payloads citizenDto.ProfileReq) (citizenDto.ProfileDTO, error)
	UploadImage(payloads citizenDto.ProfileReq) (citizenDto.ProfileReq, error)
	GetPersonalData(payload citizenDto.ProfileReq) ([]citizenDto.PersonalData, error)
	UpdateAddress(payloads citizenDto.AddressCitizenReq) error
	GetAddress(payload citizenDto.ProfileReq) (citizenDto.AddressResp, error)
	GetEmail(payloads citizenDto.ProfileReq) (citizenDto.LoginDto, error)
	UpdateEmail(payloads citizenDto.UpdateEmail) error
	UpdatePassword(payloads citizenDto.UpdatePassword) (citizenDto.LoginDto, error)

	//TODO Family Members
	CreateFamilyMember(payloads citizenDto.FamilyReq) error
	GetFamilys(payloads citizenDto.FamilyReq) ([]citizenDto.FamilylDTO, error)
	DeleteMember(payloads citizenDto.FamilylDTO) error
	GetDetailMember(payload citizenDto.FamilylDTO) (citizenDto.FamilylDTO, error)

	// TODO MEDICAL
	GetMedicalByCity(payloads citizenDto.SearchKey) ([]citizenDto.SearchDto, error)
	GetCityCitizen(id uint) (citizenDto.GetCity, error)
	GetMedicalById(medicalID uint) (citizenDto.MedicalDto, error)

	// TODO SESSION
	GetSessionByMedicalId(medicalID uint) ([]citizenDto.SessionDto, error)
	GetSessionById(id uint) (citizenDto.SessionWithVaccineId, error)

	// TODO BOOKING
	GetMaxQueue(session_id uint) (citizenDto.MaxQueue, error)
	UpdateKuotaSession(session_id uint, kuota string) error
	CreateBooking(payloads citizenDto.BookingDto) (citizenDto.BookingDto, error)
	GetLastBooking(citizenId uint) (citizenDto.TicketBooking, error)

	//TODO TICKET
	GetAllTicket(citizenId uint) ([]citizenDto.TicketDetails, error)
	GetTicketOnStatus(citizenId uint, status string) ([]citizenDto.TicketDetails, error)
	GetTicket(bookingId uint) (citizenDto.TicketDetails, error)
}

type citizenRepository struct {
	db *gorm.DB
}

func NewCitizenRepository(db *gorm.DB) *citizenRepository {
	return &citizenRepository{db}
}

// TODO REGISTER
func (u *citizenRepository) RegisterCitizen(payloads citizenDto.RegisterDto) (citizenDto.RegisterDto, error) {

	address := model.Address{
		Address:   payloads.Address,
		CreatedAT: time.Now(),
		UpdatedAT: time.Now(),
	}

	errA := u.db.Create(&address).Error
	if errA != nil {
		return payloads, errA
	}

	citizen := model.Citizen{
		AddressId: address.AddressID,
		Name:      payloads.Name,
		Nik:       payloads.Nik,
		Gender:    payloads.Gender,
		Email:     payloads.Email,
		Dob:       payloads.Dob,
		Password:  payloads.Password,
		CreatedAT: time.Now(),
		UpdatedAT: time.Now(),
	}

	errC := u.db.Create(&citizen).Error
	if errC != nil {
		return payloads, errC
	}

	return payloads, nil
}

// TODO LOGIN
func (u *citizenRepository) LoginCitizen(payloads citizenDto.LoginDto) (model.Citizen, error) {
	var citizen model.Citizen

	err := u.db.Where("email = ?", payloads.Email).First(&citizen).Error
	if err != nil {
		return citizen, err
	}

	return citizen, nil
}
