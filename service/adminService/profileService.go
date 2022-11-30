package adminService

import (
	"capstone_vaccine/dto/adminDto"
)

func (s *adminService) GetProfile(payloads adminDto.ProfileRequest) ([]adminDto.ProfilDTO, error) {

	var profile []adminDto.ProfilDTO

	res, err := s.adminRepository.GetProfile(payloads)

	if err != nil {
		return nil, err
	}
	for _, p := range res {

		profile = append(profile, adminDto.ProfilDTO{
			Name:     p.Name,
			Image:    p.Image,
			Address:  p.Address,
			Username: p.Username,
			Password: p.Password,
		})
	}
	return profile, nil
}

func (s *adminService) UpdateProfile(payloads adminDto.ProfileRequest) (adminDto.ProfileRequest, error) {

	res, err := s.adminRepository.UpdateProfile(payloads)

	if err != nil {
		return res, err
	}

	return res, nil
}
