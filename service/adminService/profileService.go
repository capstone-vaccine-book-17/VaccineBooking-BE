package adminService

import (
	"capstone_vaccine/dto/adminDto"
)

func (s *adminService) GetProfile() ([]adminDto.ProfilDTO, error) {

	var profile []adminDto.ProfilDTO

	res, err := s.adminRepository.GetProfile()

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
