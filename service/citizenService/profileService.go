package citizenService

import (
	"capstone_vaccine/dto/citizenDto"
	"capstone_vaccine/utils"
)

// TODO Get Profile
func (s *citizenService) GetProfile(payloads citizenDto.ProfileReq) (citizenDto.ProfileDTO, error) {

	res, err := s.citizenRepository.GetProfile(payloads)

	if err != nil {
		return res, err
	}

	return res, nil
}

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

// TODO Personal Data
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

func (s *citizenService) UpdateAddress(payloads citizenDto.AddressCitizenReq) error {

	err := s.citizenRepository.UpdateAddress(payloads)

	if err != nil {
		return err
	}

	return nil
}

func (s *citizenService) GetAddress(payload citizenDto.ProfileReq) (citizenDto.AddressResp, error) {

	res, err := s.citizenRepository.GetAddress(payload)

	if err != nil {
		return res, err
	}

	return res, nil
}

// TODO Get Email
func (s *citizenService) GetEmail(payloads citizenDto.ProfileReq) (citizenDto.PersonalData, error) {

	res, err := s.citizenRepository.GetEmail(payloads)

	if err != nil {
		return res, err
	}

	return res, nil
}

// TODO Get Email
func (s *citizenService) UpdateEmail(payloads citizenDto.UpdateEmail) error {

	err := s.citizenRepository.UpdateEmail(payloads)

	if err != nil {
		return err
	}

	return nil
}
