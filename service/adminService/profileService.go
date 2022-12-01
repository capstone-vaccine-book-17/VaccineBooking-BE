package adminService

import (
	"capstone_vaccine/dto/adminDto"
	"errors"

	"golang.org/x/crypto/bcrypt"
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

	dto := adminDto.ProfileRequest{
		AdminID:            payloads.AdminID,
		MedicalFacilitysId: payloads.MedicalFacilitysId,
		Name:               payloads.Name,
		Image:              payloads.Image,
		Address:            payloads.Address,
		Username:           payloads.Username,
		Password:           payloads.Password,
		NewPassword:        payloads.NewPassword,
	}
	new, _ := s.adminRepository.GetAdmin(payloads)

	if err := bcrypt.CompareHashAndPassword([]byte(new.Password), []byte(payloads.Password)); err != nil {

		return dto, errors.New("username or password incorrect")

	}

	// temp, _ := s.adminRepository.UpdateProfile(payloads)
	// // hash, _ := HashPassword(temp.NewPassword)

	_, err := s.adminRepository.UpdateProfile(dto)

	if err != nil {
		return dto, err
	}

	return dto, nil
}

func (s *adminService) UpdateImage(payloads adminDto.ProfileRequest) (adminDto.ProfilDTO, error) {

	temp := adminDto.ProfileRequest{
		MedicalFacilitysId: payloads.MedicalFacilitysId,
		Image:              payloads.Image,
	}
	res, err := s.adminRepository.UpdateImage(temp)

	out := adminDto.ProfilDTO{
		Image: res.Image,
	}

	if err != nil {
		return out, err
	}

	return out, nil

}
