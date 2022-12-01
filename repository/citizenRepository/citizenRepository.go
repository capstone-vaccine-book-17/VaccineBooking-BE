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
