package citizenService

import (
	"capstone_vaccine/dto/citizenDto"
	"capstone_vaccine/middleware"
	"capstone_vaccine/repository/citizenRepository"
	"capstone_vaccine/utils"
	"errors"
)

type CitizenService interface {
	// TODO AUTH
	LoginCitizen(payloads citizenDto.LoginDto) (citizenDto.LoginJWT, error)
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

	// TODO MEDICAL
	GetMedicalByCity(payloads citizenDto.SearchKey) ([]citizenDto.SearchDto, error)
	GetMedicalById(medicalID uint) (citizenDto.MedicalDto, error)

	// TODO SESSION
	GetSessionByMedicalId(medicalID uint) ([]citizenDto.SessionDto, error)

	// TODO BOOKING
	CreateBooking(payloads citizenDto.BookingDto) (citizenDto.BookingDto, error)
	GetLastBooking(citizenId uint) (citizenDto.TicketBooking, error)
}

type citizenService struct {
	citizenRepository citizenRepository.CitizenRepository
}

func NewCitizenService(citizenRepo citizenRepository.CitizenRepository) *citizenService {
	return &citizenService{
		citizenRepository: citizenRepo,
	}
}

// TODO REGISTER
func (s *citizenService) RegisterCitizen(payloads citizenDto.RegisterDto) (citizenDto.RegisterDto, error) {

	pw, err := utils.HashBcrypt(payloads.Password)

	if err != nil {
		return payloads, err
	}

	payloads.Password = pw

	res, err := s.citizenRepository.RegisterCitizen(payloads)

	if err != nil {
		return res, err
	}

	return res, nil
}

// TODO LOGIN
func (s *citizenService) LoginCitizen(payloads citizenDto.LoginDto) (citizenDto.LoginJWT, error) {
	var temp citizenDto.LoginJWT

	res, err := s.citizenRepository.LoginCitizen(payloads)

	if res.CitizenID < 1 {
		return temp, errors.New("username or password incorrect")
	}

	if errh := utils.CompareHash(res.Password, payloads.Password); errh != nil {
		return temp, errors.New("username or password incorrect")
	}

	token, errt := middleware.CreateTokenCitizen(res.CitizenID, res.Nik)

	temp = citizenDto.LoginJWT{
		Email: res.Email,
		Token: token,
	}

	if err != nil {
		return temp, err
	}

	if errt != nil {
		return temp, errt
	}

	return temp, nil
}
