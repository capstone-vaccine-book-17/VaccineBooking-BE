package citizenService

import (
	"capstone_vaccine/dto/citizenDto"
	"capstone_vaccine/utils"
)

// TODO Get Profile
func (s *citizenService) GetProfile(payloads citizenDto.ProfileReq) (citizenDto.ProfileDTO, error) {

	return s.citizenRepository.GetProfile(payloads)

}

// TODO Update & Upload Image
func (s *citizenService) UploadImage(payloads citizenDto.ProfileReq) (citizenDto.ProfileReq, error) {

	temp := citizenDto.ProfileReq{
		CitizenID: payloads.CitizenID,
		Image:     payloads.Image,
	}
	res, err := s.citizenRepository.UploadImage(temp)

	out := citizenDto.ProfileReq{
		Image: res.Image,
	}

	if err != nil {
		return out, err
	}

	return out, nil

}

// TODO Get Personal Data
func (s *citizenService) GetPersonalData(payload citizenDto.ProfileReq) ([]citizenDto.PersonalData, error) {

	var profile []citizenDto.PersonalData

	res, err := s.citizenRepository.GetPersonalData(payload)

	if err != nil {
		return nil, err
	}

	for _, p := range res {

		profile = append(profile, citizenDto.PersonalData{
			Name:    p.Name,
			Nik:     p.Nik,
			Age:     uint(utils.Age(p.Dob)),
			Address: p.Address,
			Email:   p.Email,
			Gender:  p.Gender,
			Dob:     p.Dob,
		})
	}

	return profile, nil

}

// TODO Update Detail Address
func (s *citizenService) UpdateAddress(payloads citizenDto.AddressCitizenReq) error {

	return s.citizenRepository.UpdateAddress(payloads)

}

// TODO GET Detail Address
func (s *citizenService) GetAddress(payload citizenDto.ProfileReq) (citizenDto.AddressResp, error) {

	return s.citizenRepository.GetAddress(payload)
}

// TODO Get Email
func (s *citizenService) GetEmail(payloads citizenDto.ProfileReq) (citizenDto.LoginDto, error) {

	return s.citizenRepository.GetEmail(payloads)

}

// TODO Update Email
func (s *citizenService) UpdateEmail(payloads citizenDto.UpdateEmail) error {

	return s.citizenRepository.UpdateEmail(payloads)

}

// TODO Update Password
func (s *citizenService) UpdatePassword(payloads citizenDto.UpdatePassword) (citizenDto.LoginDto, error) {

	return s.citizenRepository.UpdatePassword(payloads)

}
